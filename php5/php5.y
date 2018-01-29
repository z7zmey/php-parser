%{
package php5

import (
//    "strings"
//    "strconv"

    "github.com/z7zmey/php-parser/token"
    "github.com/z7zmey/php-parser/node"
//    "github.com/z7zmey/php-parser/node/scalar"
    "github.com/z7zmey/php-parser/node/name"
    "github.com/z7zmey/php-parser/node/stmt"
//    "github.com/z7zmey/php-parser/node/expr"
//    "github.com/z7zmey/php-parser/node/expr/assign_op"
//    "github.com/z7zmey/php-parser/node/expr/binary_op"
//    "github.com/z7zmey/php-parser/node/expr/cast"
)

%}

%union{
    node node.Node
    token token.Token
//    boolWithToken boolWithToken
    list []node.Node
//    foreachVariable foreachVariable
//    nodesWithEndToken *nodesWithEndToken
//    str string
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

%type <node> top_statement use_declaration use_function_declaration use_const_declaration

%type <list> top_statement_list namespace_name use_declarations use_function_declarations use_const_declarations

%%

start:
        top_statement_list
            {
                rootnode = stmt.NewStmtList($1)
            }
;

top_statement_list:
        top_statement_list top_statement                { $$ = append($1, $2) }
    |   /* empty */                                     { $$ = []node.Node{} }
;

namespace_name:
        T_STRING
            {
                namePart := name.NewNamePart($1.Value)
                positions.AddPosition(namePart, positionBuilder.NewTokenPosition($1))
                $$ = []node.Node{namePart}
                comments.AddComments(namePart, $1.Comments())
            }
    |   namespace_name T_NS_SEPARATOR T_STRING
            {
                namePart := name.NewNamePart($3.Value)
                positions.AddPosition(namePart, positionBuilder.NewTokenPosition($3))
                $$ = append($1, namePart)
                comments.AddComments(namePart, $3.Comments())
            }
;

top_statement:
        statement                                       { $$ = nil }
    |   function_declaration_statement                  { $$ = nil }
    |   class_declaration_statement                     { $$ = nil }
    |   T_HALT_COMPILER '(' ')' ';'                     { $$ = stmt.NewHaltCompiler() }
    |   T_NAMESPACE namespace_name ';'
            {
                name := name.NewName($2)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($2))
                $$ = stmt.NewNamespace(name, nil)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))

                comments.AddComments(name, ListGetFirstNodeComments($2))
                comments.AddComments($$, $1.Comments())
            }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
            {
                name := name.NewName($2)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($2))
                $$ = stmt.NewNamespace(name, $4)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $5))

                comments.AddComments(name, ListGetFirstNodeComments($2))
                comments.AddComments($$, $1.Comments())
            }
    |   T_NAMESPACE '{' top_statement_list '}'
            {
                $$ = stmt.NewNamespace(nil, $3)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
                comments.AddComments($$, $1.Comments())
            }
    |   T_USE use_declarations ';'
            {
                $$ = stmt.NewUseList(nil, $2)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
                comments.AddComments($$, $1.Comments())
            }
    |   T_USE T_FUNCTION use_function_declarations ';'
            {
                useType := node.NewIdentifier($2.Value)
                positions.AddPosition($$, positionBuilder.NewTokenPosition($2))
                comments.AddComments($$, $2.Comments())

                $$ = stmt.NewUseList(useType, $3)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
                comments.AddComments($$, $1.Comments())
            }
    |   T_USE T_CONST use_const_declarations ';'
            {
                useType := node.NewIdentifier($2.Value)
                positions.AddPosition($$, positionBuilder.NewTokenPosition($2))
                comments.AddComments($$, $2.Comments())

                $$ = stmt.NewUseList(useType, $3)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
                comments.AddComments($$, $1.Comments())
            }
    |   constant_declaration ';'                        { $$ = nil }
;

use_declarations:
        use_declarations ',' use_declaration            { $$ = append($1, $3) }
    |   use_declaration                                 { $$ = []node.Node{$1} }
;

