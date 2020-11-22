%{
package php5

import (
    "bytes"
    "strconv"

    "github.com/z7zmey/php-parser/internal/position"
    "github.com/z7zmey/php-parser/pkg/ast"
    "github.com/z7zmey/php-parser/pkg/token"
)

%}

%union{
    node                    ast.Vertex
    token                   *token.Token
    list                    []ast.Vertex
    simpleIndirectReference simpleIndirectReference

    ClassExtends            *ast.StmtClassExtends
    ClassImplements         *ast.StmtClassImplements
    InterfaceExtends        *ast.StmtInterfaceExtends
    ClosureUse              *ast.ExprClosureUse
}

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
%type <node> static_scalar_value static_operation static_var_list global_var_list
%type <node> ctor_arguments function_call_parameter_list echo_expr_list
%type <node> trait_adaptations unset_variables declare_list
%type <node> switch_case_list non_empty_function_call_parameter_list
%type <node> method_body
%type <node> foreach_statement for_statement while_statement
%type <node> foreach_variable foreach_optional_arg

%type <ClassExtends> extends_from
%type <ClassImplements> implements_list
%type <InterfaceExtends> interface_extends_list
%type <ClosureUse> lexical_vars

%type <list> top_statement_list namespace_name use_declarations use_function_declarations use_const_declarations
%type <list> inner_statement_list encaps_list isset_variables non_empty_array_pair_list
%type <list> array_pair_list assignment_list lexical_var_list elseif_list new_elseif_list non_empty_for_expr
%type <list> for_expr case_list catch_statement additional_catches
%type <list> non_empty_additional_catches parameter_list non_empty_parameter_list class_statement_list
%type <list> class_statement_list variable_modifiers method_modifiers class_variable_declaration
%type <list> interface_list trait_list trait_adaptation_list non_empty_trait_adaptation_list
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

                yylex.(*Parser).setFreeFloating(yylex.(*Parser).rootNode, token.End, yylex.(*Parser).currentToken.SkippedTokens)
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
            }
    |   /* empty */
            {
                $$ = []ast.Vertex{}
            }
;

namespace_name:
        T_STRING
            {
                $$ = []ast.Vertex{
                    &ast.NameNamePart{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        StringTkn: $1,
                        Value:     $1.Value,
                    },
                }
            }
    |   namespace_name T_NS_SEPARATOR T_STRING
            {
                $$ = append($1, &ast.NameNamePart{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($2, $3),
                    },
                    NsSeparatorTkn: $2,
                    StringTkn:      $3,
                    Value:          $3.Value,
                })
            }
;

top_statement:
        error
            {
                // error
                $$ = nil
            }
    |   statement
            {
                $$ = $1
            }
    |   function_declaration_statement
            {
                $$ = $1
            }
    |   class_declaration_statement
            {
                $$ = $1
            }
    |   T_HALT_COMPILER '(' ')' ';'
            {
                $$ = &ast.StmtHaltCompiler{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    HaltCompilerTkn:     $1,
                    OpenParenthesisTkn:  $2,
                    CloseParenthesisTkn: $3,
                    SemiColonTkn:        $4,
                }
            }
    |   T_NAMESPACE namespace_name ';'
            {
                $$ = &ast.StmtNamespace{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    NsTkn: $1,
                    Name: &ast.NameName{
                        Node:  ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Parts: $2,
                    },
                    SemiColonTkn: $3,
                }
            }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
            {
                $$ = &ast.StmtNamespace{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $5),
                    },
                    NsTkn: $1,
                    Name: &ast.NameName{
                        Node:  ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Parts: $2,
                    },
                    OpenCurlyBracket:  $3,
                    Stmts:             $4,
                    CloseCurlyBracket: $5,
                }
            }
    |   T_NAMESPACE '{' top_statement_list '}'
            {
                $$ = &ast.StmtNamespace{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    NsTkn:             $1,
                    OpenCurlyBracket:  $2,
                    Stmts:             $3,
                    CloseCurlyBracket: $4,
                }
            }
    |   T_USE use_declarations ';'
            {
                $$ = &ast.StmtUse{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    UseTkn:          $1,
                    UseDeclarations: $2,
                    SemiColonTkn:    $3,
                }
            }
    |   T_USE T_FUNCTION use_function_declarations ';'
            {
                $$ = &ast.StmtUse{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    UseTkn: $1,
                    Type: &ast.Identifier{
                        Node:  ast.Node{
                            Position: position.NewTokenPosition($2),
                        },
                        IdentifierTkn: $2,
                        Value:         $2.Value,
                    },
                    UseDeclarations: $3,
                    SemiColonTkn:    $4,
                }
            }
    |   T_USE T_CONST use_const_declarations ';'
            {
                $$ = &ast.StmtUse{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    UseTkn: $1,
                    Type: &ast.Identifier{
                        Node:  ast.Node{
                            Position: position.NewTokenPosition($2),
                        },
                        IdentifierTkn: $2,
                        Value:         $2.Value,
                    },
                    UseDeclarations: $3,
                    SemiColonTkn:    $4,
                }
            }
    |   constant_declaration ';'
            {
                $1.(*ast.StmtConstList).SemiColonTkn = $2
                $1.(*ast.StmtConstList).Node.Position = position.NewNodeTokenPosition($1, $2)
                $$ = $1
            }
;

use_declarations:
        use_declarations ',' use_declaration
            {
                $1[len($1)-1].(*ast.StmtUseDeclaration).CommaTkn = $2

                $$ = append($1, $3)
            }
    |   use_declaration
            {
                $$ = []ast.Vertex{$1}
            }
;

use_declaration:
        namespace_name
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewNodeListPosition($1),
                    },
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1),
                        },
                        Parts: $1,
                    },
                }
            }
    |   namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewNodeListTokenPosition($1, $3),
                    },
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1),
                        },
                        Parts: $1,
                    },
                    AsTkn: $2,
                    Alias: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $2),
                    },
                    NsSeparatorTkn: $1,
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Parts: $2,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    NsSeparatorTkn: $1,
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Parts: $2,
                    },
                    AsTkn: $3,
                    Alias: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($4),
                        },
                        IdentifierTkn: $4,
                        Value:         $4.Value,
                    },
                }
            }
;

use_function_declarations:
        use_function_declarations ',' use_function_declaration
            {
                $1[len($1)-1].(*ast.StmtUseDeclaration).CommaTkn = $2

                $$ = append($1, $3)
            }
    |   use_function_declaration
            {
                $$ = []ast.Vertex{$1}
            }
;

use_function_declaration:
        namespace_name
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewNodeListPosition($1),
                    },
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1),
                        },
                        Parts: $1,
                    },
                }
            }
    |   namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewNodeListTokenPosition($1, $3),
                    },
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1),
                        },
                        Parts: $1,
                    },
                    AsTkn: $2,
                    Alias: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $2),
                    },
                    NsSeparatorTkn: $1,
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Parts: $2,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    NsSeparatorTkn: $1,
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Parts: $2,
                    },
                    AsTkn: $3,
                    Alias: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($4),
                        },
                        IdentifierTkn: $4,
                        Value:         $4.Value,
                    },
                }
            }
;

use_const_declarations:
        use_const_declarations ',' use_const_declaration
            {
                $1[len($1)-1].(*ast.StmtUseDeclaration).CommaTkn = $2

                $$ = append($1, $3)
            }
    |   use_const_declaration
            {
                $$ = []ast.Vertex{$1}
            }
;

use_const_declaration:
        namespace_name
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewNodeListPosition($1),
                    },
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1),
                        },
                        Parts: $1,
                    },
                }
            }
    |   namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewNodeListTokenPosition($1, $3),
                    },
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1),
                        },
                        Parts: $1,
                    },
                    AsTkn: $2,
                    Alias: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $2),
                    },
                    NsSeparatorTkn: $1,
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Parts: $2,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    NsSeparatorTkn: $1,
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Parts: $2,
                    },
                    AsTkn: $3,
                    Alias: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($4),
                        },
                        IdentifierTkn: $4,
                        Value:         $4.Value,
                    },
                }
            }
;

constant_declaration:
        constant_declaration ',' T_STRING '=' static_scalar
            {
                constList := $1.(*ast.StmtConstList)
                constList.Node.Position = position.NewNodesPosition($1, $5)
                constList.SeparatorTkns = append(constList.SeparatorTkns, $2)
                constList.Consts = append(constList.Consts, &ast.StmtConstant{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($3, $5),
                    },
                    Name: &ast.Identifier{
                        Node:  ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                    EqualTkn: $4,
                    Expr:     $5,
                })

                $$ = $1

                yylex.(*Parser).setFreeFloating(lastNode($$.(*ast.StmtConstList).Consts).(*ast.StmtConstant).Name, token.Start, $3.SkippedTokens)
            }
    |   T_CONST T_STRING '=' static_scalar
            {
                $$ = &ast.StmtConstList{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $4),
                    },
                    ConstTkn: $1,
                    Consts: []ast.Vertex{
                        &ast.StmtConstant{
                            Node: ast.Node{
                                Position: position.NewTokenNodePosition($2, $4),
                            },
                            Name: &ast.Identifier{
                                Node:  ast.Node{
                                    Position: position.NewTokenPosition($2),
                                },
                                IdentifierTkn: $2,
                                Value:         $2.Value,
                            },
                            EqualTkn: $3,
                            Expr:     $4,
                        },
                    },
                }

                yylex.(*Parser).setFreeFloating(lastNode($$.(*ast.StmtConstList).Consts).(*ast.StmtConstant).Name, token.Start, $2.SkippedTokens)
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
            }
    |   /* empty */
            {
                $$ = []ast.Vertex{}
            }
;


inner_statement:
        error
            {
                // error
                $$ = nil
            }
    |   statement
            {
                $$ = $1
            }
    |   function_declaration_statement
            {
                $$ = $1
            }
    |   class_declaration_statement
            {
                $$ = $1
            }
    |   T_HALT_COMPILER '(' ')' ';'
            {
                $$ = &ast.StmtHaltCompiler{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    HaltCompilerTkn:     $1,
                    OpenParenthesisTkn:  $2,
                    CloseParenthesisTkn: $3,
                    SemiColonTkn:        $4,
                }
            }
;


statement:
        unticked_statement
            {
                $$ = $1
            }
    |   T_STRING ':'
            {
                $$ = &ast.StmtLabel{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    LabelName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                    ColonTkn:  $2,
                }
            }
;

