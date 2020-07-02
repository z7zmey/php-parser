%{
package php5

import (
    "bytes"
    "strconv"

	"github.com/z7zmey/php-parser/internal/position"
	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/token"
)

%}

%union{
    node                    ast.Vertex
    token                   *scanner.Token
    list                    []ast.Vertex
    simpleIndirectReference simpleIndirectReference

    ClassExtends            *ast.StmtClassExtends
    ClassImplements         *ast.StmtClassImplements
    InterfaceExtends        *ast.StmtInterfaceExtends
    ClosureUse              *ast.ExprClosureUse
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
                yylex.(*Parser).rootNode = &ast.Root{ast.Node{}, $1}
                yylex.(*Parser).rootNode.GetNode().Position = position.NewNodeListPosition($1)

                yylex.(*Parser).setFreeFloating(yylex.(*Parser).rootNode, token.End, yylex.(*Parser).currentToken.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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
                namePart.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(namePart, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_NS_SEPARATOR T_STRING
            {
                namePart := &ast.NameNamePart{ast.Node{}, $3.Value}
                $$ = append($1, namePart)

                // save position
                namePart.GetNode().Position = position.NewTokenPosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(namePart, token.Start, $3.Tokens)

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
                $$ = &ast.StmtHaltCompiler{ast.Node{}, }

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.HaltCompiller, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.OpenParenthesisToken, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.CloseParenthesisToken, $4.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE namespace_name ';'
            {
                name := &ast.NameName{ast.Node{}, $2}
                $$ = &ast.StmtNamespace{ast.Node{}, name, nil}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, token.End, $3.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
            {
                name := &ast.NameName{ast.Node{}, $2}
                $$ = &ast.StmtNamespace{ast.Node{}, name, $4}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $5)

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
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Namespace, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_declarations ';'
            {
                $$ = &ast.StmtUseList{ast.Node{}, nil, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.UseDeclarationList, $3.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE T_FUNCTION use_function_declarations ';'
            {
                useType := &ast.Identifier{ast.Node{}, $2.Value}
                $$ = &ast.StmtUseList{ast.Node{}, useType, $3}

                // save position
                useType.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(useType, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.UseDeclarationList, $4.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE T_CONST use_const_declarations ';'
            {
                useType := &ast.Identifier{ast.Node{}, $2.Value}
                $$ = &ast.StmtUseList{ast.Node{}, useType, $3}

                // save position
                useType.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(useType, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.UseDeclarationList, $4.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   constant_declaration ';'
            {
                $$ = $1

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $2.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.Tokens)

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

use_declaration:
        namespace_name
            {
                name := &ast.NameName{ast.Node{}, $1}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, nil}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($1)
                $$.GetNode().Position = position.NewNodeListPosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_AS T_STRING
            {
                name := &ast.NameName{ast.Node{}, $1}
                alias := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, alias}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($1)
                alias.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeListTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)
                yylex.(*Parser).setFreeFloating(name, token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := &ast.NameName{ast.Node{}, $2}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, nil}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($2)
                $$.GetNode().Position = position.NewNodeListPosition($2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Slash, $1.Tokens)
                yylex.(*Parser).MoveFreeFloating($2[0], name)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                name := &ast.NameName{ast.Node{}, $2}
                alias := &ast.Identifier{ast.Node{}, $4.Value}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, alias}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($2)
                alias.GetNode().Position = position.NewTokenPosition($4)
                $$.GetNode().Position = position.NewNodeListTokenPosition($2, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Slash, $1.Tokens)
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, token.End, $3.Tokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_function_declarations:
        use_function_declarations ',' use_function_declaration
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_function_declaration
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_function_declaration:
        namespace_name
            {
                name := &ast.NameName{ast.Node{}, $1}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, nil}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($1)
                $$.GetNode().Position = position.NewNodeListPosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_AS T_STRING
            {
                name := &ast.NameName{ast.Node{}, $1}
                alias := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, alias}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($1)
                alias.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeListTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)
                yylex.(*Parser).setFreeFloating(name, token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := &ast.NameName{ast.Node{}, $2}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, nil}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($2)
                $$.GetNode().Position = position.NewNodeListPosition($2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Slash, $1.Tokens)
                yylex.(*Parser).MoveFreeFloating($2[0], name)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                name := &ast.NameName{ast.Node{}, $2}
                alias := &ast.Identifier{ast.Node{}, $4.Value}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, alias}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($2)
                alias.GetNode().Position = position.NewTokenPosition($4)
                $$.GetNode().Position = position.NewNodeListTokenPosition($2, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Slash, $1.Tokens)
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, token.End, $3.Tokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_const_declarations:
        use_const_declarations ',' use_const_declaration
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_const_declaration
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_const_declaration:
        namespace_name
            {
                name := &ast.NameName{ast.Node{}, $1}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, nil}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($1)
                $$.GetNode().Position = position.NewNodeListPosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_AS T_STRING
            {
                name := &ast.NameName{ast.Node{}, $1}
                alias := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, alias}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($1)
                alias.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeListTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)
                yylex.(*Parser).setFreeFloating(name, token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := &ast.NameName{ast.Node{}, $2}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, nil}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($2)
                $$.GetNode().Position = position.NewNodeListPosition($2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Slash, $1.Tokens)
                yylex.(*Parser).MoveFreeFloating($2[0], name)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                name := &ast.NameName{ast.Node{}, $2}
                alias := &ast.Identifier{ast.Node{}, $4.Value}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, alias}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($2)
                alias.GetNode().Position = position.NewTokenPosition($4)
                $$.GetNode().Position = position.NewNodeListTokenPosition($2, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Slash, $1.Tokens)
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, token.End, $3.Tokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

constant_declaration:
        constant_declaration ',' T_STRING '=' static_scalar
            {
                name := &ast.Identifier{ast.Node{}, $3.Value}
                constant := &ast.StmtConstant{ast.Node{}, name, $5}
                constList := $1.(*ast.StmtConstList)
                lastConst := lastNode(constList.Consts)
                constList.Consts = append(constList.Consts, constant)
                $$ = $1

                // save position
                name.GetNode().Position = position.NewTokenPosition($3)
                constant.GetNode().Position = position.NewTokenNodePosition($3, $5)
                $$.GetNode().Position = position.NewNodeNodeListPosition($1, constList.Consts)

                // save comments
                yylex.(*Parser).setFreeFloating(lastConst, token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(constant, token.Start, $3.Tokens)
                yylex.(*Parser).setFreeFloating(constant, token.Name, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST T_STRING '=' static_scalar
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                constant := &ast.StmtConstant{ast.Node{}, name, $4}
                constList := []ast.Vertex{constant}
                $$ = &ast.StmtConstList{ast.Node{}, constList}

                // save position
                name.GetNode().Position = position.NewTokenPosition($2)
                constant.GetNode().Position = position.NewTokenNodePosition($2, $4)
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, constList)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(constant, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating(constant, token.Name, $3.Tokens)

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
    |   T_HALT_COMPILER '(' ')' ';'
            {
                $$ = &ast.StmtHaltCompiler{ast.Node{}, }

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.HaltCompiller, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.OpenParenthesisToken, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.CloseParenthesisToken, $4.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.Tokens)

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
                label := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.StmtLabel{ast.Node{}, label}

                // save position
                label.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Label, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unticked_statement:
        '{' inner_statement_list '}'
            {
                $$ = &ast.StmtStmtList{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_IF parenthesis_expr statement elseif_list else_single
            {
                $$ = &ast.StmtIf{ast.Node{}, $2, $3, $4, $5}

                // save position
                if $5 != nil {
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $5)
                } else if len($4) > 0 {
                    $$.GetNode().Position = position.NewTokenNodeListPosition($1, $4)
                } else {
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $3)
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                if len($2.GetNode().Tokens[token.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.If, $2.GetNode().Tokens[token.OpenParenthesisToken][:len($2.GetNode().Tokens[token.OpenParenthesisToken])-1]); delete($2.GetNode().Tokens, token.OpenParenthesisToken)
                }
                if len($2.GetNode().Tokens[token.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $2.GetNode().Tokens[token.CloseParenthesisToken][:len($2.GetNode().Tokens[token.CloseParenthesisToken])-1]); delete($2.GetNode().Tokens, token.CloseParenthesisToken)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_IF parenthesis_expr ':' inner_statement_list new_elseif_list new_else_single T_ENDIF ';'
            {
                stmts := &ast.StmtStmtList{ast.Node{}, $4}
                $$ = &ast.StmtAltIf{ast.Node{}, $2, stmts, $5, $6}

                // save position
                stmts.GetNode().Position = position.NewNodeListPosition($4)
                $$.GetNode().Position = position.NewTokensPosition($1, $8)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                if len($2.GetNode().Tokens[token.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.If, $2.GetNode().Tokens[token.OpenParenthesisToken][:len($2.GetNode().Tokens[token.OpenParenthesisToken])-1]); delete($2.GetNode().Tokens, token.OpenParenthesisToken)
                }
                if len($2.GetNode().Tokens[token.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $2.GetNode().Tokens[token.CloseParenthesisToken][:len($2.GetNode().Tokens[token.CloseParenthesisToken])-1]); delete($2.GetNode().Tokens, token.CloseParenthesisToken)
                }
                yylex.(*Parser).setFreeFloating($$, token.Cond, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $7.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $8.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $8.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_WHILE parenthesis_expr while_statement
            {
                switch n := $3.(type) {
                case *ast.StmtWhile :
                    n.Cond = $2
                case *ast.StmtAltWhile :
                    n.Cond = $2
                }

                $$ = $3

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                if len($2.GetNode().Tokens[token.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.While, $2.GetNode().Tokens[token.OpenParenthesisToken][:len($2.GetNode().Tokens[token.OpenParenthesisToken])-1]); delete($2.GetNode().Tokens, token.OpenParenthesisToken)
                }
                if len($2.GetNode().Tokens[token.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $2.GetNode().Tokens[token.CloseParenthesisToken][:len($2.GetNode().Tokens[token.CloseParenthesisToken])-1]); delete($2.GetNode().Tokens, token.CloseParenthesisToken)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DO statement T_WHILE parenthesis_expr ';'
            {
                $$ = &ast.StmtDo{ast.Node{}, $2, $4}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)
                if len($4.GetNode().Tokens[token.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.While, $4.GetNode().Tokens[token.OpenParenthesisToken][:len($4.GetNode().Tokens[token.OpenParenthesisToken])-1]); delete($4.GetNode().Tokens, token.OpenParenthesisToken)
                }
                if len($4.GetNode().Tokens[token.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.GetNode().Tokens[token.CloseParenthesisToken][:len($4.GetNode().Tokens[token.CloseParenthesisToken])-1]); delete($4.GetNode().Tokens, token.CloseParenthesisToken)
                }
                yylex.(*Parser).setFreeFloating($$, token.Cond, $5.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $5.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOR '(' for_expr ';' for_expr ';' for_expr ')' for_statement
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
                $$.GetNode().Position = position.NewTokenNodePosition($1, $9)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.For, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.InitExpr, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.CondExpr, $6.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.IncExpr, $8.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_SWITCH parenthesis_expr switch_case_list
            {
                switch n := $3.(type) {
                case *ast.StmtSwitch:
                    n.Cond = $2
                case *ast.StmtAltSwitch:
                    n.Cond = $2
                default:
                    panic("unexpected node type")
                }

                $$ = $3

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                if len($2.GetNode().Tokens[token.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Switch, $2.GetNode().Tokens[token.OpenParenthesisToken][:len($2.GetNode().Tokens[token.OpenParenthesisToken])-1]); delete($2.GetNode().Tokens, token.OpenParenthesisToken)
                }
                if len($2.GetNode().Tokens[token.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $2.GetNode().Tokens[token.CloseParenthesisToken][:len($2.GetNode().Tokens[token.CloseParenthesisToken])-1]); delete($2.GetNode().Tokens, token.CloseParenthesisToken)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BREAK ';'
            {
                $$ = &ast.StmtBreak{ast.Node{}, nil}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BREAK expr ';'
            {
                $$ = &ast.StmtBreak{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONTINUE ';'
            {
                $$ = &ast.StmtContinue{ast.Node{}, nil}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONTINUE expr ';'
            {
                $$ = &ast.StmtContinue{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_RETURN ';'
            {
                $$ = &ast.StmtReturn{ast.Node{}, nil}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_RETURN expr_without_variable ';'
            {
                $$ = &ast.StmtReturn{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_RETURN variable ';'
            {
                $$ = &ast.StmtReturn{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   yield_expr ';'
            {
                $$ = &ast.StmtExpression{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_GLOBAL global_var_list ';'
            {
                $$ = &ast.StmtGlobal{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.VarList, $3.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC static_var_list ';'
            {
                $$ = &ast.StmtStatic{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.VarList, $3.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ECHO echo_expr_list ';'
            {
                $$ = &ast.StmtEcho{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Echo, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INLINE_HTML
            {
                $$ = &ast.StmtInlineHtml{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr ';'
            {
                $$ = &ast.StmtExpression{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_UNSET '(' unset_variables ')' ';'
            {
                $$ = &ast.StmtUnset{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Unset, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.VarList, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.CloseParenthesisToken, $5.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $5.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOREACH '(' variable T_AS foreach_variable foreach_optional_arg ')' foreach_statement
            {
                if $6 == nil {
                    switch n := $8.(type) {
                    case *ast.StmtForeach :
                        n.Expr = $3
                        n.Var = $5
                    case *ast.StmtAltForeach :
                        n.Expr = $3
                        n.Var = $5
                    }
                } else {
                    switch n := $8.(type) {
                    case *ast.StmtForeach :
                        n.Expr = $3
                        n.Key = $5
                        n.Var = $6
                    case *ast.StmtAltForeach :
                        n.Expr = $3
                        n.Key = $5
                        n.Var = $6
                    }
                }

                $$ = $8

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $8)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Foreach, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.Tokens)
                if $6 != nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Key, $6.GetNode().Tokens[token.Key]); delete($6.GetNode().Tokens, token.Key)
                }
                yylex.(*Parser).setFreeFloating($$, token.Var, $7.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOREACH '(' expr_without_variable T_AS foreach_variable foreach_optional_arg ')' foreach_statement
            {
                if $6 == nil {
                    switch n := $8.(type) {
                    case *ast.StmtForeach :
                        n.Expr = $3
                        n.Var = $5
                    case *ast.StmtAltForeach :
                        n.Expr = $3
                        n.Var = $5
                    }
                } else {
                    switch n := $8.(type) {
                    case *ast.StmtForeach :
                        n.Expr = $3
                        n.Key = $5
                        n.Var = $6
                    case *ast.StmtAltForeach :
                        n.Expr = $3
                        n.Key = $5
                        n.Var = $6
                    }
                }

                // save position
                $$ = $8

                $$.GetNode().Position = position.NewTokenNodePosition($1, $8)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Foreach, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.Tokens)
                if $6 != nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Key, $6.GetNode().Tokens[token.Key]); delete($6.GetNode().Tokens, token.Key)
                }
                yylex.(*Parser).setFreeFloating($$, token.Var, $7.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DECLARE '(' declare_list ')' declare_statement
            {
                $$ = $5
                $$.(*ast.StmtDeclare).Consts = $3

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $5)

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
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRY '{' inner_statement_list '}' catch_statement finally_statement
            {
                $$ = &ast.StmtTry{ast.Node{}, $3, $5, $6}

                // save position
                if $6 == nil {
                    $$.GetNode().Position = position.NewTokenNodeListPosition($1, $5)
                } else {
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $6)
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
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_GOTO T_STRING ';'
            {
                label := &ast.Identifier{ast.Node{}, $2.Value}
                $$ = &ast.StmtGoto{ast.Node{}, label}

                // save position
                label.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(label, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Label, $3.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

catch_statement:
        /* empty */
            {
                $$ = []ast.Vertex{}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CATCH '(' fully_qualified_class_name T_VARIABLE ')' '{' inner_statement_list '}' additional_catches
            {
                identifier := &ast.Identifier{ast.Node{}, $4.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                catchNode := &ast.StmtCatch{ast.Node{}, []ast.Vertex{$3}, variable, $7}
                $$ = append([]ast.Vertex{catchNode}, $9...)

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($4)
                variable.GetNode().Position = position.NewTokenPosition($4)
                catchNode.GetNode().Position = position.NewTokensPosition($1, $8)

                // save comments
                yylex.(*Parser).setFreeFloating(catchNode, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(catchNode, token.Catch, $2.Tokens)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $4.Tokens)
                yylex.(*Parser).setFreeFloating(catchNode, token.Var, $5.Tokens)
                yylex.(*Parser).setFreeFloating(catchNode, token.Cond, $6.Tokens)
                yylex.(*Parser).setFreeFloating(catchNode, token.Stmts, $8.Tokens)

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
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Finally, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $4.Tokens)

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
                $$ = []ast.Vertex{}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_additional_catches:
        additional_catch
            {
                $$ = []ast.Vertex{$1}

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
                identifier := &ast.Identifier{ast.Node{}, $4.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.StmtCatch{ast.Node{}, []ast.Vertex{$3}, variable, $7}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($4)
                variable.GetNode().Position = position.NewTokenPosition($4)
                $$.GetNode().Position = position.NewTokensPosition($1, $8)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Catch, $2.Tokens)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $5.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $6.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $8.Tokens)

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
                name := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtFunction{ast.Node{}, $2 != nil, name, $5, nil, $8}

                // save position
                name.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewTokensPosition($1, $9)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $2.Tokens)
                    yylex.(*Parser).setFreeFloating(name, token.Start, $3.Tokens)
                } else {
                    yylex.(*Parser).setFreeFloating(name, token.Start, $3.Tokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.Name, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ParamList, $6.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Params, $7.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $9.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unticked_class_declaration_statement:
        class_entry_type T_STRING extends_from implements_list '{' class_statement_list '}'
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                switch n := $1.(type) {
                    case *ast.StmtClass :
                        n.ClassName = name
                        n.Stmts = $6
                        n.Extends = $3
                        n.Implements = $4

                    case *ast.StmtTrait :
                        // TODO: is it possible that trait extend or implement
                        n.TraitName = name
                        n.Stmts = $6
                }
                $$ = $1

                // save position
                name.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $7)

                // save comments
                yylex.(*Parser).setFreeFloating(name, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $5.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $7.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   interface_entry T_STRING interface_extends_list '{' class_statement_list '}'
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                $$ = &ast.StmtInterface{ast.Node{}, name, $3, $5}

                // save position
                name.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(name, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $6.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


class_entry_type:
        T_CLASS
            {
                $$ = &ast.StmtClass{ast.Node{}, nil, nil, nil, nil, nil, nil}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ABSTRACT T_CLASS
            {
                classModifier := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.StmtClass{ast.Node{}, nil, []ast.Vertex{classModifier}, nil, nil, nil, nil}

                // save position
                classModifier.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ModifierList, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRAIT
            {
                $$ = &ast.StmtTrait{ast.Node{}, nil, nil}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINAL T_CLASS
            {
                classModifier := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.StmtClass{ast.Node{}, nil, []ast.Vertex{classModifier}, nil, nil, nil, nil}

                // save position
                classModifier.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ModifierList, $2.Tokens)

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
                $$ = &ast.StmtClassExtends{ast.Node{}, $2};

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                $$ = &ast.StmtInterfaceExtends{ast.Node{}, $2};

                // save position
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $2)

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
    |   T_IMPLEMENTS interface_list
            {
                $$ = &ast.StmtClassImplements{ast.Node{}, $2};

                // save position
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

interface_list:
        fully_qualified_class_name
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   interface_list ',' fully_qualified_class_name
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

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
                yylex.(*Parser).setFreeFloating($$, token.Key, $1.Tokens)

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
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' assignment_list ')'
            {
                $$ = &ast.ExprList{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.List, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

for_statement:
        statement
            {
                $$ = &ast.StmtFor{ast.Node{}, nil, nil, nil, $1}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDFOR ';'
            {
                stmtList := &ast.StmtStmtList{ast.Node{}, $2}
                $$ = &ast.StmtAltFor{ast.Node{}, nil, nil, nil, stmtList}

                // save position
                stmtList.GetNode().Position = position.NewNodeListPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

foreach_statement:
        statement
            {
                $$ = &ast.StmtForeach{ast.Node{}, nil, nil, nil, $1}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDFOREACH ';'
            {
                stmtList := &ast.StmtStmtList{ast.Node{}, $2}
                $$ = &ast.StmtAltForeach{ast.Node{}, nil, nil, nil, stmtList}

                // save position
                stmtList.GetNode().Position = position.NewNodeListPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


declare_statement:
        statement
            {
                $$ = &ast.StmtDeclare{ast.Node{}, false, nil, $1}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDDECLARE ';'
            {
                stmtList := &ast.StmtStmtList{ast.Node{}, $2}
                $$ = &ast.StmtDeclare{ast.Node{}, true, nil, stmtList}

                // save position
                stmtList.GetNode().Position = position.NewNodeListPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


declare_list:
        T_STRING '=' static_scalar
            {
                name := &ast.Identifier{ast.Node{}, $1.Value}
                constant := &ast.StmtConstant{ast.Node{}, name, $3}
                $$ = []ast.Vertex{constant}

                // save position
                name.GetNode().Position = position.NewTokenPosition($1)
                constant.GetNode().Position = position.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(constant, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(constant, token.Name, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   declare_list ',' T_STRING '=' static_scalar
            {
                name := &ast.Identifier{ast.Node{}, $3.Value}
                constant := &ast.StmtConstant{ast.Node{}, name, $5}
                $$ = append($1, constant)

                // save position
                name.GetNode().Position = position.NewTokenPosition($3)
                constant.GetNode().Position = position.NewTokenNodePosition($3, $5)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(constant, token.Start, $3.Tokens)
                yylex.(*Parser).setFreeFloating(constant, token.Name, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


switch_case_list:
        '{' case_list '}'
            {
                caseList := &ast.StmtCaseList{ast.Node{}, $2}
                $$ = &ast.StmtSwitch{ast.Node{}, nil, caseList}

                // save position
                caseList.GetNode().Position = position.NewTokensPosition($1, $3)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

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
                caseList.GetNode().Position = position.NewTokensPosition($1, $4)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating(caseList, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens(caseList, token.CaseListStart, $2.Tokens)
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListEnd, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' case_list T_ENDSWITCH ';'
            {
                caseList := &ast.StmtCaseList{ast.Node{}, $2}
                $$ = &ast.StmtAltSwitch{ast.Node{}, nil, caseList}

                // save position
                caseList.GetNode().Position = position.NewNodeListPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.Tokens)
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListEnd, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' ';' case_list T_ENDSWITCH ';'
            {

                caseList := &ast.StmtCaseList{ast.Node{}, $3}
                $$ = &ast.StmtAltSwitch{ast.Node{}, nil, caseList}

                // save position
                caseList.GetNode().Position = position.NewNodeListPosition($3)
                $$.GetNode().Position = position.NewTokensPosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens(caseList, token.CaseListStart, $2.Tokens)
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListEnd, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $5.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $5.Tokens)

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
                _case.GetNode().Position = position.NewTokenNodeListPosition($2, $5)

                // save comments
                yylex.(*Parser).setFreeFloating(_case, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating(_case, token.Expr, $4.Tokens)
                yylex.(*Parser).setToken(_case, token.CaseSeparator, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                _default := &ast.StmtDefault{ast.Node{}, $4}
                $$ = append($1, _default)

                // save position
                _default.GetNode().Position = position.NewTokenNodeListPosition($2, $4)

                // save comments
                yylex.(*Parser).setFreeFloating(_default, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating(_default, token.Default, $3.Tokens)
                yylex.(*Parser).setToken(_default, token.CaseSeparator, $3.Tokens)

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
                $$.GetNode().Position = position.NewNodePosition($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDWHILE ';'
            {
                stmtList := &ast.StmtStmtList{ast.Node{}, $2}
                $$ = &ast.StmtAltWhile{ast.Node{}, nil, stmtList}

                // save position
                stmtList.GetNode().Position = position.NewNodeListPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.Tokens)

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
                _elseIf := &ast.StmtElseIf{ast.Node{}, $3, $4}
                $$ = append($1, _elseIf)

                // save position
                _elseIf.GetNode().Position = position.NewTokenNodePosition($2, $4)

                // save comments
                yylex.(*Parser).setFreeFloating(_elseIf, token.Start, $2.Tokens)
                if len($3.GetNode().Tokens[token.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens(_elseIf, token.ElseIf, $3.GetNode().Tokens[token.OpenParenthesisToken][:len($3.GetNode().Tokens[token.OpenParenthesisToken])-1]); delete($3.GetNode().Tokens, token.OpenParenthesisToken)
                }
                if len($3.GetNode().Tokens[token.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens(_elseIf, token.Expr, $3.GetNode().Tokens[token.CloseParenthesisToken][:len($3.GetNode().Tokens[token.CloseParenthesisToken])-1]); delete($3.GetNode().Tokens, token.CloseParenthesisToken)
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
                stmts := &ast.StmtStmtList{ast.Node{}, $5}
                _elseIf := &ast.StmtAltElseIf{ast.Node{}, $3, stmts}
                $$ = append($1, _elseIf)

                // save position
                stmts.GetNode().Position = position.NewNodeListPosition($5)
                _elseIf.GetNode().Position = position.NewTokenNodeListPosition($2, $5)

                // save comments
                yylex.(*Parser).setFreeFloating(_elseIf, token.Start, $2.Tokens)
                if len($3.GetNode().Tokens[token.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens(_elseIf, token.ElseIf, $3.GetNode().Tokens[token.OpenParenthesisToken][:len($3.GetNode().Tokens[token.OpenParenthesisToken])-1]); delete($3.GetNode().Tokens, token.OpenParenthesisToken)
                }
                if len($3.GetNode().Tokens[token.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloatingTokens(_elseIf, token.Expr, $3.GetNode().Tokens[token.CloseParenthesisToken][:len($3.GetNode().Tokens[token.CloseParenthesisToken])-1]); delete($3.GetNode().Tokens, token.CloseParenthesisToken)
                }
                yylex.(*Parser).setFreeFloating(_elseIf, token.Cond, $4.Tokens)

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
                $$ = &ast.StmtElse{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                stmts := &ast.StmtStmtList{ast.Node{}, $3}
                $$ = &ast.StmtAltElse{ast.Node{}, stmts}

                // save position
                stmts.GetNode().Position = position.NewNodeListPosition($3)
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Else, $2.Tokens)

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
        optional_class_type is_reference is_variadic T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $4.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.Parameter{ast.Node{}, $2 != nil, $3 != nil, $1, variable, nil}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($4)
                variable.GetNode().Position = position.NewTokenPosition($4)
                if $1 != nil {
                    $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)
                } else if $2 != nil {
                    $$.GetNode().Position = position.NewTokensPosition($2, $4)
                } else if $3 != nil {
                    $$.GetNode().Position = position.NewTokensPosition($3, $4)
                } else {
                    $$.GetNode().Position = position.NewTokenPosition($4)
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
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Ampersand, $$.GetNode().Tokens[token.Variadic]); delete($$.GetNode().Tokens, token.Variadic)
                }
                if $2 == nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.OptionalType, $$.GetNode().Tokens[token.Ampersand]); delete($$.GetNode().Tokens, token.Ampersand)
                }
                if $1 == nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Start, $$.GetNode().Tokens[token.OptionalType]); delete($$.GetNode().Tokens, token.OptionalType)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   optional_class_type is_reference is_variadic T_VARIABLE '=' static_scalar
            {
                identifier := &ast.Identifier{ast.Node{}, $4.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.Parameter{ast.Node{}, $2 != nil, $3 != nil, $1, variable, $6}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($4)
                variable.GetNode().Position = position.NewTokenPosition($4)
                if $1 != nil {
                    $$.GetNode().Position = position.NewNodesPosition($1, $6)
                } else if $2 != nil {
                    $$.GetNode().Position = position.NewTokenNodePosition($2, $6)
                } else if $3 != nil {
                    $$.GetNode().Position = position.NewTokenNodePosition($3, $6)
                } else {
                    $$.GetNode().Position = position.NewTokenNodePosition($4, $6)
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
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Ampersand, $$.GetNode().Tokens[token.Variadic]); delete($$.GetNode().Tokens, token.Variadic)
                }
                if $2 == nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.OptionalType, $$.GetNode().Tokens[token.Ampersand]); delete($$.GetNode().Tokens, token.Ampersand)
                }
                if $1 == nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Start, $$.GetNode().Tokens[token.OptionalType]); delete($$.GetNode().Tokens, token.OptionalType)
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
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CALLABLE
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                $$ = &ast.ArgumentList{ast.Node{}, nil}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ArgumentList, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' non_empty_function_call_parameter_list ')'
            {
                $$ = &ast.ArgumentList{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ArgumentList, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' yield_expr ')'
            {
                arg := &ast.Argument{ast.Node{}, false, false, $2}
                $$ = &ast.ArgumentList{ast.Node{}, []ast.Vertex{arg}}

                // save position
                arg.GetNode().Position = position.NewNodePosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ArgumentList, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


non_empty_function_call_parameter_list:
        function_call_parameter
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_function_call_parameter_list ',' function_call_parameter
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

function_call_parameter:
        expr_without_variable
            {
                $$ = &ast.Argument{ast.Node{}, false, false, $1}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable
            {
                $$ = &ast.Argument{ast.Node{}, false, false, $1}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' w_variable
            {
                $$ = &ast.Argument{ast.Node{}, false, true, $2}

                // save position
                $$.GetNode().Position = position.NewNodePosition($2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ELLIPSIS expr
            {
                $$ = &ast.Argument{ast.Node{}, true, false, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

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
        T_VARIABLE
            {
                name := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.ExprVariable{ast.Node{}, name}

                // save position
                name.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' r_variable
            {
                $$ = &ast.ExprVariable{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' '{' expr '}'
            {
                $$ = &ast.ExprVariable{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($3, token.Start, append($2.Tokens, $3.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($3, token.End, append($3.GetNode().Tokens[token.End], $4.Tokens...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


static_var_list:
        static_var_list ',' T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $3.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                staticVar := &ast.StmtStaticVar{ast.Node{}, variable, nil}
                $$ = append($1, staticVar)

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($3)
                variable.GetNode().Position = position.NewTokenPosition($3)
                staticVar.GetNode().Position = position.NewTokenPosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(staticVar, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_var_list ',' T_VARIABLE '=' static_scalar
            {
                identifier := &ast.Identifier{ast.Node{}, $3.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                staticVar := &ast.StmtStaticVar{ast.Node{}, variable, $5}
                $$ = append($1, staticVar)

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($3)
                variable.GetNode().Position = position.NewTokenPosition($3)
                staticVar.GetNode().Position = position.NewTokenNodePosition($3, $5)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(staticVar, token.Start, $3.Tokens)
                yylex.(*Parser).setFreeFloating(staticVar, token.Var, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                staticVar := &ast.StmtStaticVar{ast.Node{}, variable, nil}
                $$ = []ast.Vertex{staticVar}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)
                staticVar.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(staticVar, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' static_scalar
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                staticVar := &ast.StmtStaticVar{ast.Node{}, variable, $3}
                $$ = []ast.Vertex{staticVar}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)
                staticVar.GetNode().Position = position.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(staticVar, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(staticVar, token.Var, $2.Tokens)

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
        variable_modifiers class_variable_declaration ';'
            {
                $$ = &ast.StmtPropertyList{ast.Node{}, $1, nil, $2}

                // save position
                $$.GetNode().Position = position.NewNodeListTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)
                yylex.(*Parser).setFreeFloating($$, token.PropertyList, $3.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_constant_declaration ';'
            {
                $$ = $1

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.ConstList, $2.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_use_statement
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   method_modifiers function is_reference T_STRING '(' parameter_list ')' method_body
            {
                name := &ast.Identifier{ast.Node{}, $4.Value}
                $$ = &ast.StmtClassMethod{ast.Node{}, $3 != nil, name, $1, $6, nil, $8}

                // save position
                name.GetNode().Position = position.NewTokenPosition($4)
                if $1 == nil {
                    $$.GetNode().Position = position.NewTokenNodePosition($2, $8)
                } else {
                    $$.GetNode().Position = position.NewNodeListNodePosition($1, $8)
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
                yylex.(*Parser).setFreeFloating($$, token.Name, $5.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ParameterList, $7.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_use_statement:
        T_USE trait_list trait_adaptations
            {
                $$ = &ast.StmtTraitUse{ast.Node{}, $2, $3}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_list:
        fully_qualified_class_name
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_list ',' fully_qualified_class_name
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

                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' trait_adaptation_list '}'
            {
                $$ = &ast.StmtTraitAdaptationList{ast.Node{}, $2}

                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AdaptationList, $3.Tokens)

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
                $$ = []ast.Vertex{$1}

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
                yylex.(*Parser).setFreeFloating($$, token.NameList, $2.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_alias ';'
            {
                $$ = $1;

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Alias, $2.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_precedence:
        trait_method_reference_fully_qualified T_INSTEADOF trait_reference_list
            {
                $$ = &ast.StmtTraitUsePrecedence{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeNodeListPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_reference_list:
        fully_qualified_class_name
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_reference_list ',' fully_qualified_class_name
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_method_reference:
        T_STRING
            {
                name := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.StmtTraitMethodRef{ast.Node{}, nil, name}

                // save position
                name.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                target := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtTraitMethodRef{ast.Node{}, $1, target}

                // save position
                target.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)
                yylex.(*Parser).setFreeFloating(target, token.Start, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_alias:
        trait_method_reference T_AS trait_modifiers T_STRING
            {
                alias := &ast.Identifier{ast.Node{}, $4.Value}
                $$ = &ast.StmtTraitUseAlias{ast.Node{}, $1, $3, alias}

                // save position
                alias.GetNode().Position = position.NewTokenPosition($4)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

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
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.Tokens)

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
                $$ = &ast.StmtNop{ast.Node{}, }

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' inner_statement_list '}'
            {
                $$ = &ast.StmtStmtList{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)

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
                modifier := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = []ast.Vertex{modifier}

                // save position
                modifier.GetNode().Position = position.NewTokenPosition($1)

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
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PROTECTED
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PRIVATE
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ABSTRACT
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINAL
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_variable_declaration:
        class_variable_declaration ',' T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $3.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                property := &ast.StmtProperty{ast.Node{}, variable, nil}
                $$ = append($1, property)

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($3)
                variable.GetNode().Position = position.NewTokenPosition($3)
                property.GetNode().Position = position.NewTokenPosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(property, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_variable_declaration ',' T_VARIABLE '=' static_scalar
            {
                identifier := &ast.Identifier{ast.Node{}, $3.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                property := &ast.StmtProperty{ast.Node{}, variable, $5}
                $$ = append($1, property)

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($3)
                variable.GetNode().Position = position.NewTokenPosition($3)
                property.GetNode().Position = position.NewTokenNodePosition($3, $5)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(property, token.Start, $3.Tokens)
                yylex.(*Parser).setFreeFloating(property, token.Var, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                property := &ast.StmtProperty{ast.Node{}, variable, nil}
                $$ = []ast.Vertex{property}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)
                property.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(property, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' static_scalar
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                property := &ast.StmtProperty{ast.Node{}, variable, $3}
                $$ = []ast.Vertex{property}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)
                property.GetNode().Position = position.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(property, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating(property, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_constant_declaration:
        class_constant_declaration ',' T_STRING '=' static_scalar
            {
                name := &ast.Identifier{ast.Node{}, $3.Value}
                constant := &ast.StmtConstant{ast.Node{}, name, $5}
                constList := $1.(*ast.StmtClassConstList)
                lastConst := lastNode(constList.Consts)
                constList.Consts = append(constList.Consts, constant)
                $$ = $1

                // save position
                name.GetNode().Position = position.NewTokenPosition($3)
                constant.GetNode().Position = position.NewTokenNodePosition($3, $5)
                $1.GetNode().Position = position.NewNodesPosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating(lastConst, token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(constant, token.Start, $3.Tokens)
                yylex.(*Parser).setFreeFloating(constant, token.Name, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST T_STRING '=' static_scalar
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                constant := &ast.StmtConstant{ast.Node{}, name, $4}
                $$ = &ast.StmtClassConstList{ast.Node{}, nil, []ast.Vertex{constant}}

                // save position
                name.GetNode().Position = position.NewTokenPosition($2)
                constant.GetNode().Position = position.NewTokenNodePosition($2, $4)
                $$.GetNode().Position = position.NewTokenNodePosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(constant, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating(constant, token.Name, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

echo_expr_list:
        echo_expr_list ',' expr
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
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                $$ = []ast.Vertex{$1}

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
                fetch := &ast.ExprArrayDimFetch{ast.Node{}, nil, $3}
                $$ = append($1, fetch)

                // save position
                fetch.GetNode().Position = position.NewNodePosition($3)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' dim_offset ']'
            {
                fetch := &ast.ExprArrayDimFetch{ast.Node{}, nil, $2}
                $$ = []ast.Vertex{fetch}

                // save position
                fetch.GetNode().Position = position.NewNodePosition($2)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Var, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Expr, $3.Tokens)

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
                    $$ = &ast.ExprNew{ast.Node{}, $2, $3.(*ast.ArgumentList)}
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $3)
                } else {
                    $$ = &ast.ExprNew{ast.Node{}, $2, nil}
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $2)
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

expr_without_variable:
        T_LIST '(' assignment_list ')' '=' expr
            {
                listNode := &ast.ExprList{ast.Node{}, $3}
                $$ = &ast.ExprAssign{ast.Node{}, listNode, $6}

                // save position
                listNode.GetNode().Position = position.NewTokensPosition($1, $4)
                $$.GetNode().Position = position.NewTokenNodePosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(listNode, token.List, $2.Tokens)
                yylex.(*Parser).setFreeFloating(listNode, token.ArrayPairList, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $5.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' expr
            {
                $$ = &ast.ExprAssign{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' '&' variable
            {
                $$ = &ast.ExprAssignReference{ast.Node{}, $1, $4}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Equal, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' '&' T_NEW class_name_reference ctor_arguments
            {
                var _new *ast.ExprNew

                if $6 != nil {
                    _new = &ast.ExprNew{ast.Node{}, $5, $6.(*ast.ArgumentList)}
                } else {
                    _new = &ast.ExprNew{ast.Node{}, $5, nil}
                }
                $$ = &ast.ExprAssignReference{ast.Node{}, $1, _new}

                // save position
                if $6 != nil {
                    _new.GetNode().Position = position.NewTokenNodePosition($4, $6)
                } else {
                    _new.GetNode().Position = position.NewTokenNodePosition($4, $5)
                }
                $$.GetNode().Position = position.NewNodesPosition($1, _new)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Equal, $3.Tokens)
                yylex.(*Parser).setFreeFloating(_new, token.Start, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLONE expr
            {
                $$ = &ast.ExprClone{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_PLUS_EQUAL expr
            {
                $$ = &ast.ExprAssignPlus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MINUS_EQUAL expr
            {
                $$ = &ast.ExprAssignMinus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MUL_EQUAL expr
            {
                $$ = &ast.ExprAssignMul{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_POW_EQUAL expr
            {
                $$ = &ast.ExprAssignPow{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_DIV_EQUAL expr
            {
                $$ = &ast.ExprAssignDiv{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_CONCAT_EQUAL expr
            {
                $$ = &ast.ExprAssignConcat{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MOD_EQUAL expr
            {
                $$ = &ast.ExprAssignMod{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_AND_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_OR_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_XOR_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_SL_EQUAL expr
            {
                $$ = &ast.ExprAssignShiftLeft{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_SR_EQUAL expr
            {
                $$ = &ast.ExprAssignShiftRight{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   rw_variable T_INC
            {
                $$ = &ast.ExprPostInc{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INC rw_variable
            {
                $$ = &ast.ExprPreInc{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   rw_variable T_DEC
            {
                $$ = &ast.ExprPostDec{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DEC rw_variable
            {
                $$ = &ast.ExprPreDec{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_BOOLEAN_OR expr
            {
                $$ = &ast.ExprBinaryBooleanOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_BOOLEAN_AND expr
            {
                $$ = &ast.ExprBinaryBooleanAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_OR expr
            {
                $$ = &ast.ExprBinaryLogicalOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_AND expr
            {
                $$ = &ast.ExprBinaryLogicalAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_XOR expr
            {
                $$ = &ast.ExprBinaryLogicalXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '|' expr
            {
                $$ = &ast.ExprBinaryBitwiseOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '&' expr
            {
                $$ = &ast.ExprBinaryBitwiseAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '^' expr
            {
                $$ = &ast.ExprBinaryBitwiseXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '.' expr
            {
                $$ = &ast.ExprBinaryConcat{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '+' expr
            {
                $$ = &ast.ExprBinaryPlus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '-' expr
            {
                $$ = &ast.ExprBinaryMinus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '*' expr
            {
                $$ = &ast.ExprBinaryMul{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_POW expr
            {
                $$ = &ast.ExprBinaryPow{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '/' expr
            {
                $$ = &ast.ExprBinaryDiv{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '%' expr
            {
                $$ = &ast.ExprBinaryMod{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SL expr
            {
                $$ = &ast.ExprBinaryShiftLeft{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SR expr
            {
                $$ = &ast.ExprBinaryShiftRight{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '+' expr %prec T_INC
            {
                $$ = &ast.ExprUnaryPlus{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '-' expr %prec T_INC
            {
                $$ = &ast.ExprUnaryMinus{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '!' expr
            {
                $$ = &ast.ExprBooleanNot{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '~' expr
            {
                $$ = &ast.ExprBitwiseNot{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_IDENTICAL expr
            {
                $$ = &ast.ExprBinaryIdentical{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_NOT_IDENTICAL expr
            {
                $$ = &ast.ExprBinaryNotIdentical{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_EQUAL expr
            {
                $$ = &ast.ExprBinaryEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_NOT_EQUAL expr
            {
                $$ = &ast.ExprBinaryNotEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)
                yylex.(*Parser).setToken($$, token.Equal, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '<' expr
            {
                $$ = &ast.ExprBinarySmaller{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_SMALLER_OR_EQUAL expr
            {
                $$ = &ast.ExprBinarySmallerOrEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '>' expr
            {
                $$ = &ast.ExprBinaryGreater{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_GREATER_OR_EQUAL expr
            {
                $$ = &ast.ExprBinaryGreaterOrEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_INSTANCEOF class_name_reference
            {
                $$ = &ast.ExprInstanceOf{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   parenthesis_expr
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)

                yylex.(*Parser).setFreeFloatingTokens($1, token.Start, append($1.GetNode().Tokens[token.OpenParenthesisToken], $1.GetNode().Tokens[token.Start]...)); delete($1.GetNode().Tokens, token.OpenParenthesisToken)
                yylex.(*Parser).setFreeFloatingTokens($1, token.End, append($1.GetNode().Tokens[token.End], $1.GetNode().Tokens[token.CloseParenthesisToken]...)); delete($1.GetNode().Tokens, token.CloseParenthesisToken)
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
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, append($1.Tokens, $$.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, append($$.GetNode().Tokens[token.End], $3.Tokens...))

                for _, n := range($4) {
                    switch nn := n.(type) {
                        case *ast.ExprArrayDimFetch:
                            nn.Var = $$
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)

                        case *ast.ExprPropertyFetch:
                            nn.Var = $$
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)

                        case *ast.ExprMethodCall:
                            nn.Var = $$
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)
                    }

                    // save position
                    $$.GetNode().Position = position.NewNodesPosition($$, n)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '?' expr ':' expr
            {
                $$ = &ast.ExprTernary{ast.Node{}, $1, $3, $5}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $5)

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
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.True, $3.Tokens)

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
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOUBLE_CAST expr
            {
                $$ = &ast.ExprCastDouble{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STRING_CAST expr
            {
                $$ = &ast.ExprCastString{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ARRAY_CAST expr
            {
                $$ = &ast.ExprCastArray{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_OBJECT_CAST expr
            {
                $$ = &ast.ExprCastObject{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BOOL_CAST expr
            {
                $$ = &ast.ExprCastBool{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_UNSET_CAST expr
            {
                $$ = &ast.ExprCastUnset{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXIT exit_expr
            {
                e := $2.(*ast.ExprExit)
                $$ = $2

                if (bytes.EqualFold($1.Value, []byte("die"))) {
                    e.Die = true
                }

                // save position
                if $2.GetNode().Position == nil {
                    $$.GetNode().Position = position.NewTokenPosition($1)
                } else {
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $2)
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '@' expr
            {
                $$ = &ast.ExprErrorSuppress{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                $$ = &ast.ExprShellExec{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PRINT expr
            {
                $$ = &ast.ExprPrint{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD
            {
                $$ = &ast.ExprYield{ast.Node{}, nil, nil}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   function is_reference '(' parameter_list ')' lexical_vars '{' inner_statement_list '}'
            {
                $$ = &ast.ExprClosure{ast.Node{}, $2 != nil, false, $4, $6, nil, $8}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $9)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                if $2 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $3.Tokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $2.Tokens)
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $3.Tokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.ParameterList, $5.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.LexicalVars, $7.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $9.Tokens)

                // normalize
                if $6 == nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Params, $$.GetNode().Tokens[token.LexicalVars]); delete($$.GetNode().Tokens, token.LexicalVars)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC function is_reference '(' parameter_list ')' lexical_vars '{' inner_statement_list '}'
            {
                $$ = &ast.ExprClosure{ast.Node{}, $3 != nil, true, $5, $7, nil, $9}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $10)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Static, $2.Tokens)
                if $3 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $4.Tokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $3.Tokens)
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $4.Tokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.ParameterList, $6.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.LexicalVars, $8.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $10.Tokens)

                // normalize
                if $7 == nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Params, $$.GetNode().Tokens[token.LexicalVars]); delete($$.GetNode().Tokens, token.LexicalVars)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

yield_expr:
        T_YIELD expr_without_variable
            {
                $$ = &ast.ExprYield{ast.Node{}, nil, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD variable
            {
                $$ = &ast.ExprYield{ast.Node{}, nil, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD expr T_DOUBLE_ARROW expr_without_variable
            {
                $$ = &ast.ExprYield{ast.Node{}, $2, $4}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD expr T_DOUBLE_ARROW variable
            {
                $$ = &ast.ExprYield{ast.Node{}, $2, $4}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

combined_scalar_offset:
        combined_scalar '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   combined_scalar_offset '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONSTANT_ENCAPSED_STRING '[' dim_offset ']'
            {
                str := &ast.ScalarString{ast.Node{}, $1.Value}
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, str, $3}

                // save position
                str.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewNodeTokenPosition(str, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   general_constant '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

combined_scalar:
        T_ARRAY '(' array_pair_list ')'
            {
                $$ = &ast.ExprArray{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

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
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $3.Tokens)

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
                $$ = &ast.ExprClosureUse{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Use, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.LexicalVarList, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

lexical_var_list:
        lexical_var_list ',' T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $3.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = append($1, variable)

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($3)
                variable.GetNode().Position = position.NewTokenPosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   lexical_var_list ',' '&' T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $4.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                reference := &ast.ExprReference{ast.Node{}, variable}
                $$ = append($1, reference)

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($4)
                variable.GetNode().Position = position.NewTokenPosition($4)
                reference.GetNode().Position = position.NewTokensPosition($3, $4)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(reference, token.Start, $3.Tokens)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = []ast.Vertex{variable}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(variable, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $2.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                reference := &ast.ExprReference{ast.Node{}, variable}
                $$ = []ast.Vertex{reference}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($2)
                variable.GetNode().Position = position.NewTokenPosition($2)
                reference.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating(reference, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

function_call:
        namespace_name function_call_parameter_list
            {
                name := &ast.NameName{ast.Node{}, $1}
                $$ = &ast.ExprFunctionCall{ast.Node{}, name, $2.(*ast.ArgumentList)}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($1)
                $$.GetNode().Position = position.NewNodesPosition(name, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name function_call_parameter_list
            {
                funcName := &ast.NameRelative{ast.Node{}, $3}
                $$ = &ast.ExprFunctionCall{ast.Node{}, funcName, $4.(*ast.ArgumentList)}

                // save position
                funcName.GetNode().Position = position.NewTokenNodeListPosition($1, $3)
                $$.GetNode().Position = position.NewNodesPosition(funcName, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(funcName, token.Namespace, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name function_call_parameter_list
            {
                funcName := &ast.NameFullyQualified{ast.Node{}, $2}
                $$ = &ast.ExprFunctionCall{ast.Node{}, funcName, $3.(*ast.ArgumentList)}

                // save position
                funcName.GetNode().Position = position.NewTokenNodeListPosition($1, $2)
                $$.GetNode().Position = position.NewNodesPosition(funcName, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM variable_name function_call_parameter_list
            {
                $$ = &ast.ExprStaticCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects function_call_parameter_list
            {
                $$ = &ast.ExprStaticCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_name function_call_parameter_list
            {
                $$ = &ast.ExprStaticCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects function_call_parameter_list
            {
                $$ = &ast.ExprStaticCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_without_objects function_call_parameter_list
            {
                $$ = &ast.ExprFunctionCall{ast.Node{}, $1, $2.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $2)

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
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name
            {
                $$ = &ast.NameName{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeListPosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameRelative{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Namespace, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameFullyQualified{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

fully_qualified_class_name:
        namespace_name
            {
                $$ = &ast.NameName{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeListPosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameRelative{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Namespace, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameFullyQualified{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                yylex.(*Parser).setFreeFloating($3[0], token.Var, $2.Tokens)

                for _, n := range($3) {
                    switch nn := n.(type) {
                        case *ast.ExprArrayDimFetch:
                            nn.Var = $$
                            $$.GetNode().Position = position.NewNodesPosition($$, nn)
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)

                        case *ast.ExprPropertyFetch:
                            nn.Var = $$
                            $$.GetNode().Position = position.NewNodesPosition($$, nn)
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)
                    }
                }

                for _, n := range($4) {
                    switch nn := n.(type) {
                        case *ast.ExprArrayDimFetch:
                            nn.Var = $$
                            $$.GetNode().Position = position.NewNodesPosition($$, nn)
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)

                        case *ast.ExprPropertyFetch:
                            nn.Var = $$
                            $$.GetNode().Position = position.NewNodesPosition($$, nn)
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)
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
                $$ = []ast.Vertex{}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


dynamic_class_name_variable_property:
        T_OBJECT_OPERATOR object_property
            {
                $$ = $2

                // save comments
                yylex.(*Parser).setFreeFloating($2[0], token.Var, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

exit_expr:
        /* empty */
            {
                $$ = &ast.ExprExit{ast.Node{}, false, nil};

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' ')'
            {
                $$ = &ast.ExprExit{ast.Node{}, false, nil};

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Exit, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   parenthesis_expr
            {
                $$ = &ast.ExprExit{ast.Node{}, false, $1};

                // save position
                if bytes.Compare(yylex.(*Parser).currentToken.Value, []byte(")")) == 0 {
                    $$.GetNode().Position = position.NewTokenPosition(yylex.(*Parser).currentToken)
                } else {
                    $$.GetNode().Position = position.NewNodePosition($1)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Exit, $1.GetNode().Tokens[token.OpenParenthesisToken]); delete($1.GetNode().Tokens, token.OpenParenthesisToken)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $1.GetNode().Tokens[token.CloseParenthesisToken]); delete($1.GetNode().Tokens, token.CloseParenthesisToken)
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
                part.GetNode().Position = position.NewTokenPosition($1)

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
                $$ = &ast.ScalarLnumber{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DNUMBER
            {
                $$ = &ast.ScalarDnumber{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONSTANT_ENCAPSED_STRING
            {
                $$ = &ast.ScalarString{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LINE
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FILE
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DIR
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRAIT_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_METHOD_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FUNC_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC
            {
                encapsed := &ast.ScalarEncapsedStringPart{ast.Node{}, $2.Value}
                $$ = &ast.ScalarHeredoc{ast.Node{}, $1.Value, []ast.Vertex{encapsed}}

                // save position
                encapsed.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{ast.Node{}, $1.Value, nil}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_class_constant:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                target.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)
                yylex.(*Parser).setFreeFloating(target, token.Start, $3.Tokens)

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
                name := &ast.NameName{ast.Node{}, $1}
                $$ = &ast.ExprConstFetch{ast.Node{}, name}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($1)
                $$.GetNode().Position = position.NewNodePosition(name)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                name := &ast.NameRelative{ast.Node{}, $3}
                $$ = &ast.ExprConstFetch{ast.Node{}, name}

                // save position
                name.GetNode().Position = position.NewTokenNodeListPosition($1, $3)
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Namespace, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := &ast.NameFullyQualified{ast.Node{}, $2}
                $$ = &ast.ExprConstFetch{ast.Node{}, name}

                // save position
                name.GetNode().Position = position.NewTokenNodeListPosition($1, $2)
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ARRAY '(' static_array_pair_list ')'
            {
                $$ = &ast.ExprArray{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Array, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' static_array_pair_list ']'
            {
                $$ = &ast.ExprShortArray{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_class_constant
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '+' static_scalar_value
            {
                $$ = &ast.ExprBinaryPlus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '-' static_scalar_value
            {
                $$ = &ast.ExprBinaryMinus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '*' static_scalar_value
            {
                $$ = &ast.ExprBinaryMul{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_POW static_scalar_value
            {
                $$ = &ast.ExprBinaryPow{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '/' static_scalar_value
            {
                $$ = &ast.ExprBinaryDiv{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '%' static_scalar_value
            {
                $$ = &ast.ExprBinaryMod{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '!' static_scalar_value
            {
                $$ = &ast.ExprBooleanNot{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '~' static_scalar_value
            {
                $$ = &ast.ExprBitwiseNot{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '|' static_scalar_value
            {
                $$ = &ast.ExprBinaryBitwiseOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '&' static_scalar_value
            {
                $$ = &ast.ExprBinaryBitwiseAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '^' static_scalar_value
            {
                $$ = &ast.ExprBinaryBitwiseXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_SL static_scalar_value
            {
                $$ = &ast.ExprBinaryShiftLeft{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_SR static_scalar_value
            {
                $$ = &ast.ExprBinaryShiftRight{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '.' static_scalar_value
            {
                $$ = &ast.ExprBinaryConcat{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_LOGICAL_XOR static_scalar_value
            {
                $$ = &ast.ExprBinaryLogicalXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_LOGICAL_AND static_scalar_value
            {
                $$ = &ast.ExprBinaryLogicalAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_LOGICAL_OR static_scalar_value
            {
                $$ = &ast.ExprBinaryLogicalOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_BOOLEAN_AND static_scalar_value
            {
                $$ = &ast.ExprBinaryBooleanAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_BOOLEAN_OR static_scalar_value
            {
                $$ = &ast.ExprBinaryBooleanOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_IDENTICAL static_scalar_value
            {
                $$ = &ast.ExprBinaryIdentical{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_NOT_IDENTICAL static_scalar_value
            {
                $$ = &ast.ExprBinaryNotIdentical{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_EQUAL static_scalar_value
            {
                $$ = &ast.ExprBinaryEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_NOT_EQUAL static_scalar_value
            {
                $$ = &ast.ExprBinaryNotEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)
                yylex.(*Parser).setToken($$, token.Equal, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '<' static_scalar_value
            {
                $$ = &ast.ExprBinarySmaller{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '>' static_scalar_value
            {
                $$ = &ast.ExprBinaryGreater{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_SMALLER_OR_EQUAL static_scalar_value
            {
                $$ = &ast.ExprBinarySmallerOrEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_GREATER_OR_EQUAL static_scalar_value
            {
                $$ = &ast.ExprBinaryGreaterOrEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '?' ':' static_scalar_value
            {
                $$ = &ast.ExprTernary{ast.Node{}, $1, nil, $4}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.True, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '?' static_scalar_value ':' static_scalar_value
            {
                $$ = &ast.ExprTernary{ast.Node{}, $1, $3, $5}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $5)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.True, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '+' static_scalar_value
            {
                $$ = &ast.ExprUnaryPlus{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '-' static_scalar_value
            {
                $$ = &ast.ExprUnaryMinus{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' static_scalar_value ')'
            {
                $$ = $2

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, append($1.Tokens, $$.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, append($$.GetNode().Tokens[token.End], $3.Tokens...))

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
                name := &ast.NameName{ast.Node{}, $1}
                $$ = &ast.ExprConstFetch{ast.Node{}, name}

                // save position
                name.GetNode().Position = position.NewNodeListPosition($1)
                $$.GetNode().Position = position.NewNodePosition(name)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                name := &ast.NameRelative{ast.Node{}, $3}
                $$ = &ast.ExprConstFetch{ast.Node{}, name}

                // save position
                name.GetNode().Position = position.NewTokenNodeListPosition($1, $3)
                $$.GetNode().Position = position.NewNodePosition(name)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(name, token.Namespace, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := &ast.NameFullyQualified{ast.Node{}, $2}
                $$ = &ast.ExprConstFetch{ast.Node{}, name}

                // save position
                name.GetNode().Position = position.NewTokenNodeListPosition($1, $2)
                $$.GetNode().Position = position.NewNodePosition(name)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

scalar:
        T_STRING_VARNAME
            {
                name := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.ExprVariable{ast.Node{}, name}

                // save position
                name.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                $$ = &ast.ScalarEncapsed{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC
            {
                 $$ = &ast.ScalarHeredoc{ast.Node{}, $1.Value, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                    yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
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
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, $3, $5}
                $$ = append($1, arrayItem)

                // save position
                arrayItem.GetNode().Position = position.NewNodesPosition($3, $5)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_static_array_pair_list ',' static_scalar_value
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, nil, $3}
                $$ = append($1, arrayItem)

                // save position
                arrayItem.GetNode().Position = position.NewNodePosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_DOUBLE_ARROW static_scalar_value
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, $1, $3}
                $$ = []ast.Vertex{arrayItem}

                // save position
                arrayItem.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, nil, $1}
                $$ = []ast.Vertex{arrayItem}

                // save position
                arrayItem.GetNode().Position = position.NewNodePosition($1)

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
                if len($2.GetNode().Tokens[token.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($2, token.Start, append($2.GetNode().Tokens[token.OpenParenthesisToken], $2.GetNode().Tokens[token.Start]...))
                }
                if len($2.GetNode().Tokens[token.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($2, token.End, append($2.GetNode().Tokens[token.End], $2.GetNode().Tokens[token.CloseParenthesisToken]...))
                }
                yylex.(*Parser).setFreeFloatingTokens($2, token.OpenParenthesisToken, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($2, token.CloseParenthesisToken, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' yield_expr ')'
            {
                $$ = $2

                // save comments
                if len($2.GetNode().Tokens[token.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($2, token.Start, append($2.GetNode().Tokens[token.OpenParenthesisToken], $2.GetNode().Tokens[token.Start]...))
                }
                if len($2.GetNode().Tokens[token.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($2, token.End, append($2.GetNode().Tokens[token.End], $2.GetNode().Tokens[token.CloseParenthesisToken]...))
                }
                yylex.(*Parser).setFreeFloatingTokens($2, token.OpenParenthesisToken, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($2, token.CloseParenthesisToken, $3.Tokens)

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
                    $4[0].(*ast.ExprMethodCall).Method = $3[len($3)-1].(*ast.ExprPropertyFetch).Property
                    $3 = append($3[:len($3)-1], $4...)
                }

                // save comments
                yylex.(*Parser).setFreeFloating($3[0], token.Var, $2.Tokens)

                for _, n := range($3) {
                    switch nn := n.(type) {
                        case *ast.ExprArrayDimFetch:
                            nn.Var = $$
                            nn.GetNode().Position = position.NewNodesPosition($$, nn)
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)

                        case *ast.ExprPropertyFetch:
                            nn.Var = $$
                            nn.GetNode().Position = position.NewNodesPosition($$, nn)
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)

                        case *ast.ExprMethodCall:
                            nn.Var = $$
                            nn.GetNode().Position = position.NewNodesPosition($$, nn)
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)
                    }
                }

                for _, n := range($5) {
                    switch nn := n.(type) {
                        case *ast.ExprArrayDimFetch:
                            nn.Var = $$
                            nn.GetNode().Position = position.NewNodesPosition($$, nn)
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)

                        case *ast.ExprPropertyFetch:
                            nn.Var = $$
                            nn.GetNode().Position = position.NewNodesPosition($$, nn)
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)

                        case *ast.ExprMethodCall:
                            nn.Var = $$
                            nn.GetNode().Position = position.NewNodesPosition($$, nn)
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Var, $$)
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
                $$ = []ast.Vertex{}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


variable_property:
        T_OBJECT_OPERATOR object_property method_or_not
            {
                if $3 != nil {
                    $3[0].(*ast.ExprMethodCall).Method = $2[len($2)-1].(*ast.ExprPropertyFetch).Property
                    $2 = append($2[:len($2)-1], $3...)
                }

                $$ = $2

                // save comments
                yylex.(*Parser).setFreeFloating($2[0], token.Var, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

array_method_dereference:
        array_method_dereference '[' dim_offset ']'
            {
                fetch := &ast.ExprArrayDimFetch{ast.Node{}, nil, $3}
                $$ = append($1, fetch)

                // save position
                fetch.GetNode().Position = position.NewNodePosition($3)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   method '[' dim_offset ']'
            {
                fetch := &ast.ExprArrayDimFetch{ast.Node{}, nil, $3}
                $$ = []ast.Vertex{$1, fetch}

                // save position
                fetch.GetNode().Position = position.NewNodePosition($3)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method:
        function_call_parameter_list
            {
                $$ = &ast.ExprMethodCall{ast.Node{}, nil, nil, $1.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method_or_not:
        method
            {
                $$ = []ast.Vertex{$1}

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
                $1.last.VarName = $2

                for _, n := range($1.all) {
                    n.GetNode().Position = position.NewNodesPosition(n, $2)
                }

                $$ = $1.all[0]

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects
            {
                $$ = &ast.ExprStaticPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects
            {
                $$ = &ast.ExprStaticPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)

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
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   function_call '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.Tokens)

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
                $1.last.VarName = $2

                for _, n := range($1.all) {
                    n.GetNode().Position = position.NewNodesPosition(n, $2)
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
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   reference_variable '{' expr '}'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.Tokens)

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
                name := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.ExprVariable{ast.Node{}, name}

                // save position
                name.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' '{' expr '}'
            {
                $$ = &ast.ExprVariable{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($3, token.Start, append($2.Tokens, $3.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($3, token.End, append($3.GetNode().Tokens[token.End], $4.Tokens...))

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
                fetch := &ast.ExprPropertyFetch{ast.Node{}, nil, $1}
                $$ = []ast.Vertex{fetch}

                // save position
                fetch.GetNode().Position = position.NewNodePosition($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

object_dim_list:
        object_dim_list '[' dim_offset ']'
            {
                fetch := &ast.ExprArrayDimFetch{ast.Node{}, nil, $3}
                $$ = append($1, fetch)

                // save position
                fetch.GetNode().Position = position.NewNodePosition($3)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   object_dim_list '{' expr '}'
            {
                fetch := &ast.ExprArrayDimFetch{ast.Node{}, nil, $3}
                $$ = append($1, fetch)

                // save position
                fetch.GetNode().Position = position.NewNodePosition($3)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_name
            {
                fetch := &ast.ExprPropertyFetch{ast.Node{}, nil, $1}
                $$ = []ast.Vertex{fetch}

                // save position
                fetch.GetNode().Position = position.NewNodePosition($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

variable_name:
        T_STRING
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' expr '}'
            {
                $$ = $2

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, append($1.Tokens, $$.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, append($$.GetNode().Tokens[token.End], $3.Tokens...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

simple_indirect_reference:
        '$'
            {
                n := &ast.ExprVariable{ast.Node{}, nil}
                $$ = simpleIndirectReference{[]*ast.ExprVariable{n}, n}

                // save position
                n.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(n, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   simple_indirect_reference '$'
            {
                n := &ast.ExprVariable{ast.Node{}, nil}

                $1.last.VarName = n
                $1.all = append($1.all, n)
                $1.last = n
                $$ = $1

                // save position
                n.GetNode().Position = position.NewTokenPosition($2)

                // save comments
                yylex.(*Parser).setFreeFloating(n, token.Start, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

assignment_list:
        assignment_list ',' assignment_list_element
            {
                if len($1) == 0 {
                    $1 = []ast.Vertex{&ast.ExprArrayItem{ast.Node{}, false, nil, nil}}
                }

                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   assignment_list_element
            {
                if $1.(*ast.ExprArrayItem).Key == nil && $1.(*ast.ExprArrayItem).Val == nil {
                    $$ = []ast.Vertex{}
                } else {
                    $$ = []ast.Vertex{$1}
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


assignment_list_element:
        variable
            {
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, $1}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' assignment_list ')'
            {
                listNode := &ast.ExprList{ast.Node{}, $3}
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, listNode}

                // save position
                listNode.GetNode().Position = position.NewTokensPosition($1, $4)
                $$.GetNode().Position = position.NewNodePosition(listNode)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(listNode, token.List, $2.Tokens)
                yylex.(*Parser).setFreeFloating(listNode, token.ArrayPairList, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, nil}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


array_pair_list:
        /* empty */
            {
                $$ = []ast.Vertex{}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_array_pair_list possible_comma
            {
                $$ = $1

                if $2 != nil {
                    $$ = append($1, &ast.ExprArrayItem{ast.Node{}, false, nil, nil})
                }

                // save comments
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' expr T_DOUBLE_ARROW expr
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, $3, $5}
                $$ = append($1, arrayItem)

                // save position
                arrayItem.GetNode().Position = position.NewNodesPosition($3, $5)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_array_pair_list ',' expr
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, nil, $3}
                $$ = append($1, arrayItem)

                // save position
                arrayItem.GetNode().Position = position.NewNodePosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW expr
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, $1, $3}
                $$ = []ast.Vertex{arrayItem}

                // save position
                arrayItem.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, nil, $1}
                $$ = []ast.Vertex{arrayItem}

                // save position
                arrayItem.GetNode().Position = position.NewNodePosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_array_pair_list ',' expr T_DOUBLE_ARROW '&' w_variable
            {
                reference := &ast.ExprReference{ast.Node{}, $6}
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, $3, reference}
                $$ = append($1, arrayItem)

                // save position
                reference.GetNode().Position = position.NewTokenNodePosition($5, $6)
                arrayItem.GetNode().Position = position.NewNodesPosition($3, $6)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Expr, $4.Tokens)
                yylex.(*Parser).setFreeFloating(reference, token.Start, $5.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_array_pair_list ',' '&' w_variable
            {
                reference := &ast.ExprReference{ast.Node{}, $4}
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, nil, reference}
                $$ = append($1, arrayItem)

                // save position
                reference.GetNode().Position = position.NewTokenNodePosition($3, $4)
                arrayItem.GetNode().Position = position.NewTokenNodePosition($3, $4)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW '&' w_variable
            {
                reference := &ast.ExprReference{ast.Node{}, $4}
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, $1, reference}
                $$ = []ast.Vertex{arrayItem}

                // save position
                reference.GetNode().Position = position.NewTokenNodePosition($3, $4)
                arrayItem.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Expr, $2.Tokens)
                yylex.(*Parser).setFreeFloating(reference, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' w_variable
            {
                reference := &ast.ExprReference{ast.Node{}, $2}
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, nil, reference}
                $$ = []ast.Vertex{arrayItem}

                // save position
                reference.GetNode().Position = position.NewTokenNodePosition($1, $2)
                arrayItem.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating(arrayItem, token.Start, $1.Tokens)

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
                encapsed.GetNode().Position = position.NewTokenPosition($2)

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
                encapsed.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(encapsed, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_var:
        T_VARIABLE
            {
                name := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.ExprVariable{ast.Node{}, name}

                // save position
                name.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '[' encaps_var_offset ']'
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, variable, $3}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                fetch := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.ExprPropertyFetch{ast.Node{}, variable, fetch}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)
                fetch.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

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
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setToken($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                variable := &ast.ExprVariable{ast.Node{}, name}

                $$ = variable

                // save position
                name.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setToken($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
            {
                identifier := &ast.Identifier{ast.Node{}, $2.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, variable, $4}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($2)
                variable.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $6)

                // save comments
                yylex.(*Parser).setToken($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $3.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $5.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $6.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CURLY_OPEN variable '}'
            {
                $$ = $2;

                // save comments
                yylex.(*Parser).setToken($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_var_offset:
        T_STRING
            {
                $$ = &ast.ScalarString{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

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
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.ExprVariable{ast.Node{}, identifier}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

internal_functions_in_yacc:
        T_ISSET '(' isset_variables ')'
            {
                $$ = &ast.ExprIsset{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Isset, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.VarList, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EMPTY '(' variable ')'
            {
                $$ = &ast.ExprEmpty{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Empty, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EMPTY '(' expr ')'
            {
                $$ = &ast.ExprEmpty{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

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
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INCLUDE_ONCE expr
            {
                $$ = &ast.ExprIncludeOnce{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EVAL '(' expr ')'
            {
                $$ = &ast.ExprEval{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

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
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_REQUIRE_ONCE expr
            {
                $$ = &ast.ExprRequireOnce{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

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
                target := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                target.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)
                yylex.(*Parser).setFreeFloating(target, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                target.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)
                yylex.(*Parser).setFreeFloating(target, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_class_name_scalar:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS
            {
                target := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                target.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)
                yylex.(*Parser).setFreeFloating(target, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_name_scalar:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS
            {
                target := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                target.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)
                yylex.(*Parser).setFreeFloating(target, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

%%

type simpleIndirectReference struct {
	all  []*ast.ExprVariable
	last *ast.ExprVariable
}