use_declaration:
        namespace_name
            {
                name := name.NewName($1)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($1))
                $$ = stmt.NewUse(nil, name, nil)
                positions.AddPosition($$, positionBuilder.NewNodeListPosition($1))

                comments.AddComments(name, ListGetFirstNodeComments($1))
                comments.AddComments($$, ListGetFirstNodeComments($1))
            }
    |   namespace_name T_AS T_STRING
            {
                name := name.NewName($1)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($1))
                alias := node.NewIdentifier($3.Value)
                positions.AddPosition(alias, positionBuilder.NewTokenPosition($3))
                $$ = stmt.NewUse(nil, name, alias)
                positions.AddPosition($$, positionBuilder.NewNodeListTokenPosition($1, $3))

                comments.AddComments(name, ListGetFirstNodeComments($1))
                comments.AddComments(alias, $3.Comments())
                comments.AddComments($$, ListGetFirstNodeComments($1))
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewName($2)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($2))
                $$ = stmt.NewUse(nil, name, nil)
                positions.AddPosition($$, positionBuilder.NewNodeListPosition($2))

                comments.AddComments(name, ListGetFirstNodeComments($2))
                comments.AddComments($$, ListGetFirstNodeComments($2))
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                name := name.NewName($2)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($2))
                alias := node.NewIdentifier($4.Value)
                positions.AddPosition(alias, positionBuilder.NewTokenPosition($4))
                $$ = stmt.NewUse(nil, name, alias)
                positions.AddPosition($$, positionBuilder.NewNodeListTokenPosition($2, $4))

                comments.AddComments(name, ListGetFirstNodeComments($2))
                comments.AddComments(alias, $4.Comments())
                comments.AddComments($$, ListGetFirstNodeComments($2))
            }
;

use_function_declarations:
        use_function_declarations ',' use_function_declaration
            {
                $$ = append($1, $3)
            }
    |   use_function_declaration
            {
                $$ = []node.Node{$1}
            }
;

use_function_declaration:
        namespace_name
            {
                name := name.NewName($1)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($1))
                $$ = stmt.NewUse(nil, name, nil)
                positions.AddPosition($$, positionBuilder.NewNodeListPosition($1))

                comments.AddComments(name, ListGetFirstNodeComments($1))
                comments.AddComments($$, ListGetFirstNodeComments($1))
            }
    |   namespace_name T_AS T_STRING
            {
                name := name.NewName($1)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($1))
                alias := node.NewIdentifier($3.Value)
                positions.AddPosition(alias, positionBuilder.NewTokenPosition($3))
                $$ = stmt.NewUse(nil, name, alias)
                positions.AddPosition($$, positionBuilder.NewNodeListTokenPosition($1, $3))

                comments.AddComments(name, ListGetFirstNodeComments($1))
                comments.AddComments(alias, $3.Comments())
                comments.AddComments($$, ListGetFirstNodeComments($1))
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewName($2)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($2))
                $$ = stmt.NewUse(nil, name, nil)
                positions.AddPosition($$, positionBuilder.NewNodeListPosition($2))

                comments.AddComments(name, ListGetFirstNodeComments($2))
                comments.AddComments($$, ListGetFirstNodeComments($2))
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                name := name.NewName($2)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($2))
                alias := node.NewIdentifier($4.Value)
                positions.AddPosition(alias, positionBuilder.NewTokenPosition($4))
                $$ = stmt.NewUse(nil, name, alias)
                positions.AddPosition($$, positionBuilder.NewNodeListTokenPosition($2, $4))

                comments.AddComments(name, ListGetFirstNodeComments($2))
                comments.AddComments(alias, $4.Comments())
                comments.AddComments($$, ListGetFirstNodeComments($2))
            }
;

use_const_declarations:
        use_const_declarations ',' use_const_declaration
            {
                $$ = append($1, $3)
            }
    |   use_const_declaration
            {
                $$ = []node.Node{$1}
            }
;

use_const_declaration:
        namespace_name
            {
                name := name.NewName($1)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($1))
                $$ = stmt.NewUse(nil, name, nil)
                positions.AddPosition($$, positionBuilder.NewNodeListPosition($1))

                comments.AddComments(name, ListGetFirstNodeComments($1))
                comments.AddComments($$, ListGetFirstNodeComments($1))
            }
    |   namespace_name T_AS T_STRING
            {
                name := name.NewName($1)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($1))
                alias := node.NewIdentifier($3.Value)
                positions.AddPosition(alias, positionBuilder.NewTokenPosition($3))
                $$ = stmt.NewUse(nil, name, alias)
                positions.AddPosition($$, positionBuilder.NewNodeListTokenPosition($1, $3))

                comments.AddComments(name, ListGetFirstNodeComments($1))
                comments.AddComments(alias, $3.Comments())
                comments.AddComments($$, ListGetFirstNodeComments($1))
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewName($2)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($2))
                $$ = stmt.NewUse(nil, name, nil)
                positions.AddPosition($$, positionBuilder.NewNodeListPosition($2))

                comments.AddComments(name, ListGetFirstNodeComments($2))
                comments.AddComments($$, ListGetFirstNodeComments($2))
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                name := name.NewName($2)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($2))
                alias := node.NewIdentifier($4.Value)
                positions.AddPosition(alias, positionBuilder.NewTokenPosition($4))
                $$ = stmt.NewUse(nil, name, alias)
                positions.AddPosition($$, positionBuilder.NewNodeListTokenPosition($2, $4))

                comments.AddComments(name, ListGetFirstNodeComments($2))
                comments.AddComments(alias, $4.Comments())
                comments.AddComments($$, ListGetFirstNodeComments($2))
            }