unticked_statement:
        '{' inner_statement_list '}'
            {
                $$ = &ast.StmtStmtList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenCurlyBracket:  $1,
                    Stmts:             $2,
                    CloseCurlyBracket: $3,
                }
            }
    |   T_IF parenthesis_expr statement elseif_list else_single
            {
                pos := position.NewTokenNodePosition($1, $3)
                if $5 != nil {
                    pos = position.NewTokenNodePosition($1, $5)
                } else if len($4) > 0 {
                    pos = position.NewTokenNodeListPosition($1, $4)
                }

                $$ = &ast.StmtIf{
                    Node: ast.Node{
                        Position: pos,
                    },
                    IfTkn:               $1,
                    OpenParenthesisTkn:  $2.(*ast.ParserBrackets).OpenBracketTkn,
                    Cond:                $2.(*ast.ParserBrackets).Child,
                    CloseParenthesisTkn: $2.(*ast.ParserBrackets).CloseBracketTkn,
                    Stmt:                $3,
                    ElseIf:              $4,
                    Else:                $5,
                }
            }
    |   T_IF parenthesis_expr ':' inner_statement_list new_elseif_list new_else_single T_ENDIF ';'
            {
                $$ = &ast.StmtIf{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $8),
                    },
                    Alt:                 true,
                    IfTkn:               $1,
                    OpenParenthesisTkn:  $2.(*ast.ParserBrackets).OpenBracketTkn,
                    Cond:                $2.(*ast.ParserBrackets).Child,
                    CloseParenthesisTkn: $2.(*ast.ParserBrackets).CloseBracketTkn,
                    ColonTkn:            $3,
                    Stmt: &ast.StmtStmtList{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($4),
                        },
                        Stmts: $4,
                    },
                    ElseIf:       $5,
                    Else:         $6,
                    EndIfTkn:     $7,
                    SemiColonTkn: $8,
                }
            }
    |   T_WHILE parenthesis_expr while_statement
            {
                $3.(*ast.StmtWhile).WhileTkn = $1
                $3.(*ast.StmtWhile).OpenParenthesisTkn = $2.(*ast.ParserBrackets).OpenBracketTkn
                $3.(*ast.StmtWhile).Cond = $2.(*ast.ParserBrackets).Child
                $3.(*ast.StmtWhile).CloseParenthesisTkn = $2.(*ast.ParserBrackets).CloseBracketTkn
                $3.(*ast.StmtWhile).Node.Position = position.NewTokenNodePosition($1, $3)

                $$ = $3
            }
    |   T_DO statement T_WHILE parenthesis_expr ';'
            {
                $$ = &ast.StmtDo{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $5),
                    },
                    DoTkn:               $1,
                    Stmt:                $2,
                    WhileTkn:            $3,
                    OpenParenthesisTkn:  $4.(*ast.ParserBrackets).OpenBracketTkn,
                    Cond:                $4.(*ast.ParserBrackets).Child,
                    CloseParenthesisTkn: $4.(*ast.ParserBrackets).CloseBracketTkn,
                    SemiColonTkn:        $5,
                }
            }
    |   T_FOR '(' for_expr ';' for_expr ';' for_expr ')' for_statement
            {
                $9.(*ast.StmtFor).ForTkn = $1
                $9.(*ast.StmtFor).OpenParenthesisTkn = $2
                $9.(*ast.StmtFor).Init = $3
                $9.(*ast.StmtFor).InitSemiColonTkn = $4
                $9.(*ast.StmtFor).Cond = $5
                $9.(*ast.StmtFor).CondSemiColonTkn = $6
                $9.(*ast.StmtFor).Loop = $7
                $9.(*ast.StmtFor).CloseParenthesisTkn = $8
                $9.(*ast.StmtFor).Node.Position = position.NewTokenNodePosition($1, $9)

                $$ = $9
            }
    |   T_SWITCH parenthesis_expr switch_case_list
            {
                $3.(*ast.StmtSwitch).SwitchTkn = $1
                $3.(*ast.StmtSwitch).OpenParenthesisTkn = $2.(*ast.ParserBrackets).OpenBracketTkn
                $3.(*ast.StmtSwitch).Cond = $2.(*ast.ParserBrackets).Child
                $3.(*ast.StmtSwitch).CloseParenthesisTkn = $2.(*ast.ParserBrackets).CloseBracketTkn
                $3.(*ast.StmtSwitch).Node.Position = position.NewTokenNodePosition($1, $3)

                $$ = $3
            }
    |   T_BREAK ';'
            {
                $$ = &ast.StmtBreak{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    BreakTkn:     $1,
                    SemiColonTkn: $2,
                }
            }
    |   T_BREAK expr ';'
            {
                $$ = &ast.StmtBreak{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    BreakTkn:     $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   T_CONTINUE ';'
            {
                $$ = &ast.StmtContinue{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    ContinueTkn:  $1,
                    SemiColonTkn: $2,
                }
            }
    |   T_CONTINUE expr ';'
            {
                $$ = &ast.StmtContinue{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    ContinueTkn:  $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   T_RETURN ';'
            {
                $$ = &ast.StmtReturn{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    ReturnTkn:    $1,
                    SemiColonTkn: $2,
                }
            }
    |   T_RETURN expr_without_variable ';'
            {
                $$ = &ast.StmtReturn{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    ReturnTkn:    $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   T_RETURN variable ';'
            {
                $$ = &ast.StmtReturn{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    ReturnTkn:    $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   yield_expr ';'
            {
                $$ = &ast.StmtExpression{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.SkippedTokens)
            }
    |   T_GLOBAL global_var_list ';'
            {
                $2.(*ast.StmtGlobal).GlobalTkn = $1
                $2.(*ast.StmtGlobal).SemiColonTkn = $3
                $2.(*ast.StmtGlobal).SeparatorTkns = append($2.(*ast.StmtGlobal).SeparatorTkns, nil)
                $2.(*ast.StmtGlobal).Node.Position = position.NewTokensPosition($1, $3)

                $$ = $2
            }
    |   T_STATIC static_var_list ';'
            {
                $2.(*ast.StmtStatic).StaticTkn = $1
                $2.(*ast.StmtStatic).SemiColonTkn = $3
                $2.(*ast.StmtStatic).SeparatorTkns = append($2.(*ast.StmtStatic).SeparatorTkns, nil)
                $2.(*ast.StmtStatic).Node.Position = position.NewTokensPosition($1, $3)

                $$ = $2
            }
    |   T_ECHO echo_expr_list ';'
            {
                $2.(*ast.StmtEcho).EchoTkn = $1
                $2.(*ast.StmtEcho).SemiColonTkn = $3
                $2.(*ast.StmtEcho).Node.Position = position.NewTokensPosition($1, $3)

                $$ = $2
            }
    |   T_INLINE_HTML
            {
                $$ = &ast.StmtInlineHtml{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    InlineHtmlTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   expr ';'
            {
                $$ = &ast.StmtExpression{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.SkippedTokens)
            }
    |   T_UNSET '(' unset_variables ')' ';'
            {
                $3.(*ast.StmtUnset).UnsetTkn = $1
                $3.(*ast.StmtUnset).OpenParenthesisTkn = $2
                $3.(*ast.StmtUnset).CloseParenthesisTkn = $4
                $3.(*ast.StmtUnset).SemiColonTkn = $5
                $3.(*ast.StmtUnset).Node.Position = position.NewTokensPosition($1, $5)

                $$ = $3
            }
    |   T_FOREACH '(' variable T_AS foreach_variable foreach_optional_arg ')' foreach_statement
            {
                $8.(*ast.StmtForeach).ForeachTkn = $1
                $8.(*ast.StmtForeach).OpenParenthesisTkn = $2
                $8.(*ast.StmtForeach).Expr = $3
                $8.(*ast.StmtForeach).AsTkn = $4
                if $6 == nil {
                    $8.(*ast.StmtForeach).Var = $5
                } else {
                    $8.(*ast.StmtForeach).Key = $5
                    $8.(*ast.StmtForeach).DoubleArrowTkn = $6.(*ast.StmtForeach).DoubleArrowTkn
                    $8.(*ast.StmtForeach).Var = $6.(*ast.StmtForeach).Var
                }
                $8.(*ast.StmtForeach).CloseParenthesisTkn = $7
                $8.(*ast.StmtForeach).Node.Position = position.NewTokenNodePosition($1, $8)

                $$ = $8
            }
    |   T_FOREACH '(' expr_without_variable T_AS foreach_variable foreach_optional_arg ')' foreach_statement
            {
                $8.(*ast.StmtForeach).ForeachTkn = $1
                $8.(*ast.StmtForeach).OpenParenthesisTkn = $2
                $8.(*ast.StmtForeach).Expr = $3
                $8.(*ast.StmtForeach).AsTkn = $4
                if $6 == nil {
                    $8.(*ast.StmtForeach).Var = $5
                } else {
                    $8.(*ast.StmtForeach).Key = $5
                    $8.(*ast.StmtForeach).DoubleArrowTkn = $6.(*ast.StmtForeach).DoubleArrowTkn
                    $8.(*ast.StmtForeach).Var = $6.(*ast.StmtForeach).Var
                }
                $8.(*ast.StmtForeach).CloseParenthesisTkn = $7
                $8.(*ast.StmtForeach).Node.Position = position.NewTokenNodePosition($1, $8)

                $$ = $8
            }
    |   T_DECLARE '(' declare_list ')' declare_statement
            {
                $5.(*ast.StmtDeclare).DeclareTkn = $1
                $5.(*ast.StmtDeclare).OpenParenthesisTkn = $2
                $5.(*ast.StmtDeclare).Consts = $3.(*ast.ParserSeparatedList).Items
                $5.(*ast.StmtDeclare).SeparatorTkns = $3.(*ast.ParserSeparatedList).SeparatorTkns
                $5.(*ast.StmtDeclare).CloseParenthesisTkn = $4
                $5.(*ast.StmtDeclare).Node.Position = position.NewTokenNodePosition($1, $5)

                $$ = $5
            }
    |   ';'
            {
                $$ = &ast.StmtNop{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    SemiColonTkn: $1,
                }
            }
    |   T_TRY '{' inner_statement_list '}' catch_statement finally_statement
            {
                $$ = &ast.StmtTry{
                    TryTkn:            $1,
                    OpenCurlyBracket:  $2,
                    Stmts:             $3,
                    CloseCurlyBracket: $4,
                    Catches:           $5,
                    Finally:           $6,
                }

                if $6 == nil {
                    $$.GetNode().Position = position.NewTokenNodeListPosition($1, $5)
                } else {
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $6)
                }
            }
    |   T_THROW expr ';'
            {
                $$ = &ast.StmtThrow{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    ThrowTkn:     $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   T_GOTO T_STRING ';'
            {
                $$ = &ast.StmtGoto{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    GotoTkn: $1,
                    Label: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                    SemiColonTkn: $3,
                }
            }
;

catch_statement:
        /* empty */
            {
                $$ = []ast.Vertex{}
            }
    |   T_CATCH '(' fully_qualified_class_name T_VARIABLE ')' '{' inner_statement_list '}' additional_catches
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($4),
                    },
                    IdentifierTkn: $4,
                    Value:         $4.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                catch := &ast.StmtCatch{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $8),
                    },
                    CatchTkn:             $1,
                    OpenParenthesisTkn:   $2,
                    Types:                []ast.Vertex{$3},
                    Var:                  variable,
                    CloseParenthesisTkn:  $5,
                    OpenCurlyBracketTkn:  $6,
                    Stmts:                $7,
                    CloseCurlyBracketTkn: $8,
                }
                $$ = append([]ast.Vertex{catch}, $9...)

                // save position
                variable.GetNode().Position = position.NewTokenPosition($4)

                // save comments
                yylex.(*Parser).setFreeFloating(variable, token.Start, $4.SkippedTokens)
            }
;

finally_statement:
        /* empty */
            {
                $$ = nil
            }
    |   T_FINALLY '{' inner_statement_list '}'
            {
                $$ = &ast.StmtFinally{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    FinallyTkn:           $1,
                    OpenCurlyBracketTkn:  $2,
                    Stmts:                $3,
                    CloseCurlyBracketTkn: $4,
                }
            }
;

additional_catches:
        non_empty_additional_catches
            {
                $$ = $1
            }
    |   /* empty */
            {
                $$ = []ast.Vertex{}
            }
;

non_empty_additional_catches:
        additional_catch
            {
                $$ = []ast.Vertex{$1}
            }
    |   non_empty_additional_catches additional_catch
            {
                $$ = append($1, $2)
            }
;

additional_catch:
        T_CATCH '(' fully_qualified_class_name T_VARIABLE ')' '{' inner_statement_list '}'
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($4),
                    },
                    IdentifierTkn: $4,
                    Value:         $4.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.StmtCatch{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $8),
                    },
                    CatchTkn:             $1,
                    OpenParenthesisTkn:   $2,
                    Types:                []ast.Vertex{$3},
                    Var:                  variable,
                    CloseParenthesisTkn:  $5,
                    OpenCurlyBracketTkn:  $6,
                    Stmts:                $7,
                    CloseCurlyBracketTkn: $8,
                }

                // save position
                variable.GetNode().Position = position.NewTokenPosition($4)

                // save comments
                yylex.(*Parser).setFreeFloating(variable, token.Start, $4.SkippedTokens)
            }
;

unset_variables:
        unset_variable
            {
                $$ = &ast.StmtUnset{
                    Vars: []ast.Vertex{$1},
                }
            }
    |   unset_variables ',' unset_variable
            {
                $1.(*ast.StmtUnset).Vars = append($1.(*ast.StmtUnset).Vars, $3)
                $1.(*ast.StmtUnset).SeparatorTkns = append($1.(*ast.StmtUnset).SeparatorTkns, $2)

                $$ = $1
            }
;

unset_variable:
        variable
            {
                $$ = $1
            }
;

function_declaration_statement:
        unticked_function_declaration_statement
            {
                $$ = $1
            }
;

class_declaration_statement:
        unticked_class_declaration_statement
            {
                $$ = $1
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
                name := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                $$ = &ast.StmtFunction{ast.Node{}, $2 != nil, name, $5, nil, $8}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $9)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $2.SkippedTokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.ParamList, $6.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Params, $7.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $9.SkippedTokens)
            }
;

unticked_class_declaration_statement:
        class_entry_type T_STRING extends_from implements_list '{' class_statement_list '}'
            {
                switch n := $1.(type) {
                    case *ast.StmtClass :
                        n.Position = position.NewNodeTokenPosition($1, $7)
                        n.ClassName = &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($2),
                            },
                            IdentifierTkn: $2,
                            Value:         $2.Value,
                        }
                        n.Extends = $3
                        n.Implements = $4
                        n.OpenCurlyBracket = $5
                        n.Stmts = $6
                        n.CloseCurlyBracket = $7
                    case *ast.StmtTrait :
                        n.Position = position.NewNodeTokenPosition($1, $7)
                        n.TraitName = &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($2),
                            },
                            IdentifierTkn: $2,
                            Value:         $2.Value,
                        }
                        n.Extends = $3
                        n.Implements = $4
                        n.OpenCurlyBracket = $5
                        n.Stmts = $6
                        n.CloseCurlyBracket = $7
                }

                $$ = $1
            }
    |   interface_entry T_STRING interface_extends_list '{' class_statement_list '}'
            {
                name := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($2),
                    },
                    IdentifierTkn: $2,
                    Value:         $2.Value,
                }
                $$ = &ast.StmtInterface{ast.Node{}, name, $3, $5}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $4.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $6.SkippedTokens)
            }
