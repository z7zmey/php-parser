%{
package php5

import (
    "strconv"

    "github.com/z7zmey/php-parser/pkg/ast"
    "github.com/z7zmey/php-parser/pkg/errors"
    "github.com/z7zmey/php-parser/pkg/token"
)

%}

%union{
    node                    ast.Vertex
    token                   *token.Token
    list                    []ast.Vertex
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
%type <token> is_reference is_variadic

%type <node> top_statement use_declaration use_function_declaration use_const_declaration common_scalar
%type <node> static_class_constant compound_variable reference_variable class_name variable_class_name
%type <node> dim_offset expr expr_without_variable r_variable w_variable rw_variable variable base_variable_with_function_calls
%type <node> base_variable array_function_dereference function_call inner_statement statement unticked_statement
%type <node> statement global_var static_scalar scalar class_constant static_class_name_scalar class_name_scalar
%type <node> encaps_var encaps_var encaps_var_offset general_constant isset_variable internal_functions_in_yacc assignment_list_element
%type <node> variable_name variable_without_objects dynamic_class_name_reference new_expr class_name_reference static_member
%type <node> function_call fully_qualified_class_name combined_scalar combined_scalar_offset general_constant parenthesis_expr
%type <node> exit_expr yield_expr function_declaration_statement class_declaration_statement constant_declaration
%type <node> else_single new_else_single unset_variable declare_statement parameter_list non_empty_parameter_list
%type <node> finally_statement additional_catch unticked_function_declaration_statement unticked_class_declaration_statement
%type <node> optional_class_type parameter class_entry_type class_statement class_constant_declaration
%type <node> trait_use_statement function_call_parameter trait_adaptation_statement trait_precedence trait_alias
%type <node> trait_method_reference_fully_qualified trait_method_reference trait_modifiers member_modifier method
%type <node> static_scalar_value static_operation static_var_list global_var_list
%type <node> ctor_arguments function_call_parameter_list echo_expr_list class_variable_declaration
%type <node> trait_adaptations unset_variables declare_list non_empty_array_pair_list array_pair_list
%type <node> switch_case_list non_empty_function_call_parameter_list assignment_list lexical_var_list
%type <node> method_body trait_reference_list static_array_pair_list non_empty_static_array_pair_list
%type <node> foreach_statement for_statement while_statement isset_variables
%type <node> foreach_variable foreach_optional_arg for_expr non_empty_for_expr
%type <node> extends_from interface_list trait_list namespace_name
%type <node> implements_list use_declarations use_function_declarations use_const_declarations
%type <node> interface_extends_list
%type <node> lexical_vars

%type <list> top_statement_list
%type <list> inner_statement_list encaps_list
%type <list> elseif_list new_elseif_list
%type <list> case_list catch_statement additional_catches
%type <list> non_empty_additional_catches class_statement_list
%type <list> class_statement_list variable_modifiers method_modifiers
%type <list> trait_adaptation_list non_empty_trait_adaptation_list
%type <list> non_empty_member_modifiers backticks_expr

%type <list> chaining_dereference chaining_instance_call chaining_method_or_property instance_call variable_property
%type <list> method_or_not array_method_dereference object_property object_dim_list dynamic_class_name_variable_property
%type <list> dynamic_class_name_variable_properties variable_properties

%type <list> simple_indirect_reference

%%

start:
        top_statement_list
            {
                yylex.(*Parser).currentToken.Value = nil

                yylex.(*Parser).rootNode = &ast.Root{
                    Position: yylex.(*Parser).builder.NewNodeListPosition($1),
                    Stmts:  $1,
                    EndTkn: yylex.(*Parser).currentToken,
                }
            }
;

top_statement_list:
        top_statement_list top_statement
            {
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
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{
                        &ast.NamePart{
                            Position: yylex.(*Parser).builder.NewTokenPosition($1),
                            StringTkn: $1,
                            Value:     $1.Value,
                        },
                    },
                }
            }
    |   namespace_name T_NS_SEPARATOR T_STRING
            {
                part := &ast.NamePart{
                    Position: yylex.(*Parser).builder.NewTokenPosition($3),
                    StringTkn:      $3,
                    Value:          $3.Value,
                }

                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, part)

                $$ = $1
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
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    HaltCompilerTkn:     $1,
                    OpenParenthesisTkn:  $2,
                    CloseParenthesisTkn: $3,
                    SemiColonTkn:        $4,
                }
            }
    |   T_NAMESPACE namespace_name ';'
            {
                $$ = &ast.StmtNamespace{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    NsTkn: $1,
                    Name: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($2.(*ParserSeparatedList).Items),
                        Parts:         $2.(*ParserSeparatedList).Items,
                        SeparatorTkns: $2.(*ParserSeparatedList).SeparatorTkns,
                    },
                    SemiColonTkn: $3,
                }
            }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
            {
                $$ = &ast.StmtNamespace{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $5),
                    NsTkn: $1,
                    Name: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($2.(*ParserSeparatedList).Items),
                        Parts:         $2.(*ParserSeparatedList).Items,
                        SeparatorTkns: $2.(*ParserSeparatedList).SeparatorTkns,
                    },
                    OpenCurlyBracketTkn:  $3,
                    Stmts:                $4,
                    CloseCurlyBracketTkn: $5,
                }
            }
    |   T_NAMESPACE '{' top_statement_list '}'
            {
                $$ = &ast.StmtNamespace{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    NsTkn:                $1,
                    OpenCurlyBracketTkn:  $2,
                    Stmts:                $3,
                    CloseCurlyBracketTkn: $4,
                }
            }
    |   T_USE use_declarations ';'
            {
                $$ = &ast.StmtUseList{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    UseTkn:          $1,
                    Uses:            $2.(*ParserSeparatedList).Items,
                    SeparatorTkns:   $2.(*ParserSeparatedList).SeparatorTkns,
                    SemiColonTkn:    $3,
                }
            }
    |   T_USE T_FUNCTION use_function_declarations ';'
            {
                $$ = &ast.StmtUseList{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    UseTkn: $1,
                    Type: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($2),
                        IdentifierTkn: $2,
                        Value:         $2.Value,
                    },
                    Uses:            $3.(*ParserSeparatedList).Items,
                    SeparatorTkns:   $3.(*ParserSeparatedList).SeparatorTkns,
                    SemiColonTkn:    $4,
                }
            }
    |   T_USE T_CONST use_const_declarations ';'
            {
                $$ = &ast.StmtUseList{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    UseTkn: $1,
                    Type: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($2),
                        IdentifierTkn: $2,
                        Value:         $2.Value,
                    },
                    Uses:            $3.(*ParserSeparatedList).Items,
                    SeparatorTkns:   $3.(*ParserSeparatedList).SeparatorTkns,
                    SemiColonTkn:    $4,
                }
            }
    |   constant_declaration ';'
            {
                $1.(*ast.StmtConstList).SemiColonTkn = $2
                $1.(*ast.StmtConstList).Position = yylex.(*Parser).builder.NewNodeTokenPosition($1, $2)
                $$ = $1
            }
;