;

constant_declaration:
        constant_declaration ',' T_STRING '=' static_scalar	{  }
    |	T_CONST T_STRING '=' static_scalar {  }
;

inner_statement_list:
        inner_statement_list  {  }
    |	/* empty */
;


inner_statement:
        statement
    |	function_declaration_statement
    |	class_declaration_statement
    |	T_HALT_COMPILER '(' ')' ';'   {  }
;


statement:
        unticked_statement {  }
    |	T_STRING ':' {  }
;

unticked_statement:
        '{' inner_statement_list '}'
    |	T_IF parenthesis_expr {  } statement {  } elseif_list else_single {  }
    |	T_IF parenthesis_expr ':' {  } inner_statement_list {  } new_elseif_list new_else_single T_ENDIF ';' {  }
    |	T_WHILE {  } parenthesis_expr {  } while_statement {  }
    |	T_DO {  } statement T_WHILE {  } parenthesis_expr ';' {  }
    |	T_FOR
            '('
                for_expr
            ';' {  }
                for_expr
            ';' {  }
                for_expr
            ')' {  }
            for_statement {  }
    |	T_SWITCH parenthesis_expr	{  } switch_case_list {  }
    |	T_BREAK ';'				{  }
    |	T_BREAK expr ';'		{  }
    |	T_CONTINUE ';'			{  }
    |	T_CONTINUE expr ';'		{  }
    |	T_RETURN ';'						{  }
    |	T_RETURN expr_without_variable ';'	{  }
    |	T_RETURN variable ';'				{  }
    |	yield_expr ';' {  }
    |	T_GLOBAL global_var_list ';'
    |	T_STATIC static_var_list ';'
    |	T_ECHO echo_expr_list ';'
    |	T_INLINE_HTML			{  }
    |	expr ';'				{  }
    |	T_UNSET '(' unset_variables ')' ';'
    |	T_FOREACH '(' variable T_AS
        {  }
        foreach_variable foreach_optional_arg ')' {  }
        foreach_statement {  }
    |	T_FOREACH '(' expr_without_variable T_AS
        {  }
        foreach_variable foreach_optional_arg ')' {  }
        foreach_statement {  }
    |	T_DECLARE {  } '(' declare_list ')' declare_statement {  }
    |	';'		/* empty statement */
    |	T_TRY {  } '{' inner_statement_list '}'
        catch_statement {  }
        finally_statement {  }
    |	T_THROW expr ';' {  }
    |	T_GOTO T_STRING ';' {  }
;

catch_statement:
                /* empty */ {  }
    |	T_CATCH '(' {  } 
        fully_qualified_class_name {  }
        T_VARIABLE ')' {  }
        '{' inner_statement_list '}' {  }
        additional_catches {  }

finally_statement:
                    /* empty */ {  }
    |	T_FINALLY {  } '{' inner_statement_list '}' {  }
;

additional_catches:
        non_empty_additional_catches {  }
    |	/* empty */ {  }
;

non_empty_additional_catches:
        additional_catch {  }
    |	non_empty_additional_catches additional_catch {  }
;

additional_catch:
    T_CATCH '(' fully_qualified_class_name {  } T_VARIABLE ')' {  } '{' inner_statement_list '}' {  }
;

unset_variables:
        unset_variable
    |	unset_variables ',' unset_variable
;

unset_variable:
        variable	{  }
;

function_declaration_statement:
        unticked_function_declaration_statement	{  }
;

class_declaration_statement:
        unticked_class_declaration_statement	{  }
;

is_reference:
        /* empty */	{  }
    |	'&'			{  }
;

is_variadic:
        /* empty */ {  }
    |	T_ELLIPSIS  {  }
;

unticked_function_declaration_statement:
        function is_reference T_STRING {  }
        '(' parameter_list ')'
        '{' inner_statement_list '}' {  }
;

unticked_class_declaration_statement:
        class_entry_type T_STRING extends_from
            {  }
            implements_list
            '{'
                class_statement_list
            '}' {  }
    |	interface_entry T_STRING
            {  }
            interface_extends_list
            '{'
                class_statement_list
            '}' {  }
;


class_entry_type:
        T_CLASS			{  }
    |	T_ABSTRACT T_CLASS {  }
    |	T_TRAIT {  }
    |	T_FINAL T_CLASS {  }
;

extends_from:
        /* empty */					{  }
    |	T_EXTENDS fully_qualified_class_name	{  }
;

interface_entry:
    T_INTERFACE		{  }
;