;


class_entry_type:
        T_CLASS
            {
                $$ = &ast.StmtClass{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    ClassTkn: $1,
                }
            }
    |   T_ABSTRACT T_CLASS
            {
                $$ = &ast.StmtClass{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    Modifiers: []ast.Vertex{
                        &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($1),
                            },
                            IdentifierTkn: $1,
                            Value:         $1.Value,
                        },
                    },
                    ClassTkn: $2,
                }
            }
    |   T_TRAIT
            {
                $$ = &ast.StmtTrait{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    TraitTkn: $1,
                }
            }
    |   T_FINAL T_CLASS
            {
                $$ = &ast.StmtClass{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    Modifiers: []ast.Vertex{
                        &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($1),
                            },
                            IdentifierTkn: $1,
                            Value:         $1.Value,
                        },
                    },
                    ClassTkn: $2,
                }
            }
;

extends_from:
        /* empty */
            {
                $$ = nil
            }
    |   T_EXTENDS fully_qualified_class_name
            {
                $$ = &ast.StmtClassExtends{ast.Node{}, $2};

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
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
            }
    |   T_EXTENDS interface_list
            {
                $$ = &ast.StmtInterfaceExtends{ast.Node{}, $2};

                // save position
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

implements_list:
        /* empty */
            {
                $$ = nil
            }
    |   T_IMPLEMENTS interface_list
            {
                $$ = &ast.StmtClassImplements{ast.Node{}, $2};

                // save position
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

interface_list:
        fully_qualified_class_name
            {
                $$ = []ast.Vertex{$1}
            }
    |   interface_list ',' fully_qualified_class_name
            {
                switch n := lastNode($1).(type) {
                    case *ast.NameName: n.ListSeparatorTkn = $2
                    case *ast.NameFullyQualified: n.ListSeparatorTkn = $2
                    case *ast.NameRelative: n.ListSeparatorTkn = $2
                }
                $$ = append($1, $3)
            }
;

foreach_optional_arg:
        /* empty */
            {
                $$ = nil
            }
    |   T_DOUBLE_ARROW foreach_variable
            {
                $$ = &ast.StmtForeach{
                    DoubleArrowTkn: $1,
                    Var:            $2,
                }
            }
;

foreach_variable:
        variable
            {
                $$ = $1
            }
    |   '&' variable
            {
                $$ = &ast.ExprReference{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_LIST '(' assignment_list ')'
            {
                $$ = &ast.ExprList{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.List, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $4.SkippedTokens)
            }
;

for_statement:
        statement
            {
                $$ = &ast.StmtFor{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Stmt: $1,
                }
            }
    |   ':' inner_statement_list T_ENDFOR ';'
            {
                $$ = &ast.StmtFor{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    Alt:      true,
                    ColonTkn: $1,
                    Stmt: &ast.StmtStmtList{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Stmts: $2,
                    },
                    EndForTkn:    $3,
                    SemiColonTkn: $4,
                }
            }
;

foreach_statement:
        statement
            {
                $$ = &ast.StmtForeach{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Stmt: $1,
                }
            }
    |   ':' inner_statement_list T_ENDFOREACH ';'
            {
                $$ = &ast.StmtForeach{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    Alt:      true,
                    ColonTkn: $1,
                    Stmt: &ast.StmtStmtList{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Stmts: $2,
                    },
                    EndForeachTkn: $3,
                    SemiColonTkn:  $4,
                }
            }
;


declare_statement:
        statement
            {
                $$ = &ast.StmtDeclare{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Stmt: $1,
                }
            }
    |   ':' inner_statement_list T_ENDDECLARE ';'
            {
                $$ = &ast.StmtDeclare{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    Alt:      true,
                    ColonTkn: $1,
                    Stmt: &ast.StmtStmtList{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Stmts: $2,
                    },
                    EndDeclareTkn: $3,
                    SemiColonTkn:  $4,
                }
            }
;


declare_list:
        T_STRING '=' static_scalar
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{
                        &ast.StmtConstant{
                            Node: ast.Node{
                                Position: position.NewTokenNodePosition($1, $3),
                            },
                            Name: &ast.Identifier{
                                Node:  ast.Node{
                                    Position: position.NewTokenPosition($1),
                                },
                                IdentifierTkn: $1,
                                Value:         $1.Value,
                            },
                            EqualTkn: $2,
                            Expr:     $3,
                        },
                    },
                }

                yylex.(*Parser).setFreeFloating(lastNode($$.(*ast.ParserSeparatedList).Items).(*ast.StmtConstant).Name, token.Start, $1.SkippedTokens)
            }
    |   declare_list ',' T_STRING '=' static_scalar
            {
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append(
                    $1.(*ast.ParserSeparatedList).Items,
                    &ast.StmtConstant{
                        Node: ast.Node{
                            Position: position.NewTokenNodePosition($3, $5),
                        },
                        Name: &ast.Identifier{
                            Node:  ast.Node{
                                Position: position.NewTokenPosition($3),
                            },
                            IdentifierTkn: $3,
                            Value:         $3.Value,
                        },
                        EqualTkn: $4,
                        Expr:     $5,
                    },
                )

                $$ = $1

                yylex.(*Parser).setFreeFloating(lastNode($$.(*ast.ParserSeparatedList).Items).(*ast.StmtConstant).Name, token.Start, $3.SkippedTokens)
            }
;


switch_case_list:
        '{' case_list '}'
            {
                $$ = &ast.StmtSwitch{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenCurlyBracketTkn:  $1,
                    CaseList:             $2,
                    CloseCurlyBracketTkn: $3,
                }
            }
    |   '{' ';' case_list '}'
            {
                $$ = &ast.StmtSwitch{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    OpenCurlyBracketTkn:  $1,
                    CaseSeparatorTkn:     $2,
                    CaseList:             $3,
                    CloseCurlyBracketTkn: $4,
                }
            }
    |   ':' case_list T_ENDSWITCH ';'
            {
                $$ = &ast.StmtSwitch{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    Alt:          true,
                    ColonTkn:     $1,
                    CaseList:     $2,
                    EndSwitchTkn: $3,
                    SemiColonTkn: $4,
                }
            }
    |   ':' ';' case_list T_ENDSWITCH ';'
            {
                $$ = &ast.StmtSwitch{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $5),
                    },
                    Alt:              true,
                    ColonTkn:         $1,
                    CaseSeparatorTkn: $2,
                    CaseList:         $3,
                    EndSwitchTkn:     $4,
                    SemiColonTkn:     $5,
                }
            }
;


case_list:
        /* empty */
            {
                $$ = nil
            }
    |   case_list T_CASE expr case_separator inner_statement_list
            {
                $$ = append($1, &ast.StmtCase{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($2, $5),
                    },
                    CaseTkn:          $2,
                    Cond:             $3,
                    CaseSeparatorTkn: $4,
                    Stmts:            $5,
                })
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                $$ = append($1, &ast.StmtDefault{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($2, $4),
                    },
                    DefaultTkn:       $2,
                    CaseSeparatorTkn: $3,
                    Stmts:            $4,
                })
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
                $$ = &ast.StmtWhile{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Stmt: $1,
                }
            }
    |   ':' inner_statement_list T_ENDWHILE ';'
            {
                $$ = &ast.StmtWhile{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    Alt:      true,
                    ColonTkn: $1,
                    Stmt: &ast.StmtStmtList{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Stmts: $2,
                    },
                    EndWhileTkn:  $3,
                    SemiColonTkn: $4,
                }
            }