use_declarations:
        use_declarations ',' use_declaration
            {
                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   use_declaration
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
;

use_declaration:
        namespace_name
            {
                $$ = &ast.StmtUse{
                    Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                    Use: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                        Parts:         $1.(*ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
    |   namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUse{
                    Position: yylex.(*Parser).builder.NewNodeListTokenPosition($1.(*ParserSeparatedList).Items, $3),
                    Use: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                        Parts:         $1.(*ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ParserSeparatedList).SeparatorTkns,
                    },
                    AsTkn: $2,
                    Alias: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.StmtUse{
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $2.(*ParserSeparatedList).Items),
                    NsSeparatorTkn: $1,
                    Use: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($2.(*ParserSeparatedList).Items),
                        Parts:         $2.(*ParserSeparatedList).Items,
                        SeparatorTkns: $2.(*ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUse{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    NsSeparatorTkn: $1,
                    Use: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($2.(*ParserSeparatedList).Items),
                        Parts:         $2.(*ParserSeparatedList).Items,
                        SeparatorTkns: $2.(*ParserSeparatedList).SeparatorTkns,
                    },
                    AsTkn: $3,
                    Alias: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($4),
                        IdentifierTkn: $4,
                        Value:         $4.Value,
                    },
                }
            }
;

use_function_declarations:
        use_function_declarations ',' use_function_declaration
            {
                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   use_function_declaration
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
;

use_function_declaration:
        namespace_name
            {
                $$ = &ast.StmtUse{
                    Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                    Use: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                        Parts:         $1.(*ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
    |   namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUse{
                    Position: yylex.(*Parser).builder.NewNodeListTokenPosition($1.(*ParserSeparatedList).Items, $3),
                    Use: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                        Parts:         $1.(*ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ParserSeparatedList).SeparatorTkns,
                    },
                    AsTkn: $2,
                    Alias: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.StmtUse{
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $2.(*ParserSeparatedList).Items),
                    NsSeparatorTkn: $1,
                    Use: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($2.(*ParserSeparatedList).Items),
                        Parts:         $2.(*ParserSeparatedList).Items,
                        SeparatorTkns: $2.(*ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUse{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    NsSeparatorTkn: $1,
                    Use: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($2.(*ParserSeparatedList).Items),
                        Parts:         $2.(*ParserSeparatedList).Items,
                        SeparatorTkns: $2.(*ParserSeparatedList).SeparatorTkns,
                    },
                    AsTkn: $3,
                    Alias: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($4),
                        IdentifierTkn: $4,
                        Value:         $4.Value,
                    },
                }
            }
;

use_const_declarations:
        use_const_declarations ',' use_const_declaration
            {
                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   use_const_declaration
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
;

use_const_declaration:
        namespace_name
            {
                $$ = &ast.StmtUse{
                    Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                    Use: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                        Parts:         $1.(*ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
    |   namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUse{
                    Position: yylex.(*Parser).builder.NewNodeListTokenPosition($1.(*ParserSeparatedList).Items, $3),
                    Use: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                        Parts:         $1.(*ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ParserSeparatedList).SeparatorTkns,
                    },
                    AsTkn: $2,
                    Alias: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.StmtUse{
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $2.(*ParserSeparatedList).Items),
                    NsSeparatorTkn: $1,
                    Use: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($2.(*ParserSeparatedList).Items),
                        Parts:         $2.(*ParserSeparatedList).Items,
                        SeparatorTkns: $2.(*ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUse{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    NsSeparatorTkn: $1,
                    Use: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($2.(*ParserSeparatedList).Items),
                        Parts:         $2.(*ParserSeparatedList).Items,
                        SeparatorTkns: $2.(*ParserSeparatedList).SeparatorTkns,
                    },
                    AsTkn: $3,
                    Alias: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($4),
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
                constList.Position = yylex.(*Parser).builder.NewNodesPosition($1, $5)
                constList.SeparatorTkns = append(constList.SeparatorTkns, $2)
                constList.Consts = append(constList.Consts, &ast.StmtConstant{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($3, $5),
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                    EqualTkn: $4,
                    Expr:     $5,
                })

                $$ = $1
            }
    |   T_CONST T_STRING '=' static_scalar
            {
                $$ = &ast.StmtConstList{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $4),
                    ConstTkn: $1,
                    Consts: []ast.Vertex{
                        &ast.StmtConstant{
                            Position: yylex.(*Parser).builder.NewTokenNodePosition($2, $4),
                            Name: &ast.Identifier{
                                Position: yylex.(*Parser).builder.NewTokenPosition($2),
                                IdentifierTkn: $2,
                                Value:         $2.Value,
                            },
                            EqualTkn: $3,
                            Expr:     $4,
                        },
                    },
                }
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
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
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
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $2),
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
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
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenCurlyBracketTkn:  $1,
                    Stmts:                $2,
                    CloseCurlyBracketTkn: $3,
                }
            }
    |   T_IF parenthesis_expr statement elseif_list else_single
            {
                pos := yylex.(*Parser).builder.NewTokenNodePosition($1, $3)
                if $5 != nil {
                    pos = yylex.(*Parser).builder.NewTokenNodePosition($1, $5)
                } else if len($4) > 0 {
                    pos = yylex.(*Parser).builder.NewTokenNodeListPosition($1, $4)
                }

                $$ = &ast.StmtIf{
                    Position: pos,
                    IfTkn:               $1,
                    OpenParenthesisTkn:  $2.(*ast.ExprBrackets).OpenParenthesisTkn,
                    Cond:                $2.(*ast.ExprBrackets).Expr,
                    CloseParenthesisTkn: $2.(*ast.ExprBrackets).CloseParenthesisTkn,
                    Stmt:                $3,
                    ElseIf:              $4,
                    Else:                $5,
                }
            }
    |   T_IF parenthesis_expr ':' inner_statement_list new_elseif_list new_else_single T_ENDIF ';'
            {
                $$ = &ast.StmtIf{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $8),
                    IfTkn:               $1,
                    OpenParenthesisTkn:  $2.(*ast.ExprBrackets).OpenParenthesisTkn,
                    Cond:                $2.(*ast.ExprBrackets).Expr,
                    CloseParenthesisTkn: $2.(*ast.ExprBrackets).CloseParenthesisTkn,
                    ColonTkn:            $3,
                    Stmt: &ast.StmtStmtList{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($4),
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
                $3.(*ast.StmtWhile).OpenParenthesisTkn = $2.(*ast.ExprBrackets).OpenParenthesisTkn
                $3.(*ast.StmtWhile).Cond = $2.(*ast.ExprBrackets).Expr
                $3.(*ast.StmtWhile).CloseParenthesisTkn = $2.(*ast.ExprBrackets).CloseParenthesisTkn
                $3.(*ast.StmtWhile).Position = yylex.(*Parser).builder.NewTokenNodePosition($1, $3)

                $$ = $3
            }
    |   T_DO statement T_WHILE parenthesis_expr ';'
            {
                $$ = &ast.StmtDo{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $5),
                    DoTkn:               $1,
                    Stmt:                $2,
                    WhileTkn:            $3,
                    OpenParenthesisTkn:  $4.(*ast.ExprBrackets).OpenParenthesisTkn,
                    Cond:                $4.(*ast.ExprBrackets).Expr,
                    CloseParenthesisTkn: $4.(*ast.ExprBrackets).CloseParenthesisTkn,
                    SemiColonTkn:        $5,
                }
            }
    |   T_FOR '(' for_expr ';' for_expr ';' for_expr ')' for_statement
            {
                $9.(*ast.StmtFor).ForTkn = $1
                $9.(*ast.StmtFor).OpenParenthesisTkn = $2
                $9.(*ast.StmtFor).Init = $3.(*ParserSeparatedList).Items
                $9.(*ast.StmtFor).InitSeparatorTkns = $3.(*ParserSeparatedList).SeparatorTkns
                $9.(*ast.StmtFor).InitSemiColonTkn = $4
                $9.(*ast.StmtFor).Cond = $5.(*ParserSeparatedList).Items
                $9.(*ast.StmtFor).CondSeparatorTkns = $5.(*ParserSeparatedList).SeparatorTkns
                $9.(*ast.StmtFor).CondSemiColonTkn = $6
                $9.(*ast.StmtFor).Loop = $7.(*ParserSeparatedList).Items
                $9.(*ast.StmtFor).LoopSeparatorTkns = $7.(*ParserSeparatedList).SeparatorTkns
                $9.(*ast.StmtFor).CloseParenthesisTkn = $8
                $9.(*ast.StmtFor).Position = yylex.(*Parser).builder.NewTokenNodePosition($1, $9)

                $$ = $9
            }
    |   T_SWITCH parenthesis_expr switch_case_list
            {
                $3.(*ast.StmtSwitch).SwitchTkn = $1
                $3.(*ast.StmtSwitch).OpenParenthesisTkn = $2.(*ast.ExprBrackets).OpenParenthesisTkn
                $3.(*ast.StmtSwitch).Cond = $2.(*ast.ExprBrackets).Expr
                $3.(*ast.StmtSwitch).CloseParenthesisTkn = $2.(*ast.ExprBrackets).CloseParenthesisTkn
                $3.(*ast.StmtSwitch).Position = yylex.(*Parser).builder.NewTokenNodePosition($1, $3)

                $$ = $3
            }
    |   T_BREAK ';'
            {
                $$ = &ast.StmtBreak{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $2),
                    BreakTkn:     $1,
                    SemiColonTkn: $2,
                }
            }
    |   T_BREAK expr ';'
            {
                $$ = &ast.StmtBreak{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    BreakTkn:     $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   T_CONTINUE ';'
            {
                $$ = &ast.StmtContinue{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $2),
                    ContinueTkn:  $1,
                    SemiColonTkn: $2,
                }
            }
    |   T_CONTINUE expr ';'
            {
                $$ = &ast.StmtContinue{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    ContinueTkn:  $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   T_RETURN ';'
            {
                $$ = &ast.StmtReturn{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $2),
                    ReturnTkn:    $1,
                    SemiColonTkn: $2,
                }
            }
    |   T_RETURN expr_without_variable ';'
            {
                $$ = &ast.StmtReturn{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    ReturnTkn:    $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   T_RETURN variable ';'
            {
                $$ = &ast.StmtReturn{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    ReturnTkn:    $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   yield_expr ';'
            {
                $$ = &ast.StmtExpression{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $2),
                    Expr:         $1,
                    SemiColonTkn: $2,
                }
            }
    |   T_GLOBAL global_var_list ';'
            {
                $2.(*ast.StmtGlobal).GlobalTkn = $1
                $2.(*ast.StmtGlobal).SemiColonTkn = $3
                $2.(*ast.StmtGlobal).Position = yylex.(*Parser).builder.NewTokensPosition($1, $3)

                $$ = $2
            }
    |   T_STATIC static_var_list ';'
            {
                $2.(*ast.StmtStatic).StaticTkn = $1
                $2.(*ast.StmtStatic).SemiColonTkn = $3
                $2.(*ast.StmtStatic).Position = yylex.(*Parser).builder.NewTokensPosition($1, $3)

                $$ = $2
            }
    |   T_ECHO echo_expr_list ';'
            {
                $2.(*ast.StmtEcho).EchoTkn = $1
                $2.(*ast.StmtEcho).SemiColonTkn = $3
                $2.(*ast.StmtEcho).Position = yylex.(*Parser).builder.NewTokensPosition($1, $3)

                $$ = $2
            }
    |   T_INLINE_HTML
            {
                $$ = &ast.StmtInlineHtml{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    InlineHtmlTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   expr ';'
            {
                $$ = &ast.StmtExpression{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $2),
                    Expr:         $1,
                    SemiColonTkn: $2,
                }
            }
    |   T_UNSET '(' unset_variables ')' ';'
            {
                $3.(*ast.StmtUnset).UnsetTkn = $1
                $3.(*ast.StmtUnset).OpenParenthesisTkn = $2
                $3.(*ast.StmtUnset).CloseParenthesisTkn = $4
                $3.(*ast.StmtUnset).SemiColonTkn = $5
                $3.(*ast.StmtUnset).Position = yylex.(*Parser).builder.NewTokensPosition($1, $5)

                $$ = $3
            }
    |   T_FOREACH '(' variable T_AS foreach_variable foreach_optional_arg ')' foreach_statement
            {
                foreach := $8.(*ast.StmtForeach)

                foreach.Position = yylex.(*Parser).builder.NewTokenNodePosition($1, $8)
                foreach.ForeachTkn = $1
                foreach.OpenParenthesisTkn = $2
                foreach.Expr = $3
                foreach.AsTkn = $4
                foreach.Var = $5
                foreach.CloseParenthesisTkn = $7

                if $6 != nil {
                    foreach.Key = foreach.Var
                    foreach.DoubleArrowTkn = $6.(*ast.StmtForeach).DoubleArrowTkn
                    foreach.Var = $6.(*ast.StmtForeach).Var
                }

                if val, ok := foreach.Key.(*ast.StmtForeach); ok {
                    yylex.(*Parser).errHandlerFunc(errors.NewError("Key element cannot be a reference", val.AmpersandTkn.Position))
                    foreach.Key = val.Var
                }

                if val, ok := foreach.Var.(*ast.StmtForeach); ok {
                    foreach.AmpersandTkn    = val.AmpersandTkn
                    foreach.Var             = val.Var
                }

                $$ = foreach
            }
    |   T_FOREACH '(' expr_without_variable T_AS foreach_variable foreach_optional_arg ')' foreach_statement
            {
                foreach := $8.(*ast.StmtForeach)

                foreach.Position = yylex.(*Parser).builder.NewTokenNodePosition($1, $8)
                foreach.ForeachTkn = $1
                foreach.OpenParenthesisTkn = $2
                foreach.Expr = $3
                foreach.AsTkn = $4
                foreach.Var = $5
                foreach.CloseParenthesisTkn = $7

                if $6 != nil {
                    foreach.Key = foreach.Var
                    foreach.DoubleArrowTkn = $6.(*ast.StmtForeach).DoubleArrowTkn
                    foreach.Var = $6.(*ast.StmtForeach).Var
                }

                if val, ok := foreach.Key.(*ast.StmtForeach); ok {
                    yylex.(*Parser).errHandlerFunc(errors.NewError("Key element cannot be a reference", val.AmpersandTkn.Position))
                    foreach.Key = val.Var
                }

                if val, ok := foreach.Var.(*ast.StmtForeach); ok {
                    foreach.AmpersandTkn    = val.AmpersandTkn
                    foreach.Var             = val.Var
                }

                $$ = foreach
            }
    |   T_DECLARE '(' declare_list ')' declare_statement
            {
                $5.(*ast.StmtDeclare).DeclareTkn = $1
                $5.(*ast.StmtDeclare).OpenParenthesisTkn = $2
                $5.(*ast.StmtDeclare).Consts = $3.(*ParserSeparatedList).Items
                $5.(*ast.StmtDeclare).SeparatorTkns = $3.(*ParserSeparatedList).SeparatorTkns
                $5.(*ast.StmtDeclare).CloseParenthesisTkn = $4
                $5.(*ast.StmtDeclare).Position = yylex.(*Parser).builder.NewTokenNodePosition($1, $5)

                $$ = $5
            }
    |   ';'
            {
                $$ = &ast.StmtNop{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    SemiColonTkn: $1,
                }
            }
    |   T_TRY '{' inner_statement_list '}' catch_statement finally_statement
            {
                pos := yylex.(*Parser).builder.NewTokenNodeListPosition($1, $5)
                if $6 != nil {
                    pos = yylex.(*Parser).builder.NewTokenNodePosition($1, $6)
                }

                $$ = &ast.StmtTry{
                    Position: pos,
                    TryTkn:               $1,
                    OpenCurlyBracketTkn:  $2,
                    Stmts:                $3,
                    CloseCurlyBracketTkn: $4,
                    Catches:              $5,
                    Finally:              $6,
                }
            }
    |   T_THROW expr ';'
            {
                $$ = &ast.StmtThrow{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    ThrowTkn:     $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   T_GOTO T_STRING ';'
            {
                $$ = &ast.StmtGoto{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    GotoTkn: $1,
                    Label: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        IdentifierTkn: $2,
                        Value:         $2.Value,
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
                catch := &ast.StmtCatch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $8),
                    CatchTkn:           $1,
                    OpenParenthesisTkn: $2,
                    Types:              []ast.Vertex{$3},
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($4),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($4),
                            IdentifierTkn: $4,
                            Value:         $4.Value,
                        },
                    },
                    CloseParenthesisTkn:  $5,
                    OpenCurlyBracketTkn:  $6,
                    Stmts:                $7,
                    CloseCurlyBracketTkn: $8,
                }
                $$ = append([]ast.Vertex{catch}, $9...)
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
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
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
                $$ = &ast.StmtCatch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $8),
                    CatchTkn:           $1,
                    OpenParenthesisTkn: $2,
                    Types:              []ast.Vertex{$3},
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($4),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($4),
                            IdentifierTkn: $4,
                            Value:         $4.Value,
                        },
                    },
                    CloseParenthesisTkn:  $5,
                    OpenCurlyBracketTkn:  $6,
                    Stmts:                $7,
                    CloseCurlyBracketTkn: $8,
                }
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
                $$ = &ast.StmtFunction{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $9),
                    FunctionTkn:  $1,
                    AmpersandTkn: $2,
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                    OpenParenthesisTkn:   $4,
                    Params:               $5.(*ParserSeparatedList).Items,
                    SeparatorTkns:        $5.(*ParserSeparatedList).SeparatorTkns,
                    CloseParenthesisTkn:  $6,
                    OpenCurlyBracketTkn:  $7,
                    Stmts:                $8,
                    CloseCurlyBracketTkn: $9,
                }
            }
;

unticked_class_declaration_statement:
        class_entry_type T_STRING extends_from implements_list '{' class_statement_list '}'
            {
                switch n := $1.(type) {
                    case *ast.StmtClass :
                        className := &ast.Identifier{
                            Position:      yylex.(*Parser).builder.NewTokenPosition($2),
                            IdentifierTkn: $2,
                            Value:         $2.Value,
                        }

                        n.Position             = yylex.(*Parser).builder.NewNodeTokenPosition($1, $7)
                        n.Name                 = className
                        n.OpenCurlyBracketTkn  = $5
                        n.Stmts                = $6
                        n.CloseCurlyBracketTkn = $7

                        if $3 != nil {
                            n.ExtendsTkn = $3.(*ast.StmtClass).ExtendsTkn
                            n.Extends    = $3.(*ast.StmtClass).Extends
                        }

                        if $4 != nil {
                            n.ImplementsTkn           = $4.(*ast.StmtClass).ImplementsTkn
                            n.Implements              = $4.(*ast.StmtClass).Implements
                            n.ImplementsSeparatorTkns = $4.(*ast.StmtClass).ImplementsSeparatorTkns
                        }
                    case *ast.StmtTrait :
                        traitName := &ast.Identifier{
                            Position:      yylex.(*Parser).builder.NewTokenPosition($2),
                            IdentifierTkn: $2,
                            Value:         $2.Value,
                        }

                        n.Position             = yylex.(*Parser).builder.NewNodeTokenPosition($1, $7)
                        n.Name                 = traitName
                        n.OpenCurlyBracketTkn  = $5
                        n.Stmts                = $6
                        n.CloseCurlyBracketTkn = $7

                        if $3 != nil {
                            yylex.(*Parser).errHandlerFunc(errors.NewError("A trait cannot extend a class. Traits can only be composed from other traits with the 'use' keyword", $3.(*ast.StmtClass).Position))
                        }

                        if $4 != nil {
                            yylex.(*Parser).errHandlerFunc(errors.NewError("A trait cannot implement an interface", $4.(*ast.StmtClass).Position))
                        }
                }

                $$ = $1
            }
    |   interface_entry T_STRING interface_extends_list '{' class_statement_list '}'
            {
                iface := &ast.StmtInterface{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $6),
                    InterfaceTkn: $1,
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($2),
                        IdentifierTkn: $2,
                        Value:         $2.Value,
                    },
                    OpenCurlyBracketTkn:  $4,
                    Stmts:                $5,
                    CloseCurlyBracketTkn: $6,
                }

                if $3 != nil {
                    iface.ExtendsTkn           = $3.(*ast.StmtInterface).ExtendsTkn
                    iface.Extends              = $3.(*ast.StmtInterface).Extends
                    iface.ExtendsSeparatorTkns = $3.(*ast.StmtInterface).ExtendsSeparatorTkns
                }

                $$ = iface
            }
;


class_entry_type:
        T_CLASS
            {
                $$ = &ast.StmtClass{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    ClassTkn: $1,
                }
            }
    |   T_ABSTRACT T_CLASS
            {
                $$ = &ast.StmtClass{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $2),
                    Modifiers: []ast.Vertex{
                        &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($1),
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
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    TraitTkn: $1,
                }
            }
    |   T_FINAL T_CLASS
            {
                $$ = &ast.StmtClass{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $2),
                    Modifiers: []ast.Vertex{
                        &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($1),
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
                $$ = &ast.StmtClass{
                    Position:   yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    ExtendsTkn: $1,
                    Extends:    $2,
                }
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
                $$ = &ast.StmtInterface{
                    Position:             yylex.(*Parser).builder.NewTokenNodeListPosition($1, $2.(*ParserSeparatedList).Items),
                    ExtendsTkn:           $1,
                    Extends:              $2.(*ParserSeparatedList).Items,
                    ExtendsSeparatorTkns: $2.(*ParserSeparatedList).SeparatorTkns,
                };
            }
;

implements_list:
        /* empty */
            {
                $$ = nil
            }
    |   T_IMPLEMENTS interface_list
            {
                $$ = &ast.StmtClass{
                    Position:                yylex.(*Parser).builder.NewTokenNodeListPosition($1, $2.(*ParserSeparatedList).Items),
                    ImplementsTkn:           $1,
                    Implements:              $2.(*ParserSeparatedList).Items,
                    ImplementsSeparatorTkns: $2.(*ParserSeparatedList).SeparatorTkns,
                };
            }
;

interface_list:
        fully_qualified_class_name
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
    |   interface_list ',' fully_qualified_class_name
            {
                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, $3)

                $$ = $1
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
                $$ = &ast.StmtForeach{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    AmpersandTkn: $1,
                    Var:          $2,
                }
            }
    |   T_LIST '(' assignment_list ')'
            {
                pairList := $3.(*ParserSeparatedList)
                fistPair := pairList.Items[0].(*ast.ExprArrayItem)

                if fistPair.Key == nil && fistPair.Val == nil && len(pairList.Items) == 1 {
                    pairList.Items = nil
                }

                $$ = &ast.ExprList{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    ListTkn:         $1,
                    OpenBracketTkn:  $2,
                    Items:           $3.(*ParserSeparatedList).Items,
                    SeparatorTkns:   $3.(*ParserSeparatedList).SeparatorTkns,
                    CloseBracketTkn: $4,
                }
            }
;

for_statement:
        statement
            {
                $$ = &ast.StmtFor{
                    Position: yylex.(*Parser).builder.NewNodePosition($1),
                    Stmt: $1,
                }
            }
    |   ':' inner_statement_list T_ENDFOR ';'
            {
                $$ = &ast.StmtFor{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    ColonTkn: $1,
                    Stmt: &ast.StmtStmtList{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($2),
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
                    Position: yylex.(*Parser).builder.NewNodePosition($1),
                    Stmt: $1,
                }
            }
    |   ':' inner_statement_list T_ENDFOREACH ';'
            {
                $$ = &ast.StmtForeach{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    ColonTkn: $1,
                    Stmt: &ast.StmtStmtList{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($2),
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
                    Position: yylex.(*Parser).builder.NewNodePosition($1),
                    Stmt: $1,
                }
            }
    |   ':' inner_statement_list T_ENDDECLARE ';'
            {
                $$ = &ast.StmtDeclare{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    ColonTkn: $1,
                    Stmt: &ast.StmtStmtList{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($2),
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
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{
                        &ast.StmtConstant{
                            Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $3),
                            Name: &ast.Identifier{
                                Position: yylex.(*Parser).builder.NewTokenPosition($1),
                                IdentifierTkn: $1,
                                Value:         $1.Value,
                            },
                            EqualTkn: $2,
                            Expr:     $3,
                        },
                    },
                }
            }
    |   declare_list ',' T_STRING '=' static_scalar
            {
                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append(
                    $1.(*ParserSeparatedList).Items,
                    &ast.StmtConstant{
                        Position: yylex.(*Parser).builder.NewTokenNodePosition($3, $5),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($3),
                            IdentifierTkn: $3,
                            Value:         $3.Value,
                        },
                        EqualTkn: $4,
                        Expr:     $5,
                    },
                )

                $$ = $1
            }
;


switch_case_list:
        '{' case_list '}'
            {
                $$ = &ast.StmtSwitch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenCurlyBracketTkn:  $1,
                    Cases:                $2,
                    CloseCurlyBracketTkn: $3,
                }
            }
    |   '{' ';' case_list '}'
            {
                $$ = &ast.StmtSwitch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    OpenCurlyBracketTkn:  $1,
                    CaseSeparatorTkn:     $2,
                    Cases:                $3,
                    CloseCurlyBracketTkn: $4,
                }
            }
    |   ':' case_list T_ENDSWITCH ';'
            {
                $$ = &ast.StmtSwitch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    ColonTkn:     $1,
                    Cases:        $2,
                    EndSwitchTkn: $3,
                    SemiColonTkn: $4,
                }
            }
    |   ':' ';' case_list T_ENDSWITCH ';'
            {
                $$ = &ast.StmtSwitch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $5),
                    ColonTkn:         $1,
                    CaseSeparatorTkn: $2,
                    Cases:            $3,
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
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($2, $5),
                    CaseTkn:          $2,
                    Cond:             $3,
                    CaseSeparatorTkn: $4,
                    Stmts:            $5,
                })
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                $$ = append($1, &ast.StmtDefault{
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($2, $4),
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
                    Position: yylex.(*Parser).builder.NewNodePosition($1),
                    Stmt: $1,
                }
            }
    |   ':' inner_statement_list T_ENDWHILE ';'
            {
                $$ = &ast.StmtWhile{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    ColonTkn: $1,
                    Stmt: &ast.StmtStmtList{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($2),
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
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($2, $4),
                    ElseIfTkn:           $2,
                    OpenParenthesisTkn:  $3.(*ast.ExprBrackets).OpenParenthesisTkn,
                    Cond:                $3.(*ast.ExprBrackets).Expr,
                    CloseParenthesisTkn: $3.(*ast.ExprBrackets).CloseParenthesisTkn,
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
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($2, $5),
                    ElseIfTkn:           $2,
                    OpenParenthesisTkn:  $3.(*ast.ExprBrackets).OpenParenthesisTkn,
                    Cond:                $3.(*ast.ExprBrackets).Expr,
                    CloseParenthesisTkn: $3.(*ast.ExprBrackets).CloseParenthesisTkn,
                    ColonTkn:            $4,
                    Stmt: &ast.StmtStmtList{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($5),
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
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
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
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $3),
                    ElseTkn:  $1,
                    ColonTkn: $2,
                    Stmt: &ast.StmtStmtList{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($3),
                        Stmts: $3,
                    },
                }
            }
;


parameter_list:
        non_empty_parameter_list
            {
                $$ = $1
            }
    |   /* empty */
            {
                $$ = &ParserSeparatedList{}
            }
;

non_empty_parameter_list:
        parameter
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
    |   non_empty_parameter_list ',' parameter
            {
                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, $3)

                $$ = $1
            }