interface_extends_list:
        /* empty */
    |	T_EXTENDS interface_list
;

implements_list:
        /* empty */
    |	T_IMPLEMENTS interface_list
;

interface_list:
        fully_qualified_class_name			{  }
    |	interface_list ',' fully_qualified_class_name {  }
;

foreach_optional_arg:
        /* empty */						{  }
    |	T_DOUBLE_ARROW foreach_variable	{  }
;

foreach_variable:
        variable			{  }
    |	'&' variable		{  }
    |	T_LIST '(' {  } assignment_list ')' {  }
;

for_statement:
        statement
    |	':' inner_statement_list T_ENDFOR ';'
;


foreach_statement:
        statement
    |	':' inner_statement_list T_ENDFOREACH ';'
;


declare_statement:
        statement
    |	':' inner_statement_list T_ENDDECLARE ';'
;


declare_list:
        T_STRING '=' static_scalar					{  }
    |	declare_list ',' T_STRING '=' static_scalar	{  }
;


switch_case_list:
        '{' case_list '}'					{  }
    |	'{' ';' case_list '}'				{  }
    |	':' case_list T_ENDSWITCH ';'		{  }
    |	':' ';' case_list T_ENDSWITCH ';'	{  }
;


case_list:
        /* empty */	{  }
    |	case_list T_CASE expr case_separator {  } inner_statement_list {  }
    |	case_list T_DEFAULT case_separator {  } inner_statement_list {  }
;


case_separator:
        ':'
    |	';'
;


while_statement:
        statement
    |	':' inner_statement_list T_ENDWHILE ';'
;



elseif_list:
        /* empty */
    |	elseif_list T_ELSEIF parenthesis_expr {  } statement {  }
;


new_elseif_list:
        /* empty */
    |	new_elseif_list T_ELSEIF parenthesis_expr ':' {  } inner_statement_list {  }
;


else_single:
        /* empty */
    |	T_ELSE statement
;


new_else_single:
        /* empty */
    |	T_ELSE ':' inner_statement_list
;


parameter_list:
        non_empty_parameter_list
    |	/* empty */
;


non_empty_parameter_list:
        parameter
    |	non_empty_parameter_list ',' parameter
;

parameter:
        optional_class_type is_reference is_variadic T_VARIABLE
            {  }
    |	optional_class_type is_reference is_variadic T_VARIABLE '=' static_scalar
            {  }
;


optional_class_type:
        /* empty */					{  }
    |	T_ARRAY						{  }
    |	T_CALLABLE					{  }
    |	fully_qualified_class_name			{  }
;


function_call_parameter_list:
        '(' ')'	{  }
    |	'(' non_empty_function_call_parameter_list ')'	{  }
    |	'(' yield_expr ')'	{  }
;


non_empty_function_call_parameter_list:
        function_call_parameter
    |	non_empty_function_call_parameter_list ',' function_call_parameter
;

function_call_parameter:
        expr_without_variable	{  }
    |	variable				{  }
    |	'&' w_variable 			{  }
    |	T_ELLIPSIS expr			{  }
;

global_var_list:
        global_var_list ',' global_var	{  }
    |	global_var						{  }
;


global_var:
        T_VARIABLE			{  }
    |	'$' r_variable		{  }
    |	'$' '{' expr '}'	{  }
;


static_var_list:
        static_var_list ',' T_VARIABLE {  }
    |	static_var_list ',' T_VARIABLE '=' static_scalar {  }
    |	T_VARIABLE  {  }
    |	T_VARIABLE '=' static_scalar {  }

;


class_statement_list:
        class_statement_list class_statement
    |	/* empty */
;


class_statement:
        variable_modifiers {  } class_variable_declaration ';'
    |	class_constant_declaration ';'
    |	trait_use_statement
    |	method_modifiers function is_reference T_STRING {  }
        '(' parameter_list ')'
        method_body {  }
;

trait_use_statement:
        T_USE trait_list trait_adaptations
;

trait_list:
        fully_qualified_class_name						{  }
    |	trait_list ',' fully_qualified_class_name		{  }
;

trait_adaptations:
        ';'
    |	'{' trait_adaptation_list '}'
;

trait_adaptation_list:
        /* empty */
    |	non_empty_trait_adaptation_list
;

non_empty_trait_adaptation_list:
        trait_adaptation_statement
    |	non_empty_trait_adaptation_list trait_adaptation_statement
;

trait_adaptation_statement:
        trait_precedence ';'
    |	trait_alias ';'
;

trait_precedence:
    trait_method_reference_fully_qualified T_INSTEADOF trait_reference_list	{  }
;

trait_reference_list:
        fully_qualified_class_name									{  }
    |	trait_reference_list ',' fully_qualified_class_name			{  }
;