;



elseif_list:
        /* empty */
            {
                $$ = nil
            }
    |   elseif_list T_ELSEIF parenthesis_expr statement
            {
                $$ = append($1, &ast.StmtElseIf{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($2, $4),
                    },
                    ElseIfTkn:           $2,
                    OpenParenthesisTkn:  $3.(*ast.ParserBrackets).OpenBracketTkn,
                    Cond:                $3.(*ast.ParserBrackets).Child,
                    CloseParenthesisTkn: $3.(*ast.ParserBrackets).CloseBracketTkn,
                    Stmt:                $4,
                })
            }
;


new_elseif_list:
        /* empty */
            {
                $$ = nil
            }
    |   new_elseif_list T_ELSEIF parenthesis_expr ':' inner_statement_list
            {
                $$ = append($1, &ast.StmtElseIf{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($2, $5),
                    },
                    Alt:                 true,
                    ElseIfTkn:           $2,
                    OpenParenthesisTkn:  $3.(*ast.ParserBrackets).OpenBracketTkn,
                    Cond:                $3.(*ast.ParserBrackets).Child,
                    CloseParenthesisTkn: $3.(*ast.ParserBrackets).CloseBracketTkn,
                    ColonTkn:            $4,
                    Stmt: &ast.StmtStmtList{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($5),
                        },
                        Stmts: $5,
                    },
                })
            }
;


else_single:
        /* empty */
            {
                $$ = nil
            }
    |   T_ELSE statement
            {
                $$ = &ast.StmtElse{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    ElseTkn: $1,
                    Stmt:    $2,
                }
            }
;


new_else_single:
        /* empty */
            {
                $$ = nil
            }
    |   T_ELSE ':' inner_statement_list
            {
                $$ = &ast.StmtElse{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $3),
                    },
                    Alt:      true,
                    ElseTkn:  $1,
                    ColonTkn: $2,
                    Stmt: &ast.StmtStmtList{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($3),
                        },
                        Stmts: $3,
                    },
                }
            }
;


parameter_list:
        non_empty_parameter_list
            {
                $$ = $1;
            }
    |   /* empty */
            {
                $$ = nil
            }
;

non_empty_parameter_list:
        parameter
            {
                $$ = []ast.Vertex{$1}
            }
    |   non_empty_parameter_list ',' parameter
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
;

parameter:
        optional_class_type is_reference is_variadic T_VARIABLE
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($4),
                    },
                    IdentifierTkn: $4,
                    Value:         $4.Value,
                }

                var variable ast.Vertex
                variable = &ast.ExprVariable{ast.Node{}, identifier}
                variable.GetNode().Position = position.NewTokenPosition($4)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $4.SkippedTokens)

                if $3 != nil {
                    variable = &ast.Variadic{
                        Node: ast.Node{
                            Position: position.NewTokensPosition($3, $4),
                        },
                        VariadicTkn: $3,
                        Var:         variable,
                    }
                }

                if $2 != nil {
                    variable = &ast.Reference{
                        Node: ast.Node{
                            Position: position.NewTokensPosition($2, $4),
                        },
                        AmpersandTkn: $2,
                        Var:          variable,
                    }
                }

                pos := position.NewTokenPosition($4)
                if $1 != nil {
                    pos = position.NewNodeTokenPosition($1, $4)
                } else if $2 != nil {
                    pos = position.NewTokensPosition($2, $4)
                } else if $3 != nil {
                    pos = position.NewTokensPosition($3, $4)
                }

                $$ = &ast.Parameter{
                    Node: ast.Node{
                        Position: pos,
                    },
                    Type: $1,
                    Var:  variable,
                }
            }
    |   optional_class_type is_reference is_variadic T_VARIABLE '=' expr
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($4),
                    },
                    IdentifierTkn: $4,
                    Value:         $4.Value,
                }

                var variable ast.Vertex
                variable = &ast.ExprVariable{ast.Node{}, identifier}
                variable.GetNode().Position = position.NewTokenPosition($4)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $4.SkippedTokens)
                yylex.(*Parser).setFreeFloating(variable, token.End, $5.SkippedTokens)

                if $3 != nil {
                    variable = &ast.Variadic{
                        Node: ast.Node{
                            Position: position.NewTokensPosition($3, $4),
                        },
                        VariadicTkn: $3,
                        Var:         variable,
                    }
                }

                if $2 != nil {
                    variable = &ast.Reference{
                        Node: ast.Node{
                            Position: position.NewTokensPosition($2, $4),
                        },
                        AmpersandTkn: $2,
                        Var:          variable,
                    }
                }

                pos := position.NewTokenNodePosition($4, $6)
                if $1 != nil {
                    pos = position.NewNodesPosition($1, $6)
                } else if $2 != nil {
                    pos = position.NewTokenNodePosition($2, $6)
                } else if $3 != nil {
                    pos = position.NewTokenNodePosition($3, $6)
                }

                $$ = &ast.Parameter{
                    Node: ast.Node{
                        Position: pos,
                    },
                    Type:         $1,
                    Var:          variable,
                    EqualTkn:     $5,
                    DefaultValue: $6,
                }
            }
;


optional_class_type:
        /* empty */
            {
                $$ = nil
            }
    |   T_ARRAY
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_CALLABLE
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   fully_qualified_class_name
            {
                $$ = $1
            }
;


function_call_parameter_list:
        '(' ')'
            {
                $$ = &ast.ArgumentList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    OpenParenthesisTkn: $1,
                    CloseParenthesisTkn: $2,
                }
            }
    |   '(' non_empty_function_call_parameter_list ')'
            {
                argumentList := $2.(*ast.ArgumentList)
                argumentList.Position = position.NewTokensPosition($1, $3)
                argumentList.OpenParenthesisTkn = $1
                argumentList.CloseParenthesisTkn = $3

                $$ = argumentList
            }
    |   '(' yield_expr ')'
            {
                $$ = &ast.ArgumentList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenParenthesisTkn: $1,
                    Arguments: []ast.Vertex{
                        &ast.Argument{
                            Node: ast.Node{
                                Position: position.NewNodePosition($2),
                            },
                            Expr: $2,
                        },
                    },
                    CloseParenthesisTkn: $3,
                }
            }
;


non_empty_function_call_parameter_list:
        function_call_parameter
            {
                $$ = &ast.ArgumentList{
                    Arguments: []ast.Vertex{$1},
                }
            }
    |   non_empty_function_call_parameter_list ',' function_call_parameter
            {
                $1.(*ast.ArgumentList).SeparatorTkns = append($1.(*ast.ArgumentList).SeparatorTkns, $2)
                $1.(*ast.ArgumentList).Arguments = append($1.(*ast.ArgumentList).Arguments, $3)

                $$ = $1
            }
;

function_call_parameter:
        expr_without_variable
            {
                $$ = &ast.Argument{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Expr: $1,
                }
            }
    |   variable
            {
                $$ = &ast.Argument{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Expr: $1,
                }
            }
    |   '&' w_variable
            {
                $$ = &ast.Argument{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    AmpersandTkn: $1,
                    Expr:         $2,
                }
            }
    |   T_ELLIPSIS expr
            {
                $$ = &ast.Argument{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    VariadicTkn: $1,
                    Expr:        $2,
                }
            }
;

global_var_list:
        global_var_list ',' global_var
            {
                $1.(*ast.StmtGlobal).Vars = append($1.(*ast.StmtGlobal).Vars, $3)
                $1.(*ast.StmtGlobal).SeparatorTkns = append($1.(*ast.StmtGlobal).SeparatorTkns, $2)

                $$ = $1
            }
    |   global_var
            {
                $$ = &ast.StmtGlobal{
                    Vars: []ast.Vertex{$1},
                }
            }
;


global_var:
        T_VARIABLE
            {
                name := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                $$ = &ast.ExprVariable{ast.Node{}, name}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '$' r_variable
            {
                $$ = &ast.ExprVariable{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '$' '{' expr '}'
            {
                $$ = &ast.ExprVariable{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($3, token.Start, append($2.SkippedTokens, $3.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($3, token.End, append($3.GetNode().Tokens[token.End], $4.SkippedTokens...))
            }
;


static_var_list:
        static_var_list ',' T_VARIABLE
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}

                $1.(*ast.StmtStatic).Vars = append($1.(*ast.StmtStatic).Vars, &ast.StmtStaticVar{
                    Node: ast.Node{
                        Position:  position.NewTokenPosition($3),
                    },
                    Var: variable,
                })
                $1.(*ast.StmtStatic).SeparatorTkns = append($1.(*ast.StmtStatic).SeparatorTkns, $2)

                $$ = $1

                // save position
                variable.GetNode().Position = position.NewTokenPosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(variable, token.Start, $3.SkippedTokens)
            }
    |   static_var_list ',' T_VARIABLE '=' static_scalar
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}

                $1.(*ast.StmtStatic).Vars = append($1.(*ast.StmtStatic).Vars, &ast.StmtStaticVar{
                    Node: ast.Node{
                        Position:  position.NewTokenNodePosition($3, $5),
                    },
                    Var:      variable,
                    EqualTkn: $4,
                    Expr:     $5,
                })
                $1.(*ast.StmtStatic).SeparatorTkns = append($1.(*ast.StmtStatic).SeparatorTkns, $2)

                $$ = $1

                // save position
                variable.GetNode().Position = position.NewTokenPosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(variable, token.Start, $3.SkippedTokens)
            }
    |   T_VARIABLE
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}

                $$ = &ast.StmtStatic{
                    Vars: []ast.Vertex{
                        &ast.StmtStaticVar{
                            Node: ast.Node{
                                Position:  position.NewTokenPosition($1),
                            },
                            Var: variable,
                        },
                    },
                }

                // save position
                variable.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(variable, token.Start, $1.SkippedTokens)
            }
    |   T_VARIABLE '=' static_scalar
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}

                $$ = &ast.StmtStatic{
                    Vars: []ast.Vertex{
                        &ast.StmtStaticVar{
                            Node: ast.Node{
                                Position:  position.NewTokenNodePosition($1, $3),
                            },
                            Var:      variable,
                            EqualTkn: $2,
                            Expr:     $3,
                        },
                    },
                }

                // save position
                variable.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(variable, token.Start, $1.SkippedTokens)
            }
;


class_statement_list:
        class_statement_list class_statement
            {
                $$ = append($1, $2)
            }
    |   /* empty */
            {
                $$ = []ast.Vertex{}
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
                yylex.(*Parser).setFreeFloating($$, token.PropertyList, $3.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.SkippedTokens)
            }
    |   class_constant_declaration ';'
            {
                $1.(*ast.StmtClassConstList).SemiColonTkn = $2
                $1.(*ast.StmtClassConstList).Node.Position = position.NewNodeTokenPosition($1, $2)
                $$ = $1
            }
    |   trait_use_statement
            {
                $$ = $1
            }
    |   method_modifiers function is_reference T_STRING '(' parameter_list ')' method_body
            {
                name := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($4),
                    },
                    IdentifierTkn: $4,
                    Value:         $4.Value,
                }
                $$ = &ast.StmtClassMethod{ast.Node{}, $3 != nil, name, $1, $6, nil, $8}

                // save position
                if $1 == nil {
                    $$.GetNode().Position = position.NewTokenNodePosition($2, $8)
                } else {
                    $$.GetNode().Position = position.NewNodeListNodePosition($1, $8)
                }

                // save comments
                if len($1) > 0 {
                    yylex.(*Parser).MoveFreeFloating($1[0], $$)
                    yylex.(*Parser).setFreeFloating($$, token.ModifierList, $2.SkippedTokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Start, $2.SkippedTokens)
                }
                if $3 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $4.SkippedTokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $3.SkippedTokens)
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $4.SkippedTokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.Name, $5.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ParameterList, $7.SkippedTokens)
            }