;

parameter:
        optional_class_type is_reference is_variadic T_VARIABLE
            {
                pos := yylex.(*Parser).builder.NewTokenPosition($4)
                if $1 != nil {
                    pos = yylex.(*Parser).builder.NewNodeTokenPosition($1, $4)
                } else if $2 != nil {
                    pos = yylex.(*Parser).builder.NewTokensPosition($2, $4)
                } else if $3 != nil {
                    pos = yylex.(*Parser).builder.NewTokensPosition($3, $4)
                }

                $$ = &ast.Parameter{
                    Position: pos,
                    Type:         $1,
                    AmpersandTkn: $2,
                    VariadicTkn:  $3,
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($4),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($4),
                            IdentifierTkn: $4,
                            Value:         $4.Value,
                        },
                    },
                }
            }
    |   optional_class_type is_reference is_variadic T_VARIABLE '=' expr
            {
                pos := yylex.(*Parser).builder.NewTokenNodePosition($4, $6)
                if $1 != nil {
                    pos = yylex.(*Parser).builder.NewNodesPosition($1, $6)
                } else if $2 != nil {
                    pos = yylex.(*Parser).builder.NewTokenNodePosition($2, $6)
                } else if $3 != nil {
                    pos = yylex.(*Parser).builder.NewTokenNodePosition($3, $6)
                }

                $$ = &ast.Parameter{
                    Position: pos,
                    Type:         $1,
                    AmpersandTkn: $2,
                    VariadicTkn:  $3,
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($4),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($4),
                            IdentifierTkn: $4,
                            Value:         $4.Value,
                        },
                    },
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
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_CALLABLE
            {
                $$ = &ast.Identifier{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
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
                $$ = &ArgumentList{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $2),
                    OpenParenthesisTkn: $1,
                    CloseParenthesisTkn: $2,
                }
            }
    |   '(' non_empty_function_call_parameter_list ')'
            {
                argumentList := $2.(*ArgumentList)
                argumentList.Position = yylex.(*Parser).builder.NewTokensPosition($1, $3)
                argumentList.OpenParenthesisTkn = $1
                argumentList.CloseParenthesisTkn = $3

                $$ = argumentList
            }
    |   '(' yield_expr ')'
            {
                $$ = &ArgumentList{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenParenthesisTkn: $1,
                    Arguments: []ast.Vertex{
                        &ast.Argument{
                            Position: yylex.(*Parser).builder.NewNodePosition($2),
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
                $$ = &ArgumentList{
                    Arguments: []ast.Vertex{$1},
                }
            }
    |   non_empty_function_call_parameter_list ',' function_call_parameter
            {
                $1.(*ArgumentList).SeparatorTkns = append($1.(*ArgumentList).SeparatorTkns, $2)
                $1.(*ArgumentList).Arguments = append($1.(*ArgumentList).Arguments, $3)

                $$ = $1
            }
;

function_call_parameter:
        expr_without_variable
            {
                $$ = &ast.Argument{
                    Position: yylex.(*Parser).builder.NewNodePosition($1),
                    Expr: $1,
                }
            }
    |   variable
            {
                $$ = &ast.Argument{
                    Position: yylex.(*Parser).builder.NewNodePosition($1),
                    Expr: $1,
                }
            }
    |   '&' w_variable
            {
                $$ = &ast.Argument{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    AmpersandTkn: $1,
                    Expr:         $2,
                }
            }
    |   T_ELLIPSIS expr
            {
                $$ = &ast.Argument{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
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
                $$ = &ast.ExprVariable{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
            }
    |   '$' r_variable
            {
                $$ = &ast.ExprVariable{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    DollarTkn: $1,
                    Name:      $2,
                }
            }
    |   '$' '{' expr '}'
            {
                $$ = &ast.ExprVariable{
                    Position:             yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    DollarTkn:            $1,
                    OpenCurlyBracketTkn:  $2,
                    Name:                 $3,
                    CloseCurlyBracketTkn: $4,
                }
            }
;


static_var_list:
        static_var_list ',' T_VARIABLE
            {
                $1.(*ast.StmtStatic).Vars = append($1.(*ast.StmtStatic).Vars, &ast.StmtStaticVar{
                    Position: yylex.(*Parser).builder.NewTokenPosition($3),
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($3),
                            IdentifierTkn: $3,
                            Value:         $3.Value,
                        },
                    },
                })
                $1.(*ast.StmtStatic).SeparatorTkns = append($1.(*ast.StmtStatic).SeparatorTkns, $2)

                $$ = $1
            }
    |   static_var_list ',' T_VARIABLE '=' static_scalar
            {
                $1.(*ast.StmtStatic).Vars = append($1.(*ast.StmtStatic).Vars, &ast.StmtStaticVar{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($3, $5),
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($3),
                            IdentifierTkn: $3,
                            Value:         $3.Value,
                        },
                    },
                    EqualTkn: $4,
                    Expr:     $5,
                })
                $1.(*ast.StmtStatic).SeparatorTkns = append($1.(*ast.StmtStatic).SeparatorTkns, $2)

                $$ = $1
            }
    |   T_VARIABLE
            {
                $$ = &ast.StmtStatic{
                    Vars: []ast.Vertex{
                        &ast.StmtStaticVar{
                            Position: yylex.(*Parser).builder.NewTokenPosition($1),
                            Var: &ast.ExprVariable{
                                Position: yylex.(*Parser).builder.NewTokenPosition($1),
                                Name: &ast.Identifier{
                                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                                    IdentifierTkn: $1,
                                    Value:         $1.Value,
                                },
                            },
                        },
                    },
                }
            }
    |   T_VARIABLE '=' static_scalar
            {
                $$ = &ast.StmtStatic{
                    Vars: []ast.Vertex{
                        &ast.StmtStaticVar{
                            Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $3),
                            Var: &ast.ExprVariable{
                                Position: yylex.(*Parser).builder.NewTokenPosition($1),
                                Name: &ast.Identifier{
                                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                                    IdentifierTkn: $1,
                                    Value:         $1.Value,
                                },
                            },
                            EqualTkn: $2,
                            Expr:     $3,
                        },
                    },
                }
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
                $$ = &ast.StmtPropertyList{
                    Position: yylex.(*Parser).builder.NewNodeListTokenPosition($1, $3),
                    Modifiers:     $1,
                    Props:         $2.(*ParserSeparatedList).Items,
                    SeparatorTkns: $2.(*ParserSeparatedList).SeparatorTkns,
                    SemiColonTkn:  $3,
                }
            }
    |   class_constant_declaration ';'
            {
                $1.(*ast.StmtClassConstList).SemiColonTkn = $2
                $1.(*ast.StmtClassConstList).Position = yylex.(*Parser).builder.NewNodeTokenPosition($1, $2)
                $$ = $1
            }
    |   trait_use_statement
            {
                $$ = $1
            }
    |   method_modifiers function is_reference T_STRING '(' parameter_list ')' method_body
            {
                pos := yylex.(*Parser).builder.NewTokenNodePosition($2, $8)
                if $1 != nil {
                    pos = yylex.(*Parser).builder.NewNodeListNodePosition($1, $8)
                }

                $$ = &ast.StmtClassMethod{
                    Position: pos,
                    Modifiers:    $1,
                    FunctionTkn:  $2,
                    AmpersandTkn: $3,
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($4),
                        IdentifierTkn: $4,
                        Value:         $4.Value,
                    },
                    OpenParenthesisTkn:  $5,
                    Params:              $6.(*ParserSeparatedList).Items,
                    SeparatorTkns:       $6.(*ParserSeparatedList).SeparatorTkns,
                    CloseParenthesisTkn: $7,
                    Stmt:                $8,
                }
            }