trait_method_reference:
        T_STRING													{  }
    |	trait_method_reference_fully_qualified						{  }
;

trait_method_reference_fully_qualified:
    fully_qualified_class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING		{  }
;

trait_alias:
        trait_method_reference T_AS trait_modifiers T_STRING		{  }
    |	trait_method_reference T_AS member_modifier					{  }
;

trait_modifiers:
        /* empty */					{  } /* No change of methods visibility */
    |	member_modifier	{  } /* REM: Keep in mind, there are not only visibility modifiers */
;

method_body:
        ';' /* abstract method */		{  }
    |	'{' inner_statement_list '}'	{  }
;

variable_modifiers:
        non_empty_member_modifiers		{  }
    |	T_VAR							{  }
;

method_modifiers:
        /* empty */							{  }
    |	non_empty_member_modifiers			{  }
;

non_empty_member_modifiers:
        member_modifier						{  }
    |	non_empty_member_modifiers member_modifier	{  }
;

member_modifier:
        T_PUBLIC				{  }
    |	T_PROTECTED				{  }
    |	T_PRIVATE				{  }
    |	T_STATIC				{  }
    |	T_ABSTRACT				{  }
    |	T_FINAL					{  }
;

class_variable_declaration:
        class_variable_declaration ',' T_VARIABLE					{  }
    |	class_variable_declaration ',' T_VARIABLE '=' static_scalar	{  }
    |	T_VARIABLE						{  }
    |	T_VARIABLE '=' static_scalar	{  }
;

class_constant_declaration:
        class_constant_declaration ',' T_STRING '=' static_scalar	{  }
    |	T_CONST T_STRING '=' static_scalar	{  }
;

echo_expr_list:
        echo_expr_list ',' expr {  }
    |	expr					{  }
;


for_expr:
        /* empty */			{  }
    |	non_empty_for_expr	{  }
;

non_empty_for_expr:
        non_empty_for_expr ','	{  } expr {  }
    |	expr					{  }
;

chaining_method_or_property:
        chaining_method_or_property variable_property 	{  }
    |	variable_property 								{  }
;

chaining_dereference:
        chaining_dereference '[' dim_offset ']'	{  }
    |	'[' dim_offset ']'		{  }
;

chaining_instance_call:
        chaining_dereference 		{  } chaining_method_or_property {  }
    |	chaining_dereference 		{  }
    |	chaining_method_or_property {  }
;

instance_call:
        /* empty */ 		{  }
    |	{  }
        chaining_instance_call	{  }
;

new_expr:
        T_NEW class_name_reference {  } ctor_arguments {  }
;

expr_without_variable:
        T_LIST '(' {  } assignment_list ')' '=' expr {  }
    |	variable '=' expr		{  }
    |	variable '=' '&' variable {  }
    |	variable '=' '&' T_NEW class_name_reference {  } ctor_arguments {  }
    |	T_CLONE expr {  }
    |	variable T_PLUS_EQUAL expr 	{  }
    |	variable T_MINUS_EQUAL expr	{  }
    |	variable T_MUL_EQUAL expr		{  }
    |	variable T_POW_EQUAL expr		{  }
    |	variable T_DIV_EQUAL expr		{  }
    |	variable T_CONCAT_EQUAL expr	{  }
    |	variable T_MOD_EQUAL expr		{  }
    |	variable T_AND_EQUAL expr		{  }
    |	variable T_OR_EQUAL expr 		{  }
    |	variable T_XOR_EQUAL expr 		{  }
    |	variable T_SL_EQUAL expr	{  }
    |	variable T_SR_EQUAL expr	{  }
    |	rw_variable T_INC {  }
    |	T_INC rw_variable {  }
    |	rw_variable T_DEC {  }
    |	T_DEC rw_variable {  }
    |	expr T_BOOLEAN_OR {  } expr {  }
    |	expr T_BOOLEAN_AND {  } expr {  }
    |	expr T_LOGICAL_OR {  } expr {  }
    |	expr T_LOGICAL_AND {  } expr {  }
    |	expr T_LOGICAL_XOR expr {  }
    |	expr '|' expr	{  }
    |	expr '&' expr	{  }
    |	expr '^' expr	{  }
    |	expr '.' expr 	{  }
    |	expr '+' expr 	{  }
    |	expr '-' expr 	{  }
    |	expr '*' expr	{  }
    |	expr T_POW expr	{  }
    |	expr '/' expr	{  }
    |	expr '%' expr 	{  }
    | 	expr T_SL expr	{  }
    |	expr T_SR expr	{  }
    |	'+' expr %prec T_INC {  }
    |	'-' expr %prec T_INC {  }
    |	'!' expr {  }
    |	'~' expr {  }
    |	expr T_IS_IDENTICAL expr		{  }
    |	expr T_IS_NOT_IDENTICAL expr	{  }
    |	expr T_IS_EQUAL expr			{  }
    |	expr T_IS_NOT_EQUAL expr 		{  }
    |	expr '<' expr 					{  }
    |	expr T_IS_SMALLER_OR_EQUAL expr {  }
    |	expr '>' expr 					{  }
    |	expr T_IS_GREATER_OR_EQUAL expr {  }
    |	expr T_INSTANCEOF class_name_reference {  }
    |	parenthesis_expr 	{  }
    |	new_expr		{  }
    |	'(' new_expr ')' {  } instance_call {  }
    |	expr '?' {  }
        expr ':' {  }
        expr	 {  }
    |	expr '?' ':' {  }
        expr     {  }
    |	internal_functions_in_yacc {  }
    |	T_INT_CAST expr 	{  }
    |	T_DOUBLE_CAST expr 	{  }
    |	T_STRING_CAST expr	{  }
    |	T_ARRAY_CAST expr 	{  }
    |	T_OBJECT_CAST expr 	{  }
    |	T_BOOL_CAST expr	{  }
    |	T_UNSET_CAST expr	{  }
    |	T_EXIT exit_expr	{  }
    |	'@' {  } expr {  }
    |	scalar				{  }
    |	combined_scalar_offset {  }
    |	combined_scalar {  }
    |	'`' backticks_expr '`' {  }
    |	T_PRINT expr  {  }
    |	T_YIELD {  }
    |	function is_reference {  }
        '(' parameter_list ')' lexical_vars
        '{' inner_statement_list '}' {  }
    |	T_STATIC function is_reference {  }
        '(' parameter_list ')' lexical_vars
        '{' inner_statement_list '}' {  }