;

trait_use_statement:
        T_USE trait_list trait_adaptations
            {
                $$ = &ast.StmtTraitUse{ast.Node{}, $2, $3}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

trait_list:
        fully_qualified_class_name
            {
                $$ = []ast.Vertex{$1}
            }
    |   trait_list ',' fully_qualified_class_name
            {
                switch n := lastNode($1).(type) {
                    case *ast.NameName: n.ListSeparatorTkn = $2
                    case *ast.NameFullyQualified: n.ListSeparatorTkn = $2
                    case *ast.NameRelative: n.ListSeparatorTkn = $2
                }
                $$ = append($1, $3)
            }
;

trait_adaptations:
        ';'
            {
                $$ = &ast.StmtNop{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    SemiColonTkn: $1,
                }
            }
    |   '{' trait_adaptation_list '}'
            {
                $$ = &ast.StmtTraitAdaptationList{ast.Node{}, $2}

                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.AdaptationList, $3.SkippedTokens)
            }
;

trait_adaptation_list:
        /* empty */
            {
                $$ = nil
            }
    |   non_empty_trait_adaptation_list
            {
                $$ = $1
            }
;

non_empty_trait_adaptation_list:
        trait_adaptation_statement
            {
                $$ = []ast.Vertex{$1}
            }
    |   non_empty_trait_adaptation_list trait_adaptation_statement
            {
                $$ = append($1, $2)
            }
;

trait_adaptation_statement:
        trait_precedence ';'
            {
                $$ = $1;

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.NameList, $2.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.SkippedTokens)
            }
    |   trait_alias ';'
            {
                $$ = $1;

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Alias, $2.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.SkippedTokens)
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
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.SkippedTokens)
            }
;

trait_reference_list:
        fully_qualified_class_name
            {
                $$ = []ast.Vertex{$1}
            }
    |   trait_reference_list ',' fully_qualified_class_name
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
;

trait_method_reference:
        T_STRING
            {
                name := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                $$ = &ast.StmtTraitMethodRef{ast.Node{}, nil, name}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   trait_method_reference_fully_qualified
            {
                $$ = $1
            }
;

trait_method_reference_fully_qualified:
        fully_qualified_class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                $$ = &ast.StmtTraitMethodRef{ast.Node{}, $1, target}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
;

trait_alias:
        trait_method_reference T_AS trait_modifiers T_STRING
            {
                alias := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($4),
                    },
                    IdentifierTkn: $4,
                    Value:         $4.Value,
                }
                $$ = &ast.StmtTraitUseAlias{ast.Node{}, $1, $3, alias}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.SkippedTokens)
            }
    |   trait_method_reference T_AS member_modifier
            {
                $$ = &ast.StmtTraitUseAlias{ast.Node{}, $1, $3, nil}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.SkippedTokens)
            }
;

trait_modifiers:
        /* empty */
            {
                $$ = nil
            }
    |   member_modifier
            {
                $$ = $1
            }
;

method_body:
        ';' /* abstract method */
            {
                $$ = &ast.StmtNop{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    SemiColonTkn: $1,
                }
            }
    |   '{' inner_statement_list '}'
            {
                $$ = &ast.StmtStmtList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenCurlyBracket:  $1,
                    Stmts:             $2,
                    CloseCurlyBracket: $3,
                }
            }
;

variable_modifiers:
        non_empty_member_modifiers
            {
                $$ = $1;
            }
    |   T_VAR
            {
                $$ = []ast.Vertex{
                    &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
            }
;

method_modifiers:
        /* empty */
            {
                $$ = nil
            }
    |   non_empty_member_modifiers
            {
                $$ = $1
            }
;

non_empty_member_modifiers:
        member_modifier
            {
                $$ = []ast.Vertex{$1}
            }
    |   non_empty_member_modifiers member_modifier
            {
                $$ = append($1, $2)
            }
;

member_modifier:
        T_PUBLIC
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_PROTECTED
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_PRIVATE
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_STATIC
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_ABSTRACT
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_FINAL
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
;

class_variable_declaration:
        class_variable_declaration ',' T_VARIABLE
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                property := &ast.StmtProperty{ast.Node{}, variable, nil}
                $$ = append($1, property)

                // save position
                variable.GetNode().Position = position.NewTokenPosition($3)
                property.GetNode().Position = position.NewTokenPosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(property, token.Start, $3.SkippedTokens)
            }
    |   class_variable_declaration ',' T_VARIABLE '=' static_scalar
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                property := &ast.StmtProperty{ast.Node{}, variable, $5}
                $$ = append($1, property)

                // save position
                variable.GetNode().Position = position.NewTokenPosition($3)
                property.GetNode().Position = position.NewTokenNodePosition($3, $5)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(property, token.Start, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloating(property, token.Var, $4.SkippedTokens)
            }
    |   T_VARIABLE
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                property := &ast.StmtProperty{ast.Node{}, variable, nil}
                $$ = []ast.Vertex{property}

                // save position
                variable.GetNode().Position = position.NewTokenPosition($1)
                property.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(property, token.Start, $1.SkippedTokens)
            }
    |   T_VARIABLE '=' static_scalar
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                property := &ast.StmtProperty{ast.Node{}, variable, $3}
                $$ = []ast.Vertex{property}

                // save position
                variable.GetNode().Position = position.NewTokenPosition($1)
                property.GetNode().Position = position.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(property, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(property, token.Var, $2.SkippedTokens)
            }
;

class_constant_declaration:
        class_constant_declaration ',' T_STRING '=' static_scalar
            {
                constList := $1.(*ast.StmtClassConstList)
                constList.Node.Position = position.NewNodesPosition($1, $5)
                lastNode($$.(*ast.StmtClassConstList).Consts).(*ast.StmtConstant).CommaTkn = $2
                constList.Consts = append(constList.Consts, &ast.StmtConstant{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($3, $5),
                    },
                    Name: &ast.Identifier{
                        Node:  ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                    EqualTkn: $4,
                    Expr:     $5,
                })

                $$ = $1

                yylex.(*Parser).setFreeFloating(lastNode($$.(*ast.StmtClassConstList).Consts).(*ast.StmtConstant).Name, token.Start, $3.SkippedTokens)
            }
    |   T_CONST T_STRING '=' static_scalar
            {
                $$ = &ast.StmtClassConstList{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $4),
                    },
                    ConstTkn: $1,
                    Consts: []ast.Vertex{
                        &ast.StmtConstant{
                            Node: ast.Node{
                                Position: position.NewTokenNodePosition($2, $4),
                            },
                            Name: &ast.Identifier{
                                Node:  ast.Node{
                                    Position: position.NewTokenPosition($2),
                                },
                                IdentifierTkn: $2,
                                Value:         $2.Value,
                            },
                            EqualTkn: $3,
                            Expr:     $4,
                        },
                    },
                }

                yylex.(*Parser).setFreeFloating(lastNode($$.(*ast.StmtClassConstList).Consts).(*ast.StmtConstant).Name, token.Start, $2.SkippedTokens)
            }
;

echo_expr_list:
        echo_expr_list ',' expr
            {
                $1.(*ast.StmtEcho).Exprs = append($1.(*ast.StmtEcho).Exprs, $3)
                $1.(*ast.StmtEcho).SeparatorTkns = append($1.(*ast.StmtEcho).SeparatorTkns, $2)

                $$ = $1
            }
    |   expr
            {
                $$ = &ast.StmtEcho{
                    Exprs: []ast.Vertex{$1},
                }
            }
;


for_expr:
        /* empty */
            {
                $$ = nil
            }
    |   non_empty_for_expr
            {
                $$ = $1
            }
;

non_empty_for_expr:
        non_empty_for_expr ',' expr
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
    |   expr
            {
                $$ = []ast.Vertex{$1}
            }
;

chaining_method_or_property:
        chaining_method_or_property variable_property
            {
                $$ = append($1, $2...)
            }
    |   variable_property
            {
                $$ = $1
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
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Expr, $4.SkippedTokens)
            }
    |   '[' dim_offset ']'
            {
                fetch := &ast.ExprArrayDimFetch{ast.Node{}, nil, $2}
                $$ = []ast.Vertex{fetch}

                // save position
                fetch.GetNode().Position = position.NewNodePosition($2)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Var, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Expr, $3.SkippedTokens)
            }
;

chaining_instance_call:
        chaining_dereference chaining_method_or_property
            {
                $$ = append($1, $2...)
            }
    |   chaining_dereference
            {
                $$ = $1
            }
    |   chaining_method_or_property
            {
                $$ = $1
            }
;