;

trait_use_statement:
        T_USE trait_list trait_adaptations
            {
                traitUse := &ast.StmtTraitUse{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $3),
                    UseTkn:        $1,
                    Traits:        $2.(*ParserSeparatedList).Items,
                    SeparatorTkns: $2.(*ParserSeparatedList).SeparatorTkns,
                }

                switch n := $3.(type) {
                case *TraitAdaptationList :
                    traitUse.OpenCurlyBracketTkn = n.OpenCurlyBracketTkn
                    traitUse.Adaptations = n.Adaptations
                    traitUse.CloseCurlyBracketTkn = n.CloseCurlyBracketTkn
                case *ast.StmtNop :
                    traitUse.SemiColonTkn = n.SemiColonTkn
                };

                $$ = traitUse
            }
;

trait_list:
        fully_qualified_class_name
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
    |   trait_list ',' fully_qualified_class_name
            {
                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, $3)

                $$ = $1
            }
;

trait_adaptations:
        ';'
            {
                $$ = &ast.StmtNop{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    SemiColonTkn: $1,
                }
            }
    |   '{' trait_adaptation_list '}'
            {
                $$ = &TraitAdaptationList{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenCurlyBracketTkn:  $1,
                    Adaptations:          $2,
                    CloseCurlyBracketTkn: $3,
                }
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
                $1.(*ast.StmtTraitUsePrecedence).SemiColonTkn = $2

                $$ = $1;
            }
    |   trait_alias ';'
            {
                $1.(*ast.StmtTraitUseAlias).SemiColonTkn = $2

                $$ = $1;
            }
;

trait_precedence:
        trait_method_reference_fully_qualified T_INSTEADOF trait_reference_list
            {
                $$ = &ast.StmtTraitUsePrecedence{
                    Position: yylex.(*Parser).builder.NewNodeNodeListPosition($1, $3.(*ParserSeparatedList).Items),
                    Trait:          $1.(*TraitMethodRef).Trait,
                    DoubleColonTkn: $1.(*TraitMethodRef).DoubleColonTkn,
                    Method:         $1.(*TraitMethodRef).Method,
                    InsteadofTkn:   $2,
                    Insteadof:      $3.(*ParserSeparatedList).Items,
                    SeparatorTkns:  $3.(*ParserSeparatedList).SeparatorTkns,
                }
            }
;

trait_reference_list:
        fully_qualified_class_name
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
    |   trait_reference_list ',' fully_qualified_class_name
            {
                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, $3)

                $$ = $1
            }
;

trait_method_reference:
        T_STRING
            {
                $$ = &TraitMethodRef{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    Method: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
            }
    |   trait_method_reference_fully_qualified
            {
                $$ = $1
            }
;

trait_method_reference_fully_qualified:
        fully_qualified_class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                $$ = &TraitMethodRef{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $3),
                    Trait:          $1,
                    DoubleColonTkn: $2,
                    Method: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
;

trait_alias:
        trait_method_reference T_AS trait_modifiers T_STRING
            {
                $$ = &ast.StmtTraitUseAlias{
                    Position:       yylex.(*Parser).builder.NewNodeTokenPosition($1, $4),
                    Trait:          $1.(*TraitMethodRef).Trait,
                    DoubleColonTkn: $1.(*TraitMethodRef).DoubleColonTkn,
                    Method:         $1.(*TraitMethodRef).Method,
                    AsTkn:          $2,
                    Modifier:       $3,
                    Alias: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($4),
                        IdentifierTkn: $4,
                        Value:         $4.Value,
                    },
                }
            }
    |   trait_method_reference T_AS member_modifier
            {
                $$ = &ast.StmtTraitUseAlias{
                    Position:       yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Trait:          $1.(*TraitMethodRef).Trait,
                    DoubleColonTkn: $1.(*TraitMethodRef).DoubleColonTkn,
                    Method:         $1.(*TraitMethodRef).Method,
                    AsTkn:          $2,
                    Modifier:       $3,
                }
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
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    SemiColonTkn: $1,
                }
            }
    |   '{' inner_statement_list '}'
            {
                $$ = &ast.StmtStmtList{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenCurlyBracketTkn:  $1,
                    Stmts:                $2,
                    CloseCurlyBracketTkn: $3,
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
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
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
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_PROTECTED
            {
                $$ = &ast.Identifier{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_PRIVATE
            {
                $$ = &ast.Identifier{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_STATIC
            {
                $$ = &ast.Identifier{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_ABSTRACT
            {
                $$ = &ast.Identifier{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_FINAL
            {
                $$ = &ast.Identifier{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
;

class_variable_declaration:
        class_variable_declaration ',' T_VARIABLE
            {
                item := &ast.StmtProperty{
                    Position: yylex.(*Parser).builder.NewTokenPosition($3),
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($3),
                            IdentifierTkn: $3,
                            Value:         $3.Value,
                        },
                    },
                }

                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, item)

                $$ = $1
            }
    |   class_variable_declaration ',' T_VARIABLE '=' static_scalar
            {
                item := &ast.StmtProperty{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($3, $5),
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($3),
                            IdentifierTkn: $3,
                            Value:         $3.Value,
                        },
                    },
                    EqualTkn: $4,
                    Expr:     $5,
                }

                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, item)

                $$ = $1
            }
    |   T_VARIABLE
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{
                        &ast.StmtProperty{
                            Position: yylex.(*Parser).builder.NewTokenPosition($1),
                            Var: &ast.ExprVariable{
                                Position: yylex.(*Parser).builder.NewTokenPosition($1),
                                Name: &ast.Identifier{
                                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                                    IdentifierTkn: $1,
                                    Value:         $1.Value,
                                },
                            },
                            Expr: nil,
                        },
                    },
                }
            }
    |   T_VARIABLE '=' static_scalar
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{
                        &ast.StmtProperty{
                            Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $3),
                            Var: &ast.ExprVariable{
                                Position: yylex.(*Parser).builder.NewTokenPosition($1),
                                Name: &ast.Identifier{
                                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                                    IdentifierTkn: $1,
                                    Value:         $1.Value,
                                },
                            },
                            EqualTkn: $2,
                            Expr:     $3,
                        },
                    },
                }
            }
;

class_constant_declaration:
        class_constant_declaration ',' T_STRING '=' static_scalar
            {
                constList := $1.(*ast.StmtClassConstList)
                constList.Position = yylex.(*Parser).builder.NewNodesPosition($1, $5)
                constList.SeparatorTkns = append(constList.SeparatorTkns, $2)
                constList.Consts = append(constList.Consts, &ast.StmtConstant{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($3, $5),
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                    EqualTkn: $4,
                    Expr:     $5,
                })

                $$ = $1
            }
    |   T_CONST T_STRING '=' static_scalar
            {
                $$ = &ast.StmtClassConstList{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $4),
                    ConstTkn: $1,
                    Consts: []ast.Vertex{
                        &ast.StmtConstant{
                            Position: yylex.(*Parser).builder.NewTokenNodePosition($2, $4),
                            Name: &ast.Identifier{
                                Position: yylex.(*Parser).builder.NewTokenPosition($2),
                                IdentifierTkn: $2,
                                Value:         $2.Value,
                            },
                            EqualTkn: $3,
                            Expr:     $4,
                        },
                    },
                }
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
                $$ = &ParserSeparatedList{}
            }
    |   non_empty_for_expr
            {
                $$ = $1
            }
;

non_empty_for_expr:
        non_empty_for_expr ',' expr
            {
                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   expr
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
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
                fetch := &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($2, $4),
                    Var:             nil,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }

                $$ = append($1, fetch)
            }
    |   '[' dim_offset ']'
            {
                fetch := &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    Var:             nil,
                    OpenBracketTkn:  $1,
                    Dim:             $2,
                    CloseBracketTkn: $3,
                }

                $$ = []ast.Vertex{fetch}
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
                    $$ = &ast.ExprNew{
                        Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $3),
                        NewTkn:              $1,
                        Class:               $2,
                        OpenParenthesisTkn:  $3.(*ArgumentList).OpenParenthesisTkn,
                        Args:                $3.(*ArgumentList).Arguments,
                        SeparatorTkns:       $3.(*ArgumentList).SeparatorTkns,
                        CloseParenthesisTkn: $3.(*ArgumentList).CloseParenthesisTkn,
                    }
                } else {
                    $$ = &ast.ExprNew{
                        Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                        NewTkn: $1,
                        Class: $2,
                    }
                }
            }