;

yield_expr:
        T_YIELD expr_without_variable {  }
    |	T_YIELD variable {  }
    |	T_YIELD expr T_DOUBLE_ARROW expr_without_variable {  }
    |	T_YIELD expr T_DOUBLE_ARROW variable {  }
;

combined_scalar_offset:
        combined_scalar '[' dim_offset ']' {  }
    |	combined_scalar_offset '[' dim_offset ']' {  }
    |	T_CONSTANT_ENCAPSED_STRING '[' dim_offset ']' {  }
    |	general_constant '[' dim_offset ']' {  }
;

combined_scalar:
        T_ARRAY '(' array_pair_list ')' {  }
    |	'[' array_pair_list ']' {  }
;

function:
    T_FUNCTION {  }
;

lexical_vars:
        /* empty */
    |	T_USE '(' lexical_var_list ')'
;

lexical_var_list:
        lexical_var_list ',' T_VARIABLE			{  }
    |	lexical_var_list ',' '&' T_VARIABLE		{  }
    |	T_VARIABLE								{  }
    |	'&' T_VARIABLE							{  }
;

function_call:
        namespace_name {  }
        function_call_parameter_list {  }
    |	T_NAMESPACE T_NS_SEPARATOR namespace_name {  }
        function_call_parameter_list {  }
    |	T_NS_SEPARATOR namespace_name {  }
        function_call_parameter_list {  }
    |	class_name T_PAAMAYIM_NEKUDOTAYIM variable_name {  }
        function_call_parameter_list {  }
    |	class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects {  }
        function_call_parameter_list {  }
    |	variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_name {  }
        function_call_parameter_list {  }
    |	variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects {  }
        function_call_parameter_list {  }
    |	variable_without_objects {  }
        function_call_parameter_list {  }
;

class_name:
        T_STATIC {  }
    |	namespace_name {  }
    |	T_NAMESPACE T_NS_SEPARATOR namespace_name {  }
    |	T_NS_SEPARATOR namespace_name {  }
;

fully_qualified_class_name:
        namespace_name {  }
    |	T_NAMESPACE T_NS_SEPARATOR namespace_name {  }
    |	T_NS_SEPARATOR namespace_name {  }
;



class_name_reference:
        class_name						{  }
    |	dynamic_class_name_reference	{  }
;


dynamic_class_name_reference:
        base_variable T_OBJECT_OPERATOR {  }
            object_property {  } dynamic_class_name_variable_properties
            {  }
    |	base_variable {  }
;


dynamic_class_name_variable_properties:
        dynamic_class_name_variable_properties dynamic_class_name_variable_property
    |	/* empty */
;


dynamic_class_name_variable_property:
        T_OBJECT_OPERATOR object_property {  }
;

exit_expr:
        /* empty */	{  }
    |	'(' ')'		{  }
    |	parenthesis_expr	{  }
;

backticks_expr:
        /* empty */	{  }
    |	T_ENCAPSED_AND_WHITESPACE	{  }
    |	encaps_list	{  }