instance_call:
        /* empty */
            {
                $$ = nil
            }
    |   chaining_instance_call
            {
                $$ = $1
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
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
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
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(listNode, token.List, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(listNode, token.ArrayPairList, $4.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $5.SkippedTokens)
            }
    |   variable '=' expr
            {
                $$ = &ast.ExprAssign{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable '=' '&' variable
            {
                $$ = &ast.ExprAssignReference{ast.Node{}, $1, $4}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Equal, $3.SkippedTokens)
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
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Equal, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloating(_new, token.Start, $4.SkippedTokens)
            }
    |   T_CLONE expr
            {
                $$ = &ast.ExprClone{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   variable T_PLUS_EQUAL expr
            {
                $$ = &ast.ExprAssignPlus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_MINUS_EQUAL expr
            {
                $$ = &ast.ExprAssignMinus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_MUL_EQUAL expr
            {
                $$ = &ast.ExprAssignMul{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_POW_EQUAL expr
            {
                $$ = &ast.ExprAssignPow{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_DIV_EQUAL expr
            {
                $$ = &ast.ExprAssignDiv{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_CONCAT_EQUAL expr
            {
                $$ = &ast.ExprAssignConcat{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_MOD_EQUAL expr
            {
                $$ = &ast.ExprAssignMod{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_AND_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_OR_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_XOR_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_SL_EQUAL expr
            {
                $$ = &ast.ExprAssignShiftLeft{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_SR_EQUAL expr
            {
                $$ = &ast.ExprAssignShiftRight{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   rw_variable T_INC
            {
                $$ = &ast.ExprPostInc{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   T_INC rw_variable
            {
                $$ = &ast.ExprPreInc{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   rw_variable T_DEC
            {
                $$ = &ast.ExprPostDec{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   T_DEC rw_variable
            {
                $$ = &ast.ExprPreDec{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   expr T_BOOLEAN_OR expr
            {
                $$ = &ast.ExprBinaryBooleanOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_BOOLEAN_AND expr
            {
                $$ = &ast.ExprBinaryBooleanAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_LOGICAL_OR expr
            {
                $$ = &ast.ExprBinaryLogicalOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_LOGICAL_AND expr
            {
                $$ = &ast.ExprBinaryLogicalAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_LOGICAL_XOR expr
            {
                $$ = &ast.ExprBinaryLogicalXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '|' expr
            {
                $$ = &ast.ExprBinaryBitwiseOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '&' expr
            {
                $$ = &ast.ExprBinaryBitwiseAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '^' expr
            {
                $$ = &ast.ExprBinaryBitwiseXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '.' expr
            {
                $$ = &ast.ExprBinaryConcat{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '+' expr
            {
                $$ = &ast.ExprBinaryPlus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '-' expr
            {
                $$ = &ast.ExprBinaryMinus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '*' expr
            {
                $$ = &ast.ExprBinaryMul{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_POW expr
            {
                $$ = &ast.ExprBinaryPow{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '/' expr
            {
                $$ = &ast.ExprBinaryDiv{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '%' expr
            {
                $$ = &ast.ExprBinaryMod{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_SL expr
            {
                $$ = &ast.ExprBinaryShiftLeft{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_SR expr
            {
                $$ = &ast.ExprBinaryShiftRight{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   '+' expr %prec T_INC
            {
                $$ = &ast.ExprUnaryPlus{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '-' expr %prec T_INC
            {
                $$ = &ast.ExprUnaryMinus{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '!' expr
            {
                $$ = &ast.ExprBooleanNot{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '~' expr
            {
                $$ = &ast.ExprBitwiseNot{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   expr T_IS_IDENTICAL expr
            {
                $$ = &ast.ExprBinaryIdentical{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_IS_NOT_IDENTICAL expr
            {
                $$ = &ast.ExprBinaryNotIdentical{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_IS_EQUAL expr
            {
                $$ = &ast.ExprBinaryEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_IS_NOT_EQUAL expr
            {
                $$ = &ast.ExprBinaryNotEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Equal, $2.SkippedTokens)
            }
    |   expr '<' expr
            {
                $$ = &ast.ExprBinarySmaller{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_IS_SMALLER_OR_EQUAL expr
            {
                $$ = &ast.ExprBinarySmallerOrEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '>' expr
            {
                $$ = &ast.ExprBinaryGreater{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_IS_GREATER_OR_EQUAL expr
            {
                $$ = &ast.ExprBinaryGreaterOrEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_INSTANCEOF class_name_reference
            {
                $$ = &ast.ExprInstanceOf{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   parenthesis_expr
            {
                $$ = $1
            }
    |   new_expr
            {
                $$ = $1
            }
    |   '(' new_expr ')' instance_call
            {
                $$ = $2

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, append($1.SkippedTokens, $$.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, append($$.GetNode().Tokens[token.End], $3.SkippedTokens...))

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
            }
    |   expr '?' expr ':' expr
            {
                $$ = &ast.ExprTernary{ast.Node{}, $1, $3, $5}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $5)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.True, $4.SkippedTokens)
            }
    |   expr '?' ':' expr
            {
                $$ = &ast.ExprTernary{ast.Node{}, $1, nil, $4}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.True, $3.SkippedTokens)
            }
    |   internal_functions_in_yacc
            {
                $$ = $1
            }
    |   T_INT_CAST expr
            {
                $$ = &ast.ExprCastInt{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_DOUBLE_CAST expr
            {
                $$ = &ast.ExprCastDouble{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_STRING_CAST expr
            {
                $$ = &ast.ExprCastString{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_ARRAY_CAST expr
            {
                $$ = &ast.ExprCastArray{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_OBJECT_CAST expr
            {
                $$ = &ast.ExprCastObject{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_BOOL_CAST expr
            {
                $$ = &ast.ExprCastBool{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_UNSET_CAST expr
            {
                $$ = &ast.ExprCastUnset{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_EXIT exit_expr
            {
                $$ = &ast.ExprExit{ast.Node{}, false, $2}

                if (bytes.EqualFold($1.Value, []byte("die"))) {
                    $$.(*ast.ExprExit).Die = true
                }

                // save position
                if $2 == nil {
                    $$.GetNode().Position = position.NewTokenPosition($1)
                } else {
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $2)
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '@' expr
            {
                $$ = &ast.ExprErrorSuppress{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   scalar
            {
                $$ = $1
            }
    |   combined_scalar_offset
            {
                $$ = $1
            }
    |   combined_scalar
            {
                $$ = $1
            }
    |   '`' backticks_expr '`'
            {
                $$ = &ast.ExprShellExec{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_PRINT expr
            {
                $$ = &ast.ExprPrint{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_YIELD
            {
                $$ = &ast.ExprYield{ast.Node{}, nil, nil}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   function is_reference '(' parameter_list ')' lexical_vars '{' inner_statement_list '}'
            {
                $$ = &ast.ExprClosure{ast.Node{}, $2 != nil, false, $4, $6, nil, $8}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $9)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                if $2 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $3.SkippedTokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $2.SkippedTokens)
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $3.SkippedTokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.ParameterList, $5.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.LexicalVars, $7.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $9.SkippedTokens)

                // normalize
                if $6 == nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Params, $$.GetNode().Tokens[token.LexicalVars]); delete($$.GetNode().Tokens, token.LexicalVars)
                }
            }
    |   T_STATIC function is_reference '(' parameter_list ')' lexical_vars '{' inner_statement_list '}'
            {
                $$ = &ast.ExprClosure{ast.Node{}, $3 != nil, true, $5, $7, nil, $9}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $10)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Static, $2.SkippedTokens)
                if $3 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $4.SkippedTokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $3.SkippedTokens)
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $4.SkippedTokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.ParameterList, $6.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.LexicalVars, $8.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $10.SkippedTokens)

                // normalize
                if $7 == nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Params, $$.GetNode().Tokens[token.LexicalVars]); delete($$.GetNode().Tokens, token.LexicalVars)
                }
            }
;

yield_expr:
        T_YIELD expr_without_variable
            {
                $$ = &ast.ExprYield{ast.Node{}, nil, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_YIELD variable
            {
                $$ = &ast.ExprYield{ast.Node{}, nil, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_YIELD expr T_DOUBLE_ARROW expr_without_variable
            {
                $$ = &ast.ExprYield{ast.Node{}, $2, $4}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.SkippedTokens)
            }
    |   T_YIELD expr T_DOUBLE_ARROW variable
            {
                $$ = &ast.ExprYield{ast.Node{}, $2, $4}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.SkippedTokens)
            }
;

combined_scalar_offset:
        combined_scalar '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   combined_scalar_offset '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   T_CONSTANT_ENCAPSED_STRING '[' dim_offset ']'
            {
                str := &ast.ScalarString{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    StringTkn: $1,
                    Value:     $1.Value,
                }
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, str, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition(str, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   general_constant '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
;

combined_scalar:
        T_ARRAY '(' array_pair_list ')'
            {
                $$ = &ast.ExprArray{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Array, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $4.SkippedTokens)
            }
    |   '[' array_pair_list ']'
            {
                $$ = &ast.ExprShortArray{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $3.SkippedTokens)
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
            }
    |   T_USE '(' lexical_var_list ')'
            {
                $$ = &ast.ExprClosureUse{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Use, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.LexicalVarList, $4.SkippedTokens)
            }
;

lexical_var_list:
        lexical_var_list ',' T_VARIABLE
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = append($1, variable)

                // save position
                variable.GetNode().Position = position.NewTokenPosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $3.SkippedTokens)
            }
    |   lexical_var_list ',' '&' T_VARIABLE
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($4),
                    },
                    IdentifierTkn: $4,
                    Value:         $4.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                reference := &ast.ExprReference{ast.Node{}, variable}
                $$ = append($1, reference)

                // save position
                variable.GetNode().Position = position.NewTokenPosition($4)
                reference.GetNode().Position = position.NewTokensPosition($3, $4)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(reference, token.Start, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $4.SkippedTokens)
            }
    |   T_VARIABLE
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = []ast.Vertex{variable}

                // save position
                variable.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(variable, token.Start, $1.SkippedTokens)
            }
    |   '&' T_VARIABLE
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($2),
                    },
                    IdentifierTkn: $2,
                    Value:         $2.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                reference := &ast.ExprReference{ast.Node{}, variable}
                $$ = []ast.Vertex{reference}

                // save position
                variable.GetNode().Position = position.NewTokenPosition($2)
                reference.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating(reference, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $2.SkippedTokens)
            }
;

function_call:
        namespace_name function_call_parameter_list
            {
                name := &ast.NameName{
                    Node:  ast.Node{
                        Position: position.NewNodeListPosition($1),
                    },
                    Parts: $1,
                }
                $$ = &ast.ExprFunctionCall{ast.Node{}, name, $2.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition(name, $2)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name function_call_parameter_list
            {
                name := &ast.NameRelative{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $3),
                    },
                    NsTkn:          $1,
                    NsSeparatorTkn: $2,
                    Parts:          $3,
                }
                $$ = &ast.ExprFunctionCall{ast.Node{}, name, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition(name, $4)
            }
    |   T_NS_SEPARATOR namespace_name function_call_parameter_list
            {
                name := &ast.NameFullyQualified{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $2),
                    },
                    NsSeparatorTkn: $1,
                    Parts:          $2,
                }
                $$ = &ast.ExprFunctionCall{ast.Node{}, name, $3.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition(name, $3)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM variable_name function_call_parameter_list
            {
                $$ = &ast.ExprStaticCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects function_call_parameter_list
            {
                $$ = &ast.ExprStaticCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_name function_call_parameter_list
            {
                $$ = &ast.ExprStaticCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects function_call_parameter_list
            {
                $$ = &ast.ExprStaticCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
    |   variable_without_objects function_call_parameter_list
            {
                $$ = &ast.ExprFunctionCall{ast.Node{}, $1, $2.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
            }
;

class_name:
        T_STATIC
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   namespace_name
            {
                $$ = &ast.NameName{
                    Node:  ast.Node{
                        Position: position.NewNodeListPosition($1),
                    },
                    Parts: $1,
                }
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameRelative{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $3),
                    },
                    NsTkn:          $1,
                    NsSeparatorTkn: $2,
                    Parts:          $3,
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameFullyQualified{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $2),
                    },
                    NsSeparatorTkn: $1,
                    Parts:          $2,
                }
            }
;

fully_qualified_class_name:
        namespace_name
            {
                $$ = &ast.NameName{
                    Node:  ast.Node{
                        Position: position.NewNodeListPosition($1),
                    },
                    Parts: $1,
                }
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameRelative{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $3),
                    },
                    NsTkn:          $1,
                    NsSeparatorTkn: $2,
                    Parts:          $3,
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameFullyQualified{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $2),
                    },
                    NsSeparatorTkn: $1,
                    Parts:          $2,
                }
            }
;

class_name_reference:
        class_name
            {
                $$ = $1
            }
    |   dynamic_class_name_reference
            {
                $$ = $1
            }
;

dynamic_class_name_reference:
        base_variable T_OBJECT_OPERATOR object_property dynamic_class_name_variable_properties
            {
                $$ = $1

                // save comments
                yylex.(*Parser).setFreeFloating($3[0], token.Var, $2.SkippedTokens)

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
            }
    |   base_variable
            {
                $$ = $1
            }
;


dynamic_class_name_variable_properties:
        dynamic_class_name_variable_properties dynamic_class_name_variable_property
            {
                $$ = append($1, $2...)
            }
    |   /* empty */
            {
                $$ = []ast.Vertex{}
            }
;


dynamic_class_name_variable_property:
        T_OBJECT_OPERATOR object_property
            {
                $$ = $2

                // save comments
                yylex.(*Parser).setFreeFloating($2[0], token.Var, $1.SkippedTokens)
            }
;

exit_expr:
        /* empty */
            {
                $$ = nil
            }
    |   '(' ')'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    OpenBracketTkn:  $1,
                    CloseBracketTkn: $2,
                }

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $2.SkippedTokens)
            }
    |   parenthesis_expr
            {
                $$ = $1
            }
;

backticks_expr:
        /* empty */
            {
                $$ = []ast.Vertex{}
            }
    |   T_ENCAPSED_AND_WHITESPACE
            {
                $$ = []ast.Vertex{
                    &ast.ScalarEncapsedStringPart{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        EncapsedStrTkn: $1,
                        Value:          $1.Value,
                    },
                }
            }
    |   encaps_list
            {
                $$ = $1;
            }
;

ctor_arguments:
        /* empty */
            {
                $$ = nil
            }
    |   function_call_parameter_list
            {
                $$ = $1
            }
;

common_scalar:
        T_LNUMBER
            {
                $$ = &ast.ScalarLnumber{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    NumberTkn: $1,
                    Value:     $1.Value,
                }
            }
    |   T_DNUMBER
            {
                $$ = &ast.ScalarDnumber{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    NumberTkn: $1,
                    Value:     $1.Value,
                }
            }
    |   T_CONSTANT_ENCAPSED_STRING
            {
                $$ = &ast.ScalarString{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    StringTkn: $1,
                    Value:     $1.Value,
                }
            }
    |   T_LINE
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_FILE
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_DIR
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_TRAIT_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_METHOD_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_FUNC_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_NS_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenHeredocTkn: $1,
                    Parts: []ast.Vertex{
                        &ast.ScalarEncapsedStringPart{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($2),
                            },
                            EncapsedStrTkn: $2,
                            Value:          $2.Value,
                        },
                    },
                    CloseHeredocTkn: $3,
                }
            }
    |   T_START_HEREDOC T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    OpenHeredocTkn:  $1,
                    CloseHeredocTkn: $2,
                }
            }
;

static_class_constant:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
;

static_scalar:
        static_scalar_value
            {
                $$ = $1
            }
;

static_scalar_value:
        common_scalar
            {
                $$ = $1
            }
    |   static_class_name_scalar
            {
                $$ = $1
            }
    |   namespace_name
            {
                name := &ast.NameName{
                    Node:  ast.Node{
                        Position: position.NewNodeListPosition($1),
                    },
                    Parts: $1,
                }
                $$ = &ast.ExprConstFetch{ast.Node{}, name}

                // save position
                $$.GetNode().Position = position.NewNodePosition(name)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                name := &ast.NameRelative{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $3),
                    },
                    NsTkn:          $1,
                    NsSeparatorTkn: $2,
                    Parts:          $3,
                }
                $$ = &ast.ExprConstFetch{ast.Node{}, name}

                // save position
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $3)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := &ast.NameFullyQualified{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $2),
                    },
                    NsSeparatorTkn: $1,
                    Parts:          $2,
                }
                $$ = &ast.ExprConstFetch{ast.Node{}, name}

                // save position
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $2)
            }
    |   T_ARRAY '(' static_array_pair_list ')'
            {
                $$ = &ast.ExprArray{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Array, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $4.SkippedTokens)
            }
    |   '[' static_array_pair_list ']'
            {
                $$ = &ast.ExprShortArray{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $3.SkippedTokens)
            }
    |   static_class_constant
            {
                $$ = $1
            }
    |   T_CLASS_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   static_operation
            {
                $$ = $1
            }
;

static_operation:
        static_scalar_value '[' static_scalar_value ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   static_scalar_value '+' static_scalar_value
            {
                $$ = &ast.ExprBinaryPlus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value '-' static_scalar_value
            {
                $$ = &ast.ExprBinaryMinus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value '*' static_scalar_value
            {
                $$ = &ast.ExprBinaryMul{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_POW static_scalar_value
            {
                $$ = &ast.ExprBinaryPow{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value '/' static_scalar_value
            {
                $$ = &ast.ExprBinaryDiv{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value '%' static_scalar_value
            {
                $$ = &ast.ExprBinaryMod{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   '!' static_scalar_value
            {
                $$ = &ast.ExprBooleanNot{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '~' static_scalar_value
            {
                $$ = &ast.ExprBitwiseNot{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   static_scalar_value '|' static_scalar_value
            {
                $$ = &ast.ExprBinaryBitwiseOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value '&' static_scalar_value
            {
                $$ = &ast.ExprBinaryBitwiseAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value '^' static_scalar_value
            {
                $$ = &ast.ExprBinaryBitwiseXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_SL static_scalar_value
            {
                $$ = &ast.ExprBinaryShiftLeft{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_SR static_scalar_value
            {
                $$ = &ast.ExprBinaryShiftRight{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value '.' static_scalar_value
            {
                $$ = &ast.ExprBinaryConcat{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_LOGICAL_XOR static_scalar_value
            {
                $$ = &ast.ExprBinaryLogicalXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_LOGICAL_AND static_scalar_value
            {
                $$ = &ast.ExprBinaryLogicalAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_LOGICAL_OR static_scalar_value
            {
                $$ = &ast.ExprBinaryLogicalOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_BOOLEAN_AND static_scalar_value
            {
                $$ = &ast.ExprBinaryBooleanAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_BOOLEAN_OR static_scalar_value
            {
                $$ = &ast.ExprBinaryBooleanOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_IS_IDENTICAL static_scalar_value
            {
                $$ = &ast.ExprBinaryIdentical{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_IS_NOT_IDENTICAL static_scalar_value
            {
                $$ = &ast.ExprBinaryNotIdentical{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_IS_EQUAL static_scalar_value
            {
                $$ = &ast.ExprBinaryEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_IS_NOT_EQUAL static_scalar_value
            {
                $$ = &ast.ExprBinaryNotEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Equal, $2.SkippedTokens)
            }
    |   static_scalar_value '<' static_scalar_value
            {
                $$ = &ast.ExprBinarySmaller{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value '>' static_scalar_value
            {
                $$ = &ast.ExprBinaryGreater{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_IS_SMALLER_OR_EQUAL static_scalar_value
            {
                $$ = &ast.ExprBinarySmallerOrEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value T_IS_GREATER_OR_EQUAL static_scalar_value
            {
                $$ = &ast.ExprBinaryGreaterOrEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value '?' ':' static_scalar_value
            {
                $$ = &ast.ExprTernary{ast.Node{}, $1, nil, $4}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.True, $3.SkippedTokens)
            }
    |   static_scalar_value '?' static_scalar_value ':' static_scalar_value
            {
                $$ = &ast.ExprTernary{ast.Node{}, $1, $3, $5}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $5)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.True, $4.SkippedTokens)
            }
    |   '+' static_scalar_value
            {
                $$ = &ast.ExprUnaryPlus{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '-' static_scalar_value
            {
                $$ = &ast.ExprUnaryMinus{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '(' static_scalar_value ')'
            {
                $$ = $2

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, append($1.SkippedTokens, $$.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, append($$.GetNode().Tokens[token.End], $3.SkippedTokens...))
            }
;

general_constant:
        class_constant
            {
                $$ = $1
            }
    |   namespace_name
            {
                name := &ast.NameName{
                    Node:  ast.Node{
                        Position: position.NewNodeListPosition($1),
                    },
                    Parts: $1,
                }
                $$ = &ast.ExprConstFetch{ast.Node{}, name}

                // save position
                $$.GetNode().Position = position.NewNodePosition(name)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                name := &ast.NameRelative{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $3),
                    },
                    NsTkn:          $1,
                    NsSeparatorTkn: $2,
                    Parts:          $3,
                }
                $$ = &ast.ExprConstFetch{ast.Node{}, name}

                // save position
                $$.GetNode().Position = position.NewNodePosition(name)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := &ast.NameFullyQualified{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $2),
                    },
                    NsSeparatorTkn: $1,
                    Parts:          $2,
                }
                $$ = &ast.ExprConstFetch{ast.Node{}, name}

                // save position
                $$.GetNode().Position = position.NewNodePosition(name)
            }
;

scalar:
        T_STRING_VARNAME
            {
                name := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                $$ = &ast.ExprVariable{ast.Node{}, name}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   general_constant
            {
                $$ = $1
            }
    |   class_name_scalar
            {
                $$ = $1
            }
    |   common_scalar
            {
                $$ = $1
            }
    |   '"' encaps_list '"'
            {
                $$ = &ast.ScalarEncapsed{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenQoteTkn:  $1,
                    Parts:        $2,
                    CloseQoteTkn: $1,
                }
            }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenHeredocTkn:  $1,
                    Parts:           $2,
                    CloseHeredocTkn: $3,
                }
            }
    |   T_CLASS_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
;

static_array_pair_list:
        /* empty */
            {
                $$ = nil
            }
    |   non_empty_static_array_pair_list possible_comma
            {
                $$ = $1

                // save comments
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
                }
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
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Expr, $4.SkippedTokens)
            }
    |   non_empty_static_array_pair_list ',' static_scalar_value
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, nil, $3}
                $$ = append($1, arrayItem)

                // save position
                arrayItem.GetNode().Position = position.NewNodePosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)
            }
    |   static_scalar_value T_DOUBLE_ARROW static_scalar_value
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, $1, $3}
                $$ = []ast.Vertex{arrayItem}

                // save position
                arrayItem.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Expr, $2.SkippedTokens)
            }
    |   static_scalar_value
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, nil, $1}
                $$ = []ast.Vertex{arrayItem}

                // save position
                arrayItem.GetNode().Position = position.NewNodePosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)
            }
;

expr:
        r_variable
            {
                $$ = $1
            }
    |   expr_without_variable
            {
                $$ = $1
            }
;

parenthesis_expr:
        '(' expr ')'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn:  $1,
                    Child:           $2,
                    CloseBracketTkn: $3,
                }

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.SkippedTokens)
            }
    |   '(' yield_expr ')'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn:  $1,
                    Child:           $2,
                    CloseBracketTkn: $3,
                }

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.SkippedTokens)
            }
;


r_variable:
        variable
            {
                $$ = $1
            }
;


w_variable:
        variable
            {
                $$ = $1
            }
;

rw_variable:
        variable
            {
                $$ = $1
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
                yylex.(*Parser).setFreeFloating($3[0], token.Var, $2.SkippedTokens)

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
            }
    |   base_variable_with_function_calls
            {
                $$ = $1
            }
;

variable_properties:
        variable_properties variable_property
            {
                $$ = append($1, $2...)
            }
    |   /* empty */
            {
                $$ = []ast.Vertex{}
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
                yylex.(*Parser).setFreeFloating($2[0], token.Var, $1.SkippedTokens)
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
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Expr, $4.SkippedTokens)
            }
    |   method '[' dim_offset ']'
            {
                fetch := &ast.ExprArrayDimFetch{ast.Node{}, nil, $3}
                $$ = []ast.Vertex{$1, fetch}

                // save position
                fetch.GetNode().Position = position.NewNodePosition($3)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Expr, $4.SkippedTokens)
            }
;

method:
        function_call_parameter_list
            {
                $$ = &ast.ExprMethodCall{ast.Node{}, nil, nil, $1.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)
            }
;

method_or_not:
        method
            {
                $$ = []ast.Vertex{$1}
            }
    |   array_method_dereference
            {
                $$ = $1
            }
    |   /* empty */
            {
                $$ = nil
            }
;

variable_without_objects:
        reference_variable
            {
                $$ = $1
            }
    |   simple_indirect_reference reference_variable
            {
                $1.last.VarName = $2

                for _, n := range($1.all) {
                    n.GetNode().Position = position.NewNodesPosition(n, $2)
                }

                $$ = $1.all[0]
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
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects
            {
                $$ = &ast.ExprStaticPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
;

variable_class_name:
        reference_variable
            {
                $$ = $1
            }
;

array_function_dereference:
        array_function_dereference '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   function_call '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
;

base_variable_with_function_calls:
        base_variable
            {
                $$ = $1
            }
    |   array_function_dereference
            {
                $$ = $1
            }
    |   function_call
            {
                $$ = $1
            }
;


base_variable:
        reference_variable
            {
                $$ = $1
            }
    |   simple_indirect_reference reference_variable
            {
                $1.last.VarName = $2

                for _, n := range($1.all) {
                    n.GetNode().Position = position.NewNodesPosition(n, $2)
                }

                $$ = $1.all[0]
            }
    |   static_member
            {
                $$ = $1
            }
;

reference_variable:
        reference_variable '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   reference_variable '{' expr '}'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   compound_variable
            {
                $$ = $1
            }
;


compound_variable:
        T_VARIABLE
            {
                name := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                $$ = &ast.ExprVariable{ast.Node{}, name}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '$' '{' expr '}'
            {
                $$ = &ast.ExprVariable{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($3, token.Start, append($2.SkippedTokens, $3.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($3, token.End, append($3.GetNode().Tokens[token.End], $4.SkippedTokens...))
            }
;

dim_offset:
        /* empty */
            {
                $$ = nil
            }
    |   expr
            {
                $$ = $1
            }
;


object_property:
        object_dim_list
            {
                $$ = $1
            }
    |   variable_without_objects
            {
                fetch := &ast.ExprPropertyFetch{ast.Node{}, nil, $1}
                $$ = []ast.Vertex{fetch}

                // save position
                fetch.GetNode().Position = position.NewNodePosition($1)
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
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Expr, $4.SkippedTokens)
            }
    |   object_dim_list '{' expr '}'
            {
                fetch := &ast.ExprArrayDimFetch{ast.Node{}, nil, $3}
                $$ = append($1, fetch)

                // save position
                fetch.GetNode().Position = position.NewNodePosition($3)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(fetch, token.Expr, $4.SkippedTokens)
            }
    |   variable_name
            {
                fetch := &ast.ExprPropertyFetch{ast.Node{}, nil, $1}
                $$ = []ast.Vertex{fetch}

                // save position
                fetch.GetNode().Position = position.NewNodePosition($1)
            }
;

variable_name:
        T_STRING
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   '{' expr '}'
            {
                $$ = $2

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, append($1.SkippedTokens, $$.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, append($$.GetNode().Tokens[token.End], $3.SkippedTokens...))
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
                yylex.(*Parser).setFreeFloating(n, token.Start, $1.SkippedTokens)
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
                yylex.(*Parser).setFreeFloating(n, token.Start, $2.SkippedTokens)
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
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
    |   assignment_list_element
            {
                if $1.(*ast.ExprArrayItem).Key == nil && $1.(*ast.ExprArrayItem).Val == nil {
                    $$ = []ast.Vertex{}
                } else {
                    $$ = []ast.Vertex{$1}
                }
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
            }
    |   T_LIST '(' assignment_list ')'
            {
                listNode := &ast.ExprList{ast.Node{}, $3}
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, listNode}

                // save position
                listNode.GetNode().Position = position.NewTokensPosition($1, $4)
                $$.GetNode().Position = position.NewNodePosition(listNode)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(listNode, token.List, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(listNode, token.ArrayPairList, $4.SkippedTokens)
            }
    |   /* empty */
            {
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, nil}
            }
;


array_pair_list:
        /* empty */
            {
                $$ = []ast.Vertex{}
            }
    |   non_empty_array_pair_list possible_comma
            {
                $$ = $1

                if $2 != nil {
                    $$ = append($1, &ast.ExprArrayItem{ast.Node{}, false, nil, nil})
                }

                // save comments
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
                }
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
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Expr, $4.SkippedTokens)
            }
    |   non_empty_array_pair_list ',' expr
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, nil, $3}
                $$ = append($1, arrayItem)

                // save position
                arrayItem.GetNode().Position = position.NewNodePosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)
            }
    |   expr T_DOUBLE_ARROW expr
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, $1, $3}
                $$ = []ast.Vertex{arrayItem}

                // save position
                arrayItem.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Expr, $2.SkippedTokens)
            }
    |   expr
            {
                arrayItem := &ast.ExprArrayItem{ast.Node{}, false, nil, $1}
                $$ = []ast.Vertex{arrayItem}

                // save position
                arrayItem.GetNode().Position = position.NewNodePosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)
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
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Expr, $4.SkippedTokens)
                yylex.(*Parser).setFreeFloating(reference, token.Start, $5.SkippedTokens)
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
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(arrayItem, token.Start, $3.SkippedTokens)
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
                yylex.(*Parser).setFreeFloating(arrayItem, token.Expr, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(reference, token.Start, $3.SkippedTokens)
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
                yylex.(*Parser).setFreeFloating(arrayItem, token.Start, $1.SkippedTokens)
            }
;

encaps_list:
        encaps_list encaps_var
            {
                $$ = append($1, $2)
            }
    |   encaps_list T_ENCAPSED_AND_WHITESPACE
            {
                $$ = append(
                    $1,
                    &ast.ScalarEncapsedStringPart{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($2),
                        },
                        EncapsedStrTkn: $2,
                        Value:          $2.Value,
                    },
                )
            }
    |   encaps_var
            {
                $$ = []ast.Vertex{$1}
            }
    |   T_ENCAPSED_AND_WHITESPACE encaps_var
            {
                $$ = []ast.Vertex{
                    &ast.ScalarEncapsedStringPart{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        EncapsedStrTkn: $1,
                        Value:          $1.Value,
                    },
                    $2,
                }
            }
;

encaps_var:
        T_VARIABLE
            {
                name := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                $$ = &ast.ExprVariable{ast.Node{}, name}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_VARIABLE '[' encaps_var_offset ']'
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, variable, $3}

                // save position
                variable.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                fetch := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                $$ = &ast.ExprPropertyFetch{ast.Node{}, variable, fetch}

                // save position
                variable.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES expr '}'
            {
                variable := &ast.ExprVariable{ast.Node{}, $2}

                $$ = variable

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setToken($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.SkippedTokens)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
            {
                name := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($2),
                    },
                    IdentifierTkn: $2,
                    Value:         $2.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, name}

                $$ = variable

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setToken($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.SkippedTokens)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($2),
                    },
                    IdentifierTkn: $2,
                    Value:         $2.Value,
                }
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, variable, $4}

                // save position
                variable.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $6)

                // save comments
                yylex.(*Parser).setToken(variable, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $5.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $6.SkippedTokens)
            }
    |   T_CURLY_OPEN variable '}'
            {
                $$ = $2;

                // save comments
                yylex.(*Parser).setToken($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.SkippedTokens)
            }
;

encaps_var_offset:
        T_STRING
            {
                $$ = &ast.ScalarString{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    StringTkn: $1,
                    Value:     $1.Value,
                }
            }
    |   T_NUM_STRING
            {
                // TODO: add option to handle 64 bit integer
                if _, err := strconv.Atoi(string($1.Value)); err == nil {
                    $$ = &ast.ScalarLnumber{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        NumberTkn: $1,
                        Value:     $1.Value,
                    }
                } else {
                    $$ = &ast.ScalarString{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        StringTkn: $1,
                        Value:     $1.Value,
                    }
                }

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_VARIABLE
            {
                identifier := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
                $$ = &ast.ExprVariable{ast.Node{}, identifier}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

internal_functions_in_yacc:
        T_ISSET '(' isset_variables ')'
            {
                $$ = &ast.ExprIsset{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Isset, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.VarList, $4.SkippedTokens)
            }
    |   T_EMPTY '(' variable ')'
            {
                exprBrackets := &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($2, $4),
                    },
                    OpenBracketTkn:  $2,
                    Child:           $3,
                    CloseBracketTkn: $4,
                }
                $$ = &ast.ExprEmpty{ast.Node{}, exprBrackets}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(exprBrackets, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(exprBrackets, token.End, $4.SkippedTokens)
            }
    |   T_EMPTY '(' expr ')'
            {
                exprBrackets := &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($2, $4),
                    },
                    OpenBracketTkn:  $2,
                    Child:           $3,
                    CloseBracketTkn: $4,
                }
                $$ = &ast.ExprEmpty{ast.Node{}, exprBrackets}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(exprBrackets, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(exprBrackets, token.End, $4.SkippedTokens)
            }
    |   T_INCLUDE expr
            {
                $$ = &ast.ExprInclude{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_INCLUDE_ONCE expr
            {
                $$ = &ast.ExprIncludeOnce{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_EVAL '(' expr ')'
            {
                exprBrackets := &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($2, $4),
                    },
                    OpenBracketTkn:  $2,
                    Child:           $3,
                    CloseBracketTkn: $4,
                }
                $$ = &ast.ExprEval{ast.Node{}, exprBrackets}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(exprBrackets, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(exprBrackets, token.End, $4.SkippedTokens)
            }
    |   T_REQUIRE expr
            {
                $$ = &ast.ExprRequire{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_REQUIRE_ONCE expr
            {
                $$ = &ast.ExprRequireOnce{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

isset_variables:
        isset_variable
            {
                $$ = []ast.Vertex{$1}
            }
    |   isset_variables ',' isset_variable
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
;

isset_variable:
        variable
            {
                $$ = $1
            }
    |   expr_without_variable
            {
                $$ = $1
            }
;

class_constant:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
;

static_class_name_scalar:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS
            {
                target := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
;

class_name_scalar:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS
            {
                target := &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    IdentifierTkn: $3,
                    Value:         $3.Value,
                }
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
;

%%

type simpleIndirectReference struct {
	all  []*ast.ExprVariable
	last *ast.ExprVariable
}