;

expr_without_variable:
        T_LIST '(' assignment_list ')' '=' expr
            {
                $$ = &ast.ExprAssign{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $6),
                    Var: &ast.ExprList{
                        Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                        ListTkn:         $1,
                        OpenBracketTkn:  $2,
                        Items:           $3.(*ParserSeparatedList).Items,
                        SeparatorTkns:   $3.(*ParserSeparatedList).SeparatorTkns,
                        CloseBracketTkn: $4,
                    },
                    EqualTkn: $5,
                    Expr:     $6,
                }
            }
    |   variable '=' expr
            {
                $$ = &ast.ExprAssign{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable '=' '&' variable
            {
                $$ = &ast.ExprAssignReference{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $4),
                    Var:          $1,
                    EqualTkn:     $2,
                    AmpersandTkn: $3,
                    Expr:         $4,
                }
            }
    |   variable '=' '&' T_NEW class_name_reference ctor_arguments
            {
                var _new *ast.ExprNew
                if $6 != nil {
                    _new = &ast.ExprNew{
                        Position: yylex.(*Parser).builder.NewTokenNodePosition($4, $6),
                        NewTkn:              $4,
                        Class:               $5,
                        OpenParenthesisTkn:  $6.(*ArgumentList).OpenParenthesisTkn,
                        Args:                $6.(*ArgumentList).Arguments,
                        SeparatorTkns:       $6.(*ArgumentList).SeparatorTkns,
                        CloseParenthesisTkn: $6.(*ArgumentList).CloseParenthesisTkn,
                    }
                } else {
                    _new = &ast.ExprNew{
                        Position: yylex.(*Parser).builder.NewTokenNodePosition($4, $5),
                        NewTkn: $4,
                        Class:  $5,
                    }
                }

                $$ = &ast.ExprAssignReference{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, _new),
                    Var:          $1,
                    EqualTkn:     $2,
                    AmpersandTkn: $3,
                    Expr:         _new,
                }
            }
    |   T_CLONE expr
            {
                $$ = &ast.ExprClone{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    CloneTkn: $1,
                    Expr:     $2,
                }
            }
    |   variable T_PLUS_EQUAL expr
            {
                $$ = &ast.ExprAssignPlus{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_MINUS_EQUAL expr
            {
                $$ = &ast.ExprAssignMinus{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_MUL_EQUAL expr
            {
                $$ = &ast.ExprAssignMul{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_POW_EQUAL expr
            {
                $$ = &ast.ExprAssignPow{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_DIV_EQUAL expr
            {
                $$ = &ast.ExprAssignDiv{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_CONCAT_EQUAL expr
            {
                $$ = &ast.ExprAssignConcat{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_MOD_EQUAL expr
            {
                $$ = &ast.ExprAssignMod{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_AND_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseAnd{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_OR_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseOr{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_XOR_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseXor{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_SL_EQUAL expr
            {
                $$ = &ast.ExprAssignShiftLeft{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_SR_EQUAL expr
            {
                $$ = &ast.ExprAssignShiftRight{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   rw_variable T_INC
            {
                $$ = &ast.ExprPostInc{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $2),
                    Var:    $1,
                    IncTkn: $2,
                }
            }
    |   T_INC rw_variable
            {
                $$ = &ast.ExprPreInc{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    IncTkn: $1,
                    Var:    $2,
                }
            }
    |   rw_variable T_DEC
            {
                $$ = &ast.ExprPostDec{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $2),
                    Var:    $1,
                    DecTkn: $2,
                }
            }
    |   T_DEC rw_variable
            {
                $$ = &ast.ExprPreDec{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    DecTkn: $1,
                    Var:    $2,
                }
            }
    |   expr T_BOOLEAN_OR expr
            {
                $$ = &ast.ExprBinaryBooleanOr{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_BOOLEAN_AND expr
            {
                $$ = &ast.ExprBinaryBooleanAnd{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_LOGICAL_OR expr
            {
                $$ = &ast.ExprBinaryLogicalOr{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_LOGICAL_AND expr
            {
                $$ = &ast.ExprBinaryLogicalAnd{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_LOGICAL_XOR expr
            {
                $$ = &ast.ExprBinaryLogicalXor{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '|' expr
            {
                $$ = &ast.ExprBinaryBitwiseOr{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '&' expr
            {
                $$ = &ast.ExprBinaryBitwiseAnd{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '^' expr
            {
                $$ = &ast.ExprBinaryBitwiseXor{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '.' expr
            {
                $$ = &ast.ExprBinaryConcat{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '+' expr
            {
                $$ = &ast.ExprBinaryPlus{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '-' expr
            {
                $$ = &ast.ExprBinaryMinus{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '*' expr
            {
                $$ = &ast.ExprBinaryMul{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_POW expr
            {
                $$ = &ast.ExprBinaryPow{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '/' expr
            {
                $$ = &ast.ExprBinaryDiv{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '%' expr
            {
                $$ = &ast.ExprBinaryMod{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_SL expr
            {
                $$ = &ast.ExprBinaryShiftLeft{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_SR expr
            {
                $$ = &ast.ExprBinaryShiftRight{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   '+' expr %prec T_INC
            {
                $$ = &ast.ExprUnaryPlus{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    PlusTkn: $1,
                    Expr:    $2,
                }
            }
    |   '-' expr %prec T_INC
            {
                $$ = &ast.ExprUnaryMinus{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    MinusTkn: $1,
                    Expr:     $2,
                }
            }
    |   '!' expr
            {
                $$ = &ast.ExprBooleanNot{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    ExclamationTkn: $1,
                    Expr:           $2,
                }
            }
    |   '~' expr
            {
                $$ = &ast.ExprBitwiseNot{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    TildaTkn: $1,
                    Expr:     $2,
                }
            }
    |   expr T_IS_IDENTICAL expr
            {
                $$ = &ast.ExprBinaryIdentical{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_IS_NOT_IDENTICAL expr
            {
                $$ = &ast.ExprBinaryNotIdentical{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_IS_EQUAL expr
            {
                $$ = &ast.ExprBinaryEqual{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_IS_NOT_EQUAL expr
            {
                $$ = &ast.ExprBinaryNotEqual{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '<' expr
            {
                $$ = &ast.ExprBinarySmaller{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_IS_SMALLER_OR_EQUAL expr
            {
                $$ = &ast.ExprBinarySmallerOrEqual{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '>' expr
            {
                $$ = &ast.ExprBinaryGreater{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_IS_GREATER_OR_EQUAL expr
            {
                $$ = &ast.ExprBinaryGreaterOrEqual{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_INSTANCEOF class_name_reference
            {
                $$ = &ast.ExprInstanceOf{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Expr:          $1,
                    InstanceOfTkn: $2,
                    Class:         $3,
                }
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
                $$ = &ast.ExprBrackets{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenParenthesisTkn:  $1,
                    Expr:                $2,
                    CloseParenthesisTkn: $3,
                }

                for _, n := range($4) {
                    switch nn := n.(type) {
                        case *ast.ExprFunctionCall:
                            nn.Function = $$
                            nn.Position = yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn

                        case *ast.ExprArrayDimFetch:
                            nn.Var = $$
                            nn.Position = yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn

                        case *ast.ExprPropertyFetch:
                            nn.Var = $$
                            nn.Position = yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn

                        case *ast.ExprMethodCall:
                            nn.Var = $$
                            nn.Position = yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn
                    }
                }
            }
    |   expr '?' expr ':' expr
            {
                $$ = &ast.ExprTernary{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $5),
                    Cond:        $1,
                    QuestionTkn: $2,
                    IfTrue:      $3,
                    ColonTkn:    $4,
                    IfFalse:     $5,
                }
            }
    |   expr '?' ':' expr
            {
                $$ = &ast.ExprTernary{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $4),
                    Cond:        $1,
                    QuestionTkn: $2,
                    ColonTkn:    $3,
                    IfFalse:     $4,
                }
            }
    |   internal_functions_in_yacc
            {
                $$ = $1
            }
    |   T_INT_CAST expr
            {
                $$ = &ast.ExprCastInt{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_DOUBLE_CAST expr
            {
                $$ = &ast.ExprCastDouble{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_STRING_CAST expr
            {
                $$ = &ast.ExprCastString{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_ARRAY_CAST expr
            {
                $$ = &ast.ExprCastArray{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_OBJECT_CAST expr
            {
                $$ = &ast.ExprCastObject{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_BOOL_CAST expr
            {
                $$ = &ast.ExprCastBool{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_UNSET_CAST expr
            {
                $$ = &ast.ExprCastUnset{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_EXIT exit_expr
            {
                exit := &ast.ExprExit{
                    ExitTkn: $1,
                }

                if $2 == nil {
                    exit.Position = yylex.(*Parser).builder.NewTokenPosition($1)
                } else {
                    exit.Position       = yylex.(*Parser).builder.NewTokenNodePosition($1, $2)
                    exit.OpenParenthesisTkn  = $2.(*ast.ExprBrackets).OpenParenthesisTkn
                    exit.Expr                = $2.(*ast.ExprBrackets).Expr
                    exit.CloseParenthesisTkn = $2.(*ast.ExprBrackets).CloseParenthesisTkn
                }

                $$ = exit
            }
    |   '@' expr
            {
                $$ = &ast.ExprErrorSuppress{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    AtTkn: $1,
                    Expr:  $2,
                }
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
                $$ = &ast.ExprShellExec{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenBacktickTkn:  $1,
                    Parts:            $2,
                    CloseBacktickTkn: $3,
                }
            }
    |   T_PRINT expr
            {
                $$ = &ast.ExprPrint{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    PrintTkn: $1,
                    Expr:     $2,
                }
            }
    |   T_YIELD
            {
                $$ = &ast.ExprYield{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    YieldTkn: $1,
                }
            }
    |   function is_reference '(' parameter_list ')' lexical_vars '{' inner_statement_list '}'
            {
                closure := $6.(*ast.ExprClosure)

                closure.Position             = yylex.(*Parser).builder.NewTokensPosition($1, $9)
                closure.FunctionTkn          = $1
                closure.AmpersandTkn         = $2
                closure.OpenParenthesisTkn   = $3
                closure.Params               = $4.(*ParserSeparatedList).Items
                closure.SeparatorTkns        = $4.(*ParserSeparatedList).SeparatorTkns
                closure.CloseParenthesisTkn  = $5
                closure.OpenCurlyBracketTkn  = $7
                closure.Stmts                = $8
                closure.CloseCurlyBracketTkn = $9

                $$ = closure
            }
    |   T_STATIC function is_reference '(' parameter_list ')' lexical_vars '{' inner_statement_list '}'
            {
                closure := $7.(*ast.ExprClosure)
                
                closure.Position             = yylex.(*Parser).builder.NewTokensPosition($1, $10)
                closure.StaticTkn            = $1
                closure.FunctionTkn          = $2
                closure.AmpersandTkn         = $3
                closure.OpenParenthesisTkn   = $4
                closure.Params               = $5.(*ParserSeparatedList).Items
                closure.SeparatorTkns        = $5.(*ParserSeparatedList).SeparatorTkns
                closure.CloseParenthesisTkn  = $6
                closure.OpenCurlyBracketTkn  = $8
                closure.Stmts                = $9
                closure.CloseCurlyBracketTkn = $10

                $$ = closure
            }
;

yield_expr:
        T_YIELD expr_without_variable
            {
                $$ = &ast.ExprYield{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    YieldTkn: $1,
                    Val:      $2,
                }
            }
    |   T_YIELD variable
            {
                $$ = &ast.ExprYield{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    YieldTkn: $1,
                    Val:      $2,
                }
            }
    |   T_YIELD expr T_DOUBLE_ARROW expr_without_variable
            {
                $$ = &ast.ExprYield{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $4),
                    YieldTkn:       $1,
                    Key:            $2,
                    DoubleArrowTkn: $3,
                    Val:            $4,
                }
            }
    |   T_YIELD expr T_DOUBLE_ARROW variable
            {
                $$ = &ast.ExprYield{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $4),
                    YieldTkn:       $1,
                    Key:            $2,
                    DoubleArrowTkn: $3,
                    Val:            $4,
                }
            }
;

combined_scalar_offset:
        combined_scalar '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $4),
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   combined_scalar_offset '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $4),
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   T_CONSTANT_ENCAPSED_STRING '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    Var: &ast.ScalarString{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        StringTkn: $1,
                        Value:     $1.Value,
                    },
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   general_constant '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $4),
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
;

combined_scalar:
        T_ARRAY '(' array_pair_list ')'
            {
                $$ = &ast.ExprArray{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    ArrayTkn:        $1,
                    OpenBracketTkn:  $2,
                    Items:           $3.(*ParserSeparatedList).Items,
                    SeparatorTkns:   $3.(*ParserSeparatedList).SeparatorTkns,
                    CloseBracketTkn: $4,
                }
            }
    |   '[' array_pair_list ']'
            {
                $$ = &ast.ExprArray{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenBracketTkn:  $1,
                    Items:           $2.(*ParserSeparatedList).Items,
                    SeparatorTkns:   $2.(*ParserSeparatedList).SeparatorTkns,
                    CloseBracketTkn: $3,
                }
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
                $$ = &ast.ExprClosure{}
            }
    |   T_USE '(' lexical_var_list ')'
            {
                $$ = &ast.ExprClosure{
                    UseTkn:                 $1,
                    UseOpenParenthesisTkn:  $2,
                    Uses:                   $3.(*ParserSeparatedList).Items,
                    UseSeparatorTkns:       $3.(*ParserSeparatedList).SeparatorTkns,
                    UseCloseParenthesisTkn: $4,
                }
            }
;

lexical_var_list:
        lexical_var_list ',' T_VARIABLE
            {
                variable := &ast.ExprClosureUse{
                    Position: yylex.(*Parser).builder.NewTokenPosition($3),
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($3),
                            IdentifierTkn: $3,
                            Value:         $3.Value,
                        },
                    },
                }

                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, variable)

                $$ = $1
            }
    |   lexical_var_list ',' '&' T_VARIABLE
            {
                variable := &ast.ExprClosureUse{
                    Position: yylex.(*Parser).builder.NewTokensPosition($3, $4),
                    AmpersandTkn: $3,
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($4),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($4),
                            IdentifierTkn: $4,
                            Value:         $4.Value,
                        },
                    },
                }

                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, variable)

                $$ = $1
            }
    |   T_VARIABLE
            {
                variable := &ast.ExprClosureUse{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($1),
                            IdentifierTkn: $1,
                            Value:         $1.Value,
                        },
                    },
                }

                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{ variable },
                }
            }
    |   '&' T_VARIABLE
            {
                variable := &ast.ExprClosureUse{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $2),
                    AmpersandTkn: $1,
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($2),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($2),
                            IdentifierTkn: $2,
                            Value:         $2.Value,
                        },
                    },
                }

                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{ variable },
                }
            }
;

function_call:
        namespace_name function_call_parameter_list
            {
                $$ = &ast.ExprFunctionCall{
                    Position: yylex.(*Parser).builder.NewNodeListNodePosition($1.(*ParserSeparatedList).Items, $2),
                    Function: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                        Parts:         $1.(*ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ParserSeparatedList).SeparatorTkns,
                    },
                    OpenParenthesisTkn:  $2.(*ArgumentList).OpenParenthesisTkn,
                    Args:                $2.(*ArgumentList).Arguments,
                    SeparatorTkns:       $2.(*ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $2.(*ArgumentList).CloseParenthesisTkn,
                }
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name function_call_parameter_list
            {
                $$ = &ast.ExprFunctionCall{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $4),
                    Function: &ast.NameRelative{
                        Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $3.(*ParserSeparatedList).Items),
                        NsTkn:          $1,
                        NsSeparatorTkn: $2,
                        Parts:          $3.(*ParserSeparatedList).Items,
                        SeparatorTkns:  $3.(*ParserSeparatedList).SeparatorTkns,
                    },
                    OpenParenthesisTkn:  $4.(*ArgumentList).OpenParenthesisTkn,
                    Args:                $4.(*ArgumentList).Arguments,
                    SeparatorTkns:       $4.(*ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $4.(*ArgumentList).CloseParenthesisTkn,
                }
            }
    |   T_NS_SEPARATOR namespace_name function_call_parameter_list
            {
                $$ = &ast.ExprFunctionCall{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $3),
                    Function: &ast.NameFullyQualified{
                        Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $2.(*ParserSeparatedList).Items),
                        NsSeparatorTkn: $1,
                        Parts:          $2.(*ParserSeparatedList).Items,
                        SeparatorTkns:  $2.(*ParserSeparatedList).SeparatorTkns,
                    },
                    OpenParenthesisTkn:  $3.(*ArgumentList).OpenParenthesisTkn,
                    Args:                $3.(*ArgumentList).Arguments,
                    SeparatorTkns:       $3.(*ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $3.(*ArgumentList).CloseParenthesisTkn,
                }
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM variable_name function_call_parameter_list
            {
                staticCall := &ast.ExprStaticCall{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $4),
                    Class:               $1,
                    DoubleColonTkn:      $2,
                    Call:                $3,
                    OpenParenthesisTkn:  $4.(*ArgumentList).OpenParenthesisTkn,
                    Args:                $4.(*ArgumentList).Arguments,
                    SeparatorTkns:       $4.(*ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $4.(*ArgumentList).CloseParenthesisTkn,
                }

                if brackets, ok := $3.(*ParserBrackets); ok {
                    staticCall.OpenCurlyBracketTkn  = brackets.OpenBracketTkn
                    staticCall.Call                 = brackets.Child
                    staticCall.CloseCurlyBracketTkn = brackets.CloseBracketTkn
                }

                $$ = staticCall
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects function_call_parameter_list
            {
                $$ = &ast.ExprStaticCall{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $4),
                    Class:               $1,
                    DoubleColonTkn:      $2,
                    Call:                $3,
                    OpenParenthesisTkn:  $4.(*ArgumentList).OpenParenthesisTkn,
                    Args:                $4.(*ArgumentList).Arguments,
                    SeparatorTkns:       $4.(*ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $4.(*ArgumentList).CloseParenthesisTkn,
                }
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_name function_call_parameter_list
            {
                staticCall := &ast.ExprStaticCall{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $4),
                    Class:               $1,
                    DoubleColonTkn:      $2,
                    Call:                $3,
                    OpenParenthesisTkn:  $4.(*ArgumentList).OpenParenthesisTkn,
                    Args:                $4.(*ArgumentList).Arguments,
                    SeparatorTkns:       $4.(*ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $4.(*ArgumentList).CloseParenthesisTkn,
                }

                if brackets, ok := $3.(*ParserBrackets); ok {
                    staticCall.OpenCurlyBracketTkn  = brackets.OpenBracketTkn
                    staticCall.Call                 = brackets.Child
                    staticCall.CloseCurlyBracketTkn = brackets.CloseBracketTkn
                }

                $$ = staticCall
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects function_call_parameter_list
            {
                $$ = &ast.ExprStaticCall{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $4),
                    Class:               $1,
                    DoubleColonTkn:      $2,
                    Call:                $3,
                    OpenParenthesisTkn:  $4.(*ArgumentList).OpenParenthesisTkn,
                    Args:                $4.(*ArgumentList).Arguments,
                    SeparatorTkns:       $4.(*ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $4.(*ArgumentList).CloseParenthesisTkn,
                }
            }
    |   variable_without_objects function_call_parameter_list
            {
                $$ = &ast.ExprFunctionCall{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $2),
                    Function:            $1,
                    OpenParenthesisTkn:  $2.(*ArgumentList).OpenParenthesisTkn,
                    Args:                $2.(*ArgumentList).Arguments,
                    SeparatorTkns:       $2.(*ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $2.(*ArgumentList).CloseParenthesisTkn,
                }
            }
;

class_name:
        T_STATIC
            {
                $$ = &ast.Identifier{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   namespace_name
            {
                $$ = &ast.Name{
                    Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                    Parts:         $1.(*ParserSeparatedList).Items,
                    SeparatorTkns: $1.(*ParserSeparatedList).SeparatorTkns,
                }
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameRelative{
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $3.(*ParserSeparatedList).Items),
                    NsTkn:          $1,
                    NsSeparatorTkn: $2,
                    Parts:          $3.(*ParserSeparatedList).Items,
                    SeparatorTkns:  $3.(*ParserSeparatedList).SeparatorTkns,
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameFullyQualified{
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $2.(*ParserSeparatedList).Items),
                    NsSeparatorTkn: $1,
                    Parts:          $2.(*ParserSeparatedList).Items,
                    SeparatorTkns:  $2.(*ParserSeparatedList).SeparatorTkns,
                }
            }
;

fully_qualified_class_name:
        namespace_name
            {
                $$ = &ast.Name{
                    Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                    Parts:         $1.(*ParserSeparatedList).Items,
                    SeparatorTkns: $1.(*ParserSeparatedList).SeparatorTkns,
                }
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameRelative{
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $3.(*ParserSeparatedList).Items),
                    NsTkn:          $1,
                    NsSeparatorTkn: $2,
                    Parts:          $3.(*ParserSeparatedList).Items,
                    SeparatorTkns:  $3.(*ParserSeparatedList).SeparatorTkns,
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameFullyQualified{
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $2.(*ParserSeparatedList).Items),
                    NsSeparatorTkn: $1,
                    Parts:          $2.(*ParserSeparatedList).Items,
                    SeparatorTkns:  $2.(*ParserSeparatedList).SeparatorTkns,
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

                $3[0].(*ast.ExprPropertyFetch).ObjectOperatorTkn = $2

                for _, n := range($3) {
                    switch nn := n.(type) {
                        case *ast.ExprArrayDimFetch:
                            nn.Var = $$
                            *$$.GetPosition() = *yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn

                        case *ast.ExprPropertyFetch:
                            nn.Var = $$
                            *$$.GetPosition() = *yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn
                    }
                }

                for _, n := range($4) {
                    switch nn := n.(type) {
                        case *ast.ExprArrayDimFetch:
                            nn.Var = $$
                            *$$.GetPosition() = *yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn

                        case *ast.ExprPropertyFetch:
                            nn.Var = $$
                            *$$.GetPosition() = *yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn
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
                $2[0].(*ast.ExprPropertyFetch).ObjectOperatorTkn = $1

                $$ = $2
            }
;

exit_expr:
        /* empty */
            {
                $$ = nil
            }
    |   '(' ')'
            {
                $$ = &ast.ExprBrackets{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $2),
                    OpenParenthesisTkn:  $1,
                    CloseParenthesisTkn: $2,
                }
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
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
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
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    NumberTkn: $1,
                    Value:     $1.Value,
                }
            }
    |   T_DNUMBER
            {
                $$ = &ast.ScalarDnumber{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    NumberTkn: $1,
                    Value:     $1.Value,
                }
            }
    |   T_CONSTANT_ENCAPSED_STRING
            {
                $$ = &ast.ScalarString{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    StringTkn: $1,
                    Value:     $1.Value,
                }
            }
    |   T_LINE
            {
                $$ = &ast.ScalarMagicConstant{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_FILE
            {
                $$ = &ast.ScalarMagicConstant{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_DIR
            {
                $$ = &ast.ScalarMagicConstant{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_TRAIT_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_METHOD_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_FUNC_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_NS_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenHeredocTkn: $1,
                    Parts: []ast.Vertex{
                        &ast.ScalarEncapsedStringPart{
                            Position: yylex.(*Parser).builder.NewTokenPosition($2),
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
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $2),
                    OpenHeredocTkn:  $1,
                    CloseHeredocTkn: $2,
                }
            }
;

static_class_constant:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                $$ = &ast.ExprClassConstFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $3),
                    Class:          $1,
                    DoubleColonTkn: $2,
                    Const: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
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
                $$ = &ast.ExprConstFetch{
                    Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                    Const: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                        Parts:         $1.(*ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.ExprConstFetch{
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $3.(*ParserSeparatedList).Items),
                    Const: &ast.NameRelative{
                        Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $3.(*ParserSeparatedList).Items),
                        NsTkn:          $1,
                        NsSeparatorTkn: $2,
                        Parts:          $3.(*ParserSeparatedList).Items,
                        SeparatorTkns:  $3.(*ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.ExprConstFetch{
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $2.(*ParserSeparatedList).Items),
                    Const: &ast.NameFullyQualified{
                        Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $2.(*ParserSeparatedList).Items),
                        NsSeparatorTkn: $1,
                        Parts:          $2.(*ParserSeparatedList).Items,
                        SeparatorTkns:  $2.(*ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
    |   T_ARRAY '(' static_array_pair_list ')'
            {
                $$ = &ast.ExprArray{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    ArrayTkn:        $1,
                    OpenBracketTkn:  $2,
                    Items:           $3.(*ParserSeparatedList).Items,
                    SeparatorTkns:   $3.(*ParserSeparatedList).SeparatorTkns,
                    CloseBracketTkn: $4,
                }
            }
    |   '[' static_array_pair_list ']'
            {
                $$ = &ast.ExprArray{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenBracketTkn:  $1,
                    Items:           $2.(*ParserSeparatedList).Items,
                    SeparatorTkns:   $2.(*ParserSeparatedList).SeparatorTkns,
                    CloseBracketTkn: $3,
                }
            }
    |   static_class_constant
            {
                $$ = $1
            }
    |   T_CLASS_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
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
                $$ = &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $4),
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   static_scalar_value '+' static_scalar_value
            {
                $$ = &ast.ExprBinaryPlus{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value '-' static_scalar_value
            {
                $$ = &ast.ExprBinaryMinus{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value '*' static_scalar_value
            {
                $$ = &ast.ExprBinaryMul{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_POW static_scalar_value
            {
                $$ = &ast.ExprBinaryPow{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value '/' static_scalar_value
            {
                $$ = &ast.ExprBinaryDiv{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value '%' static_scalar_value
            {
                $$ = &ast.ExprBinaryMod{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   '!' static_scalar_value
            {
                $$ = &ast.ExprBooleanNot{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    ExclamationTkn: $1,
                    Expr:           $2,
                }
            }
    |   '~' static_scalar_value
            {
                $$ = &ast.ExprBitwiseNot{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    TildaTkn: $1,
                    Expr:     $2,
                }
            }
    |   static_scalar_value '|' static_scalar_value
            {
                $$ = &ast.ExprBinaryBitwiseOr{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value '&' static_scalar_value
            {
                $$ = &ast.ExprBinaryBitwiseAnd{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value '^' static_scalar_value
            {
                $$ = &ast.ExprBinaryBitwiseXor{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_SL static_scalar_value
            {
                $$ = &ast.ExprBinaryShiftLeft{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_SR static_scalar_value
            {
                $$ = &ast.ExprBinaryShiftRight{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value '.' static_scalar_value
            {
                $$ = &ast.ExprBinaryConcat{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_LOGICAL_XOR static_scalar_value
            {
                $$ = &ast.ExprBinaryLogicalXor{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_LOGICAL_AND static_scalar_value
            {
                $$ = &ast.ExprBinaryLogicalAnd{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_LOGICAL_OR static_scalar_value
            {
                $$ = &ast.ExprBinaryLogicalOr{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_BOOLEAN_AND static_scalar_value
            {
                $$ = &ast.ExprBinaryBooleanAnd{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_BOOLEAN_OR static_scalar_value
            {
                $$ = &ast.ExprBinaryBooleanOr{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_IS_IDENTICAL static_scalar_value
            {
                $$ = &ast.ExprBinaryIdentical{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_IS_NOT_IDENTICAL static_scalar_value
            {
                $$ = &ast.ExprBinaryNotIdentical{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_IS_EQUAL static_scalar_value
            {
                $$ = &ast.ExprBinaryEqual{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_IS_NOT_EQUAL static_scalar_value
            {
                $$ = &ast.ExprBinaryNotEqual{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value '<' static_scalar_value
            {
                $$ = &ast.ExprBinarySmaller{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value '>' static_scalar_value
            {
                $$ = &ast.ExprBinaryGreater{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_IS_SMALLER_OR_EQUAL static_scalar_value
            {
                $$ = &ast.ExprBinarySmallerOrEqual{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value T_IS_GREATER_OR_EQUAL static_scalar_value
            {
                $$ = &ast.ExprBinaryGreaterOrEqual{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   static_scalar_value '?' ':' static_scalar_value
            {
                $$ = &ast.ExprTernary{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $4),
                    Cond:        $1,
                    QuestionTkn: $2,
                    ColonTkn:    $3,
                    IfFalse:     $4,
                }
            }
    |   static_scalar_value '?' static_scalar_value ':' static_scalar_value
            {
                $$ = &ast.ExprTernary{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $5),
                    Cond:        $1,
                    QuestionTkn: $2,
                    IfTrue:      $3,
                    ColonTkn:    $4,
                    IfFalse:     $5,
                }
            }
    |   '+' static_scalar_value
            {
                $$ = &ast.ExprUnaryPlus{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    PlusTkn: $1,
                    Expr:    $2,
                }
            }
    |   '-' static_scalar_value
            {
                $$ = &ast.ExprUnaryMinus{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    MinusTkn: $1,
                    Expr:     $2,
                }
            }
    |   '(' static_scalar_value ')'
            {
                $$ = &ast.ExprBrackets{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenParenthesisTkn:  $1,
                    Expr:                $2,
                    CloseParenthesisTkn: $3,
                }
            }
;

general_constant:
        class_constant
            {
                $$ = $1
            }
    |   namespace_name
            {
                $$ = &ast.ExprConstFetch{
                    Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                    Const: &ast.Name{
                        Position: yylex.(*Parser).builder.NewNodeListPosition($1.(*ParserSeparatedList).Items),
                        Parts:         $1.(*ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.ExprConstFetch{
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $3.(*ParserSeparatedList).Items),
                    Const: &ast.NameRelative{
                        Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $3.(*ParserSeparatedList).Items),
                        NsTkn:          $1,
                        NsSeparatorTkn: $2,
                        Parts:          $3.(*ParserSeparatedList).Items,
                        SeparatorTkns:  $3.(*ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.ExprConstFetch{
                    Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $2.(*ParserSeparatedList).Items),
                    Const: &ast.NameFullyQualified{
                        Position: yylex.(*Parser).builder.NewTokenNodeListPosition($1, $2.(*ParserSeparatedList).Items),
                        NsSeparatorTkn: $1,
                        Parts:          $2.(*ParserSeparatedList).Items,
                        SeparatorTkns:  $2.(*ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
;

scalar:
        T_STRING_VARNAME
            {
                $$ = &ast.ExprVariable{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
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
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenQuoteTkn:  $1,
                    Parts:         $2,
                    CloseQuoteTkn: $3,
                }
            }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenHeredocTkn:  $1,
                    Parts:           $2,
                    CloseHeredocTkn: $3,
                }
            }
    |   T_CLASS_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
;

static_array_pair_list:
        /* empty */
            {
                $$ = &ParserSeparatedList{}
            }
    |   non_empty_static_array_pair_list possible_comma
            {
                if $2 != nil {
                    $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                    $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, &ast.ExprArrayItem{})
                }

                $$ = $1
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
                arrayItem := &ast.ExprArrayItem{
                    Position: yylex.(*Parser).builder.NewNodesPosition($3, $5),
                    Key:            $3,
                    DoubleArrowTkn: $4,
                    Val:            $5,
                }

                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, arrayItem)

                $$ = $1
            }
    |   non_empty_static_array_pair_list ',' static_scalar_value
            {
                arrayItem := &ast.ExprArrayItem{
                    Position: yylex.(*Parser).builder.NewNodePosition($3),
                    Val: $3,
                }

                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, arrayItem)

                $$ = $1
            }
    |   static_scalar_value T_DOUBLE_ARROW static_scalar_value
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{
                        &ast.ExprArrayItem{
                            Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                            Key:            $1,
                            DoubleArrowTkn: $2,
                            Val:            $3,
                        },
                    },
                }
            }
    |   static_scalar_value
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{
                        &ast.ExprArrayItem{
                            Position: yylex.(*Parser).builder.NewNodePosition($1),
                            Val: $1,
                        },
                    },
                }
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
                $$ = &ast.ExprBrackets{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenParenthesisTkn:  $1,
                    Expr:                $2,
                    CloseParenthesisTkn: $3,
                }
            }
    |   '(' yield_expr ')'
            {
                $$ = &ast.ExprBrackets{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenParenthesisTkn:  $1,
                    Expr:                $2,
                    CloseParenthesisTkn: $3,
                }
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

                $3[0].(*ast.ExprPropertyFetch).ObjectOperatorTkn = $2

                if $4 != nil {
                    last := $3[len($3)-1]
                    switch l := last.(type) {
                        case *ast.ExprArrayDimFetch:
                            mc := $4[0].(*ast.ExprMethodCall)
                            $3 = append($3, &ast.ExprFunctionCall{
                                    Position: yylex.(*Parser).builder.NewNodePosition(mc),
                                    OpenParenthesisTkn:  mc.OpenParenthesisTkn,
                                    Args:                mc.Args,
                                    SeparatorTkns:       mc.SeparatorTkns,
                                    CloseParenthesisTkn: mc.CloseParenthesisTkn,
                                },
                            )
                            $3 = append($3, $4[1:len($4)]...)
                        case *ast.ExprPropertyFetch:
                            $4[0].(*ast.ExprMethodCall).OpenCurlyBracketTkn = l.OpenCurlyBracketTkn
                            $4[0].(*ast.ExprMethodCall).Method = l.Prop
                            $4[0].(*ast.ExprMethodCall).CloseCurlyBracketTkn = l.CloseCurlyBracketTkn
                            $4[0].(*ast.ExprMethodCall).ObjectOperatorTkn = l.ObjectOperatorTkn
                            $3 = append($3[:len($3)-1], $4...)
                    }
                }

                for _, n := range($3) {
                    switch nn := n.(type) {
                        case *ast.ExprFunctionCall:
                            nn.Function = $$
                            nn.Position = yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn

                        case *ast.ExprArrayDimFetch:
                            nn.Var = $$
                            nn.Position = yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn

                        case *ast.ExprPropertyFetch:
                            nn.Var = $$
                            nn.Position = yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn

                        case *ast.ExprMethodCall:
                            nn.Var = $$
                            nn.Position = yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn
                    }
                }

                for _, n := range($5) {
                    switch nn := n.(type) {
                        case *ast.ExprFunctionCall:
                            nn.Function = $$
                            nn.Position = yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn

                        case *ast.ExprArrayDimFetch:
                            nn.Var = $$
                            nn.Position = yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn

                        case *ast.ExprPropertyFetch:
                            nn.Var = $$
                            nn.Position = yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn

                        case *ast.ExprMethodCall:
                            nn.Var = $$
                            nn.Position = yylex.(*Parser).builder.NewNodesPosition($$, nn)
                            $$ = nn
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
                $2[0].(*ast.ExprPropertyFetch).ObjectOperatorTkn = $1

                if $3 != nil {
                    last := $2[len($2)-1]
                    switch l := last.(type) {
                        case *ast.ExprArrayDimFetch:
                            mc := $3[0].(*ast.ExprMethodCall)
                            $2 = append($2, &ast.ExprFunctionCall{
                                    Position: yylex.(*Parser).builder.NewNodePosition(mc),
                                    OpenParenthesisTkn:  mc.OpenParenthesisTkn,
                                    Args:                mc.Args,
                                    SeparatorTkns:       mc.SeparatorTkns,
                                    CloseParenthesisTkn: mc.OpenParenthesisTkn,
                                },
                            )
                            $2 = append($2, $3[1:len($3)]...)
                        case *ast.ExprPropertyFetch:
                            $3[0].(*ast.ExprMethodCall).OpenCurlyBracketTkn = l.OpenCurlyBracketTkn
                            $3[0].(*ast.ExprMethodCall).Method = l.Prop
                            $3[0].(*ast.ExprMethodCall).CloseCurlyBracketTkn = l.CloseCurlyBracketTkn
                            $3[0].(*ast.ExprMethodCall).ObjectOperatorTkn = l.ObjectOperatorTkn
                            $2 = append($2[:len($2)-1], $3...)
                    }
                }

                $$ = $2
            }
;

array_method_dereference:
        array_method_dereference '[' dim_offset ']'
            {
                fetch := &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($2, $4),
                    Var:             nil,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }

                $$ = append($1, fetch)
            }
    |   method '[' dim_offset ']'
            {
                fetch := &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($2, $4),
                    Var:             nil,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }

                $$ = []ast.Vertex{$1, fetch}
            }
;

method:
        function_call_parameter_list
            {
                $$ = &ast.ExprMethodCall{
                    Position: yylex.(*Parser).builder.NewNodePosition($1),
                    OpenParenthesisTkn:  $1.(*ArgumentList).OpenParenthesisTkn,
                    Args:                $1.(*ArgumentList).Arguments,
                    SeparatorTkns:       $1.(*ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $1.(*ArgumentList).CloseParenthesisTkn,
                }
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
                for i := len($1)-1; i>=0; i-- {
                    $1[i].(*ast.ExprVariable).Name = $2
                    $1[i].(*ast.ExprVariable).Position = yylex.(*Parser).builder.NewNodesPosition($1[i], $2)
                    $2 = $1[i]
                }

                $$ = $1[0]
            }
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects
            {
                $$ = &ast.ExprStaticPropertyFetch{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Class:          $1,
                    DoubleColonTkn: $2,
                    Prop:           $3,
                }
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects
            {
                $$ = &ast.ExprStaticPropertyFetch{
                    Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                    Class:          $1,
                    DoubleColonTkn: $2,
                    Prop:           $3,
                }
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
                $$ = &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $4),
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   function_call '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $4),
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
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
                for i := len($1)-1; i>=0; i-- {
                    $1[i].(*ast.ExprVariable).Name = $2
                    $1[i].(*ast.ExprVariable).Position = yylex.(*Parser).builder.NewNodesPosition($1[i], $2)
                    $2 = $1[i]
                }

                $$ = $1[0]
            }
    |   static_member
            {
                $$ = $1
            }
;

reference_variable:
        reference_variable '[' dim_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $4),
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   reference_variable '{' expr '}'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $4),
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   compound_variable
            {
                $$ = $1
            }
;


compound_variable:
        T_VARIABLE
            {
                $$ = &ast.ExprVariable{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
            }
    |   '$' '{' expr '}'
            {
                $$ = &ast.ExprVariable{
                    Position:             yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    DollarTkn:            $1,
                    OpenCurlyBracketTkn:  $2,
                    Name:                 $3,
                    CloseCurlyBracketTkn: $4,
                }
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
                $$ = []ast.Vertex{
                    &ast.ExprPropertyFetch{
                        Position: yylex.(*Parser).builder.NewNodePosition($1),
                        Prop:     $1,
                    },
                }
            }
;

object_dim_list:
        object_dim_list '[' dim_offset ']'
            {
                fetch := &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($2, $4),
                    Var:             nil,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }

                $$ = append($1, fetch)
            }
    |   object_dim_list '{' expr '}'
            {
                fetch := &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($2, $4),
                    Var:             nil,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }

                $$ = append($1, fetch)
            }
    |   variable_name
            {
                property := &ast.ExprPropertyFetch{
                    Position: yylex.(*Parser).builder.NewNodePosition($1),
                    Prop:     $1,
                }

                if brackets, ok := $1.(*ParserBrackets); ok {
                    property.OpenCurlyBracketTkn  = brackets.OpenBracketTkn
                    property.Prop                 = brackets.Child
                    property.CloseCurlyBracketTkn = brackets.CloseBracketTkn
                }

                $$ = []ast.Vertex{ property }
            }
;

variable_name:
        T_STRING
            {
                $$ = &ast.Identifier{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   '{' expr '}'
            {
                $$ = &ParserBrackets{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenBracketTkn:  $1,
                    Child:           $2,
                    CloseBracketTkn: $3,
                }
            }
;

simple_indirect_reference:
        '$'
            {
                $$ = []ast.Vertex{
                    &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        DollarTkn: $1,
                    },
                }
            }
    |   simple_indirect_reference '$'
            {
                $$ = append($1, &ast.ExprVariable{
                    Position: yylex.(*Parser).builder.NewTokenPosition($2),
                    DollarTkn: $2,
                })
            }
;

assignment_list:
        assignment_list ',' assignment_list_element
            {
                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   assignment_list_element
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
;


assignment_list_element:
        variable
            {
                $$ = &ast.ExprArrayItem{
                    Position: yylex.(*Parser).builder.NewNodePosition($1),
                    Val: $1,
                }
            }
    |   T_LIST '(' assignment_list ')'
            {
                pairList := $3.(*ParserSeparatedList)
                fistPair := pairList.Items[0].(*ast.ExprArrayItem)

                if fistPair.Key == nil && fistPair.Val == nil && len(pairList.Items) == 1 {
                    pairList.Items = nil
                }

                $$ = &ast.ExprArrayItem{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    Val: &ast.ExprList{
                        Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                        ListTkn:         $1,
                        OpenBracketTkn:  $2,
                        Items:           $3.(*ParserSeparatedList).Items,
                        SeparatorTkns:   $3.(*ParserSeparatedList).SeparatorTkns,
                        CloseBracketTkn: $4,
                    },
                }
            }
    |   /* empty */
            {
                $$ = &ast.ExprArrayItem{}
            }
;


array_pair_list:
        /* empty */
            {
                $$ = &ParserSeparatedList{}
            }
    |   non_empty_array_pair_list possible_comma
            {
                if $2 != nil {
                    $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                    $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, &ast.ExprArrayItem{})
                }

                $$ = $1
            }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' expr T_DOUBLE_ARROW expr
            {
                arrayItem := &ast.ExprArrayItem{
                    Position: yylex.(*Parser).builder.NewNodesPosition($3, $5),
                    Key:            $3,
                    DoubleArrowTkn: $4,
                    Val:            $5,
                }

                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, arrayItem)

                $$ = $1
            }
    |   non_empty_array_pair_list ',' expr
            {
                arrayItem := &ast.ExprArrayItem{
                    Position: yylex.(*Parser).builder.NewNodePosition($3),
                    Val:            $3,
                }

                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, arrayItem)

                $$ = $1
            }
    |   expr T_DOUBLE_ARROW expr
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{
                        &ast.ExprArrayItem{
                            Position: yylex.(*Parser).builder.NewNodesPosition($1, $3),
                            Key:            $1,
                            DoubleArrowTkn: $2,
                            Val:            $3,
                        },
                    },
                }
            }
    |   expr
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{
                        &ast.ExprArrayItem{
                            Position: yylex.(*Parser).builder.NewNodePosition($1),
                            Val: $1,
                        },
                    },
                }
            }
    |   non_empty_array_pair_list ',' expr T_DOUBLE_ARROW '&' w_variable
            {
                arrayItem := &ast.ExprArrayItem{
                    Position:       yylex.(*Parser).builder.NewNodesPosition($3, $6),
                    Key:            $3,
                    DoubleArrowTkn: $4,
                    AmpersandTkn:   $5,
                    Val:            $6,
                }

                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, arrayItem)

                $$ = $1
            }
    |   non_empty_array_pair_list ',' '&' w_variable
            {
                arrayItem := &ast.ExprArrayItem{
                    Position:     yylex.(*Parser).builder.NewTokenNodePosition($3, $4),
                    AmpersandTkn: $3,
                    Val:          $4,
                }

                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, arrayItem)

                $$ = $1
            }
    |   expr T_DOUBLE_ARROW '&' w_variable
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{
                        &ast.ExprArrayItem{
                            Position:       yylex.(*Parser).builder.NewNodesPosition($1, $4),
                            Key:            $1,
                            DoubleArrowTkn: $2,
                            AmpersandTkn:   $3,
                            Val:            $4,
                        },
                    },
                }
            }
    |   '&' w_variable
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{
                        &ast.ExprArrayItem{
                            Position:     yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                            AmpersandTkn: $1,
                            Val:          $2,
                        },
                    },
                }
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
                        Position: yylex.(*Parser).builder.NewTokenPosition($2),
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
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
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
                $$ = &ast.ExprVariable{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
            }
    |   T_VARIABLE '[' encaps_var_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($1),
                            IdentifierTkn: $1,
                            Value:         $1.Value,
                        },
                    },
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING
            {
                $$ = &ast.ExprPropertyFetch{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    Var: &ast.ExprVariable{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        Name: &ast.Identifier{
                            Position: yylex.(*Parser).builder.NewTokenPosition($1),
                            IdentifierTkn: $1,
                            Value:         $1.Value,
                        },
                    },
                    ObjectOperatorTkn: $2,
                    Prop: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES expr '}'
            {
                $$ = &ast.ScalarEncapsedStringVar{
                    Position:                  yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    DollarOpenCurlyBracketTkn: $1,
                    Name:                      $2,
                    CloseCurlyBracketTkn:      $3,
                }
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
            {
                $$ = &ast.ScalarEncapsedStringVar{
                    Position:                  yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    DollarOpenCurlyBracketTkn: $1,
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($2),
                        IdentifierTkn: $2,
                        Value:         $2.Value,
                    },
                    CloseCurlyBracketTkn: $3,
                }
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
            {
                $$ = &ast.ScalarEncapsedStringVar{
                    Position:                  yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    DollarOpenCurlyBracketTkn: $1,
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($2),
                        IdentifierTkn: $2,
                        Value:         $2.Value,
                    },
                    OpenSquareBracketTkn:  $3,
                    Dim:                   $4,
                    CloseSquareBracketTkn: $5,
                    CloseCurlyBracketTkn:  $6,
                }
            }
    |   T_CURLY_OPEN variable '}'
            {
                $$ = &ast.ScalarEncapsedStringBrackets{
                    Position:             yylex.(*Parser).builder.NewTokensPosition($1, $3),
                    OpenCurlyBracketTkn:  $1,
                    Var:                  $2,
                    CloseCurlyBracketTkn: $3,
                }
            }
;

encaps_var_offset:
        T_STRING
            {
                $$ = &ast.ScalarString{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    StringTkn: $1,
                    Value:     $1.Value,
                }
            }
    |   T_NUM_STRING
            {
                // TODO: add option to handle 64 bit integer
                if _, err := strconv.Atoi(string($1.Value)); err == nil {
                    $$ = &ast.ScalarLnumber{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        NumberTkn: $1,
                        Value:     $1.Value,
                    }
                } else {
                    $$ = &ast.ScalarString{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        StringTkn: $1,
                        Value:     $1.Value,
                    }
                }
            }
    |   T_VARIABLE
            {
                $$ = &ast.ExprVariable{
                    Position: yylex.(*Parser).builder.NewTokenPosition($1),
                    Name: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($1),
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
            }
;

internal_functions_in_yacc:
        T_ISSET '(' isset_variables ')'
            {
                $$ = &ast.ExprIsset{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    IssetTkn:            $1,
                    OpenParenthesisTkn:  $2,
                    Vars:                $3.(*ParserSeparatedList).Items,
                    SeparatorTkns:       $3.(*ParserSeparatedList).SeparatorTkns,
                    CloseParenthesisTkn: $4,
                }
            }
    |   T_EMPTY '(' variable ')'
            {
                $$ = &ast.ExprEmpty{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    EmptyTkn:            $1,
                    OpenParenthesisTkn:  $2,
                    Expr:                $3,
                    CloseParenthesisTkn: $4,
                }
            }
    |   T_EMPTY '(' expr ')'
            {
                $$ = &ast.ExprEmpty{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    EmptyTkn:            $1,
                    OpenParenthesisTkn:  $2,
                    Expr:                $3,
                    CloseParenthesisTkn: $4,
                }
            }
    |   T_INCLUDE expr
            {
                $$ = &ast.ExprInclude{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    IncludeTkn: $1,
                    Expr:       $2,
                }
            }
    |   T_INCLUDE_ONCE expr
            {
                $$ = &ast.ExprIncludeOnce{
                    Position:       yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    IncludeOnceTkn: $1,
                    Expr:           $2,
                }
            }
    |   T_EVAL '(' expr ')'
            {
                $$ = &ast.ExprEval{
                    Position: yylex.(*Parser).builder.NewTokensPosition($1, $4),
                    EvalTkn:             $1,
                    OpenParenthesisTkn:  $2,
                    Expr:                $3,
                    CloseParenthesisTkn: $4,
                }
            }
    |   T_REQUIRE expr
            {
                $$ = &ast.ExprRequire{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    RequireTkn: $1,
                    Expr:       $2,
                }
            }
    |   T_REQUIRE_ONCE expr
            {
                $$ = &ast.ExprRequireOnce{
                    Position: yylex.(*Parser).builder.NewTokenNodePosition($1, $2),
                    RequireOnceTkn: $1,
                    Expr:           $2,
                }
            }
;

isset_variables:
        isset_variable
            {
                $$ = &ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
    |   isset_variables ',' isset_variable
            {
                $1.(*ParserSeparatedList).SeparatorTkns = append($1.(*ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ParserSeparatedList).Items = append($1.(*ParserSeparatedList).Items, $3)

                $$ = $1
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
                $$ = &ast.ExprClassConstFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $3),
                    Class:          $1,
                    DoubleColonTkn: $2,
                    Const: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                $$ = &ast.ExprClassConstFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $3),
                    Class:          $1,
                    DoubleColonTkn: $2,
                    Const: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
;

static_class_name_scalar:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS
            {
                $$ = &ast.ExprClassConstFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $3),
                    Class:          $1,
                    DoubleColonTkn: $2,
                    Const: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
;

class_name_scalar:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS
            {
                $$ = &ast.ExprClassConstFetch{
                    Position: yylex.(*Parser).builder.NewNodeTokenPosition($1, $3),
                    Class:          $1,
                    DoubleColonTkn: $2,
                    Const: &ast.Identifier{
                        Position: yylex.(*Parser).builder.NewTokenPosition($3),
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
;

%%