;


ctor_arguments:
        /* empty */	{  }
    |	function_call_parameter_list 	{  }
;


common_scalar:
        T_LNUMBER 					{  }
    |	T_DNUMBER 					{  }
    |	T_CONSTANT_ENCAPSED_STRING	{  }
    |	T_LINE 						{  }
    |	T_FILE 						{  }
    |	T_DIR   					{  }
    |	T_TRAIT_C					{  }
    |	T_METHOD_C					{  }
    |	T_FUNC_C					{  }
    |	T_NS_C						{  }
    |	T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC {  }
    |	T_START_HEREDOC T_END_HEREDOC {  }
;

static_class_constant:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING {  }
;

static_scalar: /* compile-time evaluated scalars */
    static_scalar_value {  }
;

static_scalar_value:
        common_scalar	{  }
    |	static_class_name_scalar	{  }
    |	namespace_name 		{  }
    |	T_NAMESPACE T_NS_SEPARATOR namespace_name {  }
    |	T_NS_SEPARATOR namespace_name {  }
    |	T_ARRAY '(' static_array_pair_list ')' {  }
    |	'[' static_array_pair_list ']' {  }
    |	static_class_constant {  }
    |	T_CLASS_C			{  }
    |	static_operation {  }
;

static_operation:
        static_scalar_value '[' static_scalar_value ']' {  }
    |	static_scalar_value '+' static_scalar_value {  }
    |	static_scalar_value '-' static_scalar_value {  }
    |	static_scalar_value '*' static_scalar_value {  }
    |	static_scalar_value T_POW static_scalar_value {  }
    |	static_scalar_value '/' static_scalar_value {  }
    |	static_scalar_value '%' static_scalar_value {  }
    |	'!' static_scalar_value {  }
    |	'~' static_scalar_value {  }
    |	static_scalar_value '|' static_scalar_value {  }
    |	static_scalar_value '&' static_scalar_value {  }
    |	static_scalar_value '^' static_scalar_value {  }
    |	static_scalar_value T_SL static_scalar_value {  }
    |	static_scalar_value T_SR static_scalar_value {  }
    |	static_scalar_value '.' static_scalar_value {  }
    |	static_scalar_value T_LOGICAL_XOR static_scalar_value {  }
    |	static_scalar_value T_LOGICAL_AND static_scalar_value {  }
    |	static_scalar_value T_LOGICAL_OR static_scalar_value {  }
    |	static_scalar_value T_BOOLEAN_AND static_scalar_value {  }
    |	static_scalar_value T_BOOLEAN_OR static_scalar_value {  }
    |	static_scalar_value T_IS_IDENTICAL static_scalar_value {  }
    |	static_scalar_value T_IS_NOT_IDENTICAL static_scalar_value {  }
    |	static_scalar_value T_IS_EQUAL static_scalar_value {  }
    |	static_scalar_value T_IS_NOT_EQUAL static_scalar_value {  }
    |	static_scalar_value '<' static_scalar_value {  }
    |	static_scalar_value '>' static_scalar_value {  }
    |	static_scalar_value T_IS_SMALLER_OR_EQUAL static_scalar_value {  }
    |	static_scalar_value T_IS_GREATER_OR_EQUAL static_scalar_value {  }
    |	static_scalar_value '?' ':' static_scalar_value {  }
    |	static_scalar_value '?' static_scalar_value ':' static_scalar_value {  }
    |	'+' static_scalar_value {  }
    |	'-' static_scalar_value {  }
    |	'(' static_scalar_value ')' {  }
;

general_constant:
        class_constant {  }
    |	namespace_name	{  }
    |	T_NAMESPACE T_NS_SEPARATOR namespace_name {  }
    |	T_NS_SEPARATOR namespace_name {  }
;

scalar:
        T_STRING_VARNAME {  }
    |	general_constant {  }
    |	class_name_scalar {  }
    |	common_scalar {  }
    |	'"' encaps_list '"' {  }
    |	T_START_HEREDOC encaps_list T_END_HEREDOC {  }
    |	T_CLASS_C {  }
;

static_array_pair_list:
        /* empty */ {  }
    |	non_empty_static_array_pair_list possible_comma	{  }
;

possible_comma:
        /* empty */
    |	','
;

non_empty_static_array_pair_list:
        non_empty_static_array_pair_list ',' static_scalar_value T_DOUBLE_ARROW static_scalar_value {  }
    |	non_empty_static_array_pair_list ',' static_scalar_value {  }
    |	static_scalar_value T_DOUBLE_ARROW static_scalar_value {  }
    |	static_scalar_value {  }
;

expr:
        r_variable					{  }
    |	expr_without_variable		{  }
;

parenthesis_expr:
        '(' expr ')'		{  }
    |	'(' yield_expr ')'	{  }
;


r_variable:
    variable {  }
;


w_variable:
    variable	{  }
;

rw_variable:
    variable	{  }
;

variable:
        base_variable_with_function_calls T_OBJECT_OPERATOR {  }
            object_property { } method_or_not variable_properties
            {  }
    |	base_variable_with_function_calls {  }
;

variable_properties:
        variable_properties variable_property {  }
    |	/* empty */ {  }
;


variable_property:
        T_OBJECT_OPERATOR object_property {  } method_or_not {  }
;

array_method_dereference:
        array_method_dereference '[' dim_offset ']' {  }
    |	method '[' dim_offset ']' {  }
;

method:
        {  }
        function_call_parameter_list {  }
;

method_or_not:
        method						{  }
    |	array_method_dereference	{  }
    |	/* empty */ {  }
;

variable_without_objects:
        reference_variable {  }
    |	simple_indirect_reference reference_variable {  }
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects {  }
    |	variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects {  }

;

variable_class_name:
        reference_variable {  }
;

array_function_dereference:
        array_function_dereference '[' dim_offset ']' {  }
    |	function_call {  }
        '[' dim_offset ']' {  }
;

base_variable_with_function_calls:
        base_variable				{  }
    |	array_function_dereference	{  }
    |	function_call {  }
;


base_variable:
        reference_variable {  }
    |	simple_indirect_reference reference_variable {  }
    |	static_member {  }
;

reference_variable:
        reference_variable '[' dim_offset ']'	{  }
    |	reference_variable '{' expr '}'		{  }
    |	compound_variable			{  }
;


compound_variable:
        T_VARIABLE			{  }
    |	'$' '{' expr '}'	{  }
;

dim_offset:
        /* empty */		{  }
    |	expr			{  }
;


object_property:
        object_dim_list {  }
    |	variable_without_objects {  }
;

object_dim_list:
        object_dim_list '[' dim_offset ']'	{  }
    |	object_dim_list '{' expr '}'		{  }
    |	variable_name { }
;

variable_name:
        T_STRING		{  }
    |	'{' expr '}'	{  }
;

simple_indirect_reference:
        '$' {  }
    |	simple_indirect_reference '$' {  }
;

assignment_list:
        assignment_list ',' assignment_list_element
    |	assignment_list_element
;


assignment_list_element:
        variable								{  }
    |	T_LIST '(' {  } assignment_list ')'	{  }
    |	/* empty */							{  }
;


array_pair_list:
        /* empty */ {  }
    |	non_empty_array_pair_list possible_comma	{  }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' expr T_DOUBLE_ARROW expr	{  }
    |	non_empty_array_pair_list ',' expr			{  }
    |	expr T_DOUBLE_ARROW expr	{  }
    |	expr 				{  }
    |	non_empty_array_pair_list ',' expr T_DOUBLE_ARROW '&' w_variable {  }
    |	non_empty_array_pair_list ',' '&' w_variable {  }
    |	expr T_DOUBLE_ARROW '&' w_variable	{  }
    |	'&' w_variable 			{  }
;

encaps_list:
        encaps_list encaps_var { }
    |	encaps_list T_ENCAPSED_AND_WHITESPACE	{ }
    |	encaps_var { }
    |	T_ENCAPSED_AND_WHITESPACE encaps_var	{ }
;



encaps_var:
        T_VARIABLE {  }
    |	T_VARIABLE '[' {  } encaps_var_offset ']'	{  }
    |	T_VARIABLE T_OBJECT_OPERATOR T_STRING {  }
    |	T_DOLLAR_OPEN_CURLY_BRACES expr '}' {  }
    |	T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}' {  }
    |	T_CURLY_OPEN variable '}' {  }
;


encaps_var_offset:
        T_STRING		{  }
    |	T_NUM_STRING	{  }
    |	T_VARIABLE		{  }
;


internal_functions_in_yacc:
        T_ISSET '(' isset_variables ')' {  }
    |	T_EMPTY '(' variable ')'	{  }
    |	T_EMPTY '(' expr_without_variable ')' {  }
    |	T_INCLUDE expr 			{  }
    |	T_INCLUDE_ONCE expr 	{  }
    |	T_EVAL '(' expr ')' 	{  }
    |	T_REQUIRE expr			{  }
    |	T_REQUIRE_ONCE expr		{  }
;

isset_variables:
        isset_variable			{  }
    |	isset_variables ',' {  } isset_variable {  }
;

isset_variable:
        variable				{  }
    |	expr_without_variable	{  }
;

class_constant:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING {  }
    |	variable_class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING {  }
;

static_class_name_scalar:
    class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS {  }
;

class_name_scalar:
    class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS {  }
;

%%