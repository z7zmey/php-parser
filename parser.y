%{
package main

import (
  "bytes"
  "fmt"
  "os"
  "io"
)

type node struct {
  name string
  children []node
  attributes map[string]string
}

func (n node) String() string {
  buf := new(bytes.Buffer)
  n.print(buf, " ")
  return buf.String()
}

func (n node) print(out io.Writer, indent string) {
  if (len(n.attributes) > 0) {
    fmt.Fprintf(out, "\n%v%v %s", indent, n.name, n.attributes)
  } else {
    fmt.Fprintf(out, "\n%v%v", indent, n.name)
  }
  for _, nn := range n.children { 
    nn.print(out, indent + "  ") 
  }
}

func Node(name string) node { 
  return node{name: name, attributes: make(map[string]string)} 
}

func (n node) append(nn...node) node {
  n.children = append(n.children, nn...)
  return n 
}

func (n node) attribute(key string, value string) node {
  n.attributes[key] = value
  return n 
}

%}

%union{
    node node
    token string
    value string
}

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

%type <value> class_modifier
%type <value> is_reference
%type <value> is_variadic
%type <value> returns_ref
%type <value> reserved_non_modifiers
%type <value> semi_reserved

%type <node> top_statement namespace_name name statement function_declaration_statement
%type <node> class_declaration_statement trait_declaration_statement
%type <node> interface_declaration_statement interface_extends_list
%type <node> group_use_declaration inline_use_declarations inline_use_declaration
%type <node> mixed_group_use_declaration use_declaration unprefixed_use_declaration
%type <node> unprefixed_use_declarations const_decl inner_statement
%type <node> expr optional_expr while_statement for_statement foreach_variable
%type <node> foreach_statement declare_statement finally_statement unset_variable variable
%type <node> extends_from parameter optional_type argument expr_without_variable global_var
%type <node> static_var class_statement trait_adaptation trait_precedence trait_alias
%type <node> absolute_trait_method_reference trait_method_reference property echo_expr
%type <node> new_expr anonymous_class class_name class_name_reference simple_variable
%type <node> internal_functions_in_yacc
%type <node> exit_expr scalar backticks_expr lexical_var function_call member_name property_name
%type <node> variable_class_name dereferencable_scalar constant dereferencable
%type <node> callable_expr callable_variable static_member new_variable
%type <node> encaps_var encaps_var_offset isset_variables
%type <node> top_statement_list use_declarations const_list inner_statement_list if_stmt
%type <node> alt_if_stmt for_exprs switch_case_list global_var_list static_var_list
%type <node> echo_expr_list unset_variables catch_name_list catch_list parameter_list class_statement_list
%type <node> implements_list case_list if_stmt_without_else
%type <node> non_empty_parameter_list argument_list non_empty_argument_list property_list
%type <node> class_const_list class_const_decl name_list trait_adaptations method_body non_empty_for_exprs
%type <node> ctor_arguments alt_if_stmt_without_else trait_adaptation_list lexical_vars
%type <node> lexical_var_list encaps_list
%type <node> array_pair non_empty_array_pair_list array_pair_list possible_array_pair
%type <node> isset_variable type return_type type_expr
%type <node> identifier

%type <node> variable_modifiers
%type <node> method_modifiers non_empty_member_modifiers member_modifier
%type <node> class_modifiers use_type

%%

/////////////////////////////////////////////////////////////////////////

start:
    top_statement_list                                  { fmt.Println($1) }
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
        reserved_non_modifiers {$$=$1}
    | T_STATIC {$$=$1} | T_ABSTRACT {$$=$1} | T_FINAL {$$=$1} | T_PRIVATE {$$=$1} | T_PROTECTED {$$=$1} | T_PUBLIC {$$=$1}
;

identifier:
        T_STRING                                        { $$ = Node("identifier").attribute("value", $1) }
    |   semi_reserved                                   { $$ = Node("identifier").attribute("value", $1) }
;

top_statement_list:
        top_statement_list top_statement                { $$ = $1.append($2); }
    |   /* empty */                                     { $$ = Node("Statements") }
;

namespace_name:
        T_STRING                                        { $$ = Node("NamespaceParts").append(Node($1)); }
    |   namespace_name T_NS_SEPARATOR T_STRING          { $$ = $1.append(Node($3)); }
;

name:
      namespace_name                                    { $$ = Node("Name").append($1); }
    | T_NAMESPACE T_NS_SEPARATOR namespace_name         { $$ = Node("Name").append($3).attribute("Relative", "true"); }
    | T_NS_SEPARATOR namespace_name                     { $$ = Node("Name").append($2).attribute("FullyQualified", "true"); }
;

top_statement:
        statement                                       { $$ = $1; }
    |   function_declaration_statement                  { $$ = $1; }
    |   class_declaration_statement                     { $$ = $1; }
    |   trait_declaration_statement                     { $$ = $1; }
    |   interface_declaration_statement                 { $$ = $1; }
    |   T_HALT_COMPILER '(' ')' ';'                     { $$ = Node("THaltCompiler") }
    |   T_NAMESPACE namespace_name ';'                  { $$ = Node("Namespace").append($2); }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
                                                        { $$ = Node("Namespace").append($2).append($4) }
    |   T_NAMESPACE '{' top_statement_list '}'          { $$ = Node("Namespace").append($3) }
    |   T_USE mixed_group_use_declaration ';'           { $$ = $2; }
    |   T_USE use_type group_use_declaration ';'        { $$ = $3.append($2) }
    |   T_USE use_declarations ';'                      { $$ = $2; }
    |   T_USE use_type use_declarations ';'             { $$ = $3.append($2) }
    |   T_CONST const_list ';'                          { $$ = $2; }
;

use_type:
        T_FUNCTION                                      { $$ = Node("FuncUseType"); }
    |   T_CONST                                         { $$ = Node("ConstUseType"); }
;

group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
                                                        { $$ = Node("GroupUse").append($1).append($4) }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
                                                        { $$ = Node("GroupUse").append($2).append($5) }
;

mixed_group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
                                                        { $$ = Node("MixedGroupUse").append($1).append($4); }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
                                                        { $$ = Node("MixedGroupUse").append($2).append($5); }
;

possible_comma:
        /* empty */
    |   ','
;

inline_use_declarations:
        inline_use_declarations ',' inline_use_declaration
                                                        { $$ = $1.append($3) }
    |   inline_use_declaration                          { $$ = Node("UseList").append($1) }
;

unprefixed_use_declarations:
        unprefixed_use_declarations ',' unprefixed_use_declaration
                                                        { $$ = $1.append($3) }
    |   unprefixed_use_declaration                      { $$ = Node("UseList").append($1) }
;

use_declarations:
        use_declarations ',' use_declaration            { $$ = $1.append($3) }
    |   use_declaration                                 { $$ = Node("UseList").append($1) }
;

inline_use_declaration:
        unprefixed_use_declaration                      { $$ = $1; }
    |   use_type unprefixed_use_declaration             { $$ = $2.append($1) }
;

unprefixed_use_declaration:
        namespace_name                                  { $$ = Node("UseElem").append($1); }
    |   namespace_name T_AS T_STRING                    { $$ = Node("UseElem").append($1).append(Node("as").attribute("value", $3)); }
;

use_declaration:
        unprefixed_use_declaration                      { $$ = $1; }
    |   T_NS_SEPARATOR unprefixed_use_declaration       { $$ = $2; }
;

const_list:
        const_list ',' const_decl                       { $$ = $1.append($3) }
    |   const_decl                                      { $$ = Node("ConstList").append($1) }
;

inner_statement_list:
        inner_statement_list inner_statement            { $$ = $1.append($2); }
    |   /* empty */                                     { $$ = Node("stmt") }
;

inner_statement:
    statement                                           { $$ = $1; }
    |   function_declaration_statement                  { $$ = $1; }
    |   class_declaration_statement                     { $$ = $1; }
    |   trait_declaration_statement                     { $$ = $1; }
    |   interface_declaration_statement                 { $$ = $1; }
    |   T_HALT_COMPILER '(' ')' ';'                     { $$ = Node("THaltCompiler") }

statement:
    '{' inner_statement_list '}'                        { $$ = $2; }
    |   if_stmt                                         { $$ = $1; }
    |   alt_if_stmt                                     { $$ = $1; }
    |   T_WHILE '(' expr ')' while_statement
        {
            $$ = Node("While").
                append(Node("expr").append($3)).
                append(Node("stmt").append($5));
        }
    |   T_DO statement T_WHILE '(' expr ')' ';'
        {
            $$ = Node("DoWhile").
                append(Node("expr").append($5)).
                append(Node("stmt").append($2));
        }
    |   T_FOR '(' for_exprs ';' for_exprs ';' for_exprs ')' for_statement
            {
                $$ = Node("For").
                    append(Node("expr1").append($3)).
                    append(Node("expr2").append($5)).
                    append(Node("expr3").append($7)).
                    append(Node("stmt").append($9))
            }
    |   T_SWITCH '(' expr ')' switch_case_list          { $$ = Node("Switch").append(Node("expr").append($3)).append($5); }
    |   T_BREAK optional_expr ';'                       { $$ = Node("Break").append($2) }
    |   T_CONTINUE optional_expr ';'                    { $$ = Node("Continue").append($2) }
    |   T_RETURN optional_expr ';'                      { $$ = Node("Return").append($2) }
    |   T_GLOBAL global_var_list ';'                    { $$ = $2; }
    |   T_STATIC static_var_list ';'                    { $$ = $2; }
    |   T_ECHO echo_expr_list ';'                       { $$ = $2; }
    |   T_INLINE_HTML                                   { $$ = Node("Echo").append(Node("InlineHtml").attribute("value", $1)) }
    |   expr ';'                                        { $$ = $1; }
    |   T_UNSET '(' unset_variables possible_comma ')' ';' 
                                                        { $$ = Node("Unset").append($3); }
    |   T_FOREACH '(' expr T_AS foreach_variable ')' foreach_statement
            {
                $$ = Node("Foreach").
                    append(Node("expr").append($3)).
                    append(Node("ForeachVariable").append($5)).
                    append($7);
            }
    |   T_FOREACH '(' expr T_AS foreach_variable T_DOUBLE_ARROW foreach_variable ')' foreach_statement
            {
                $$ = Node("Foreach").
                    append(Node("expr").append($3)).
                    append(Node("ForeachKey").append($5)).
                    append(Node("ForeachVariable").append($7)).
                    append($9);
            }
    |   T_DECLARE '(' const_list ')' declare_statement  { $$ =  Node("Declare").append($3).append($5) }
    |   ';' /* empty statement */                       { $$ = Node(""); }
    |   T_TRY '{' inner_statement_list '}' catch_list finally_statement
            {
                $$ = Node("Try").
                    append($3).
                    append($5).
                    append($6);
            }
    |   T_THROW expr ';'                                { $$ = Node("Throw").append($2) }
    |   T_GOTO T_STRING ';'                             { $$ = Node("GoTo").attribute("Label", $2) }
    |   T_STRING ':'                                    { $$ = Node("Label").attribute("name", $1) }

catch_list:
        /* empty */                                     { $$ = Node("CatchList") }
    |   catch_list T_CATCH '(' catch_name_list T_VARIABLE ')' '{' inner_statement_list '}'
                                                        { $$ = $1.append($4).append(Node("Variable").attribute("name", $5)).append($8) }
;
catch_name_list:
        name                                            { $$ = Node("CatchNameList").append($1) }
    |   catch_name_list '|' name                        { $$ = $1.append($3) }
;

finally_statement:
        /* empty */                                     { $$ = Node(""); }
    |   T_FINALLY '{' inner_statement_list '}'          { $$ = Node("Finnaly").append($3) }
;

unset_variables:
        unset_variable                                  { $$ = Node("UnsetVariablesList").append($1) }
    |   unset_variables ',' unset_variable              { $$ = $1.append($3) }
;

unset_variable:
    variable                                            { $$ = $1 }
;

function_declaration_statement:
    T_FUNCTION returns_ref T_STRING '(' parameter_list ')' return_type '{' inner_statement_list '}'
        {
            $$ = Node("Function").
                attribute("name", $3).
                attribute("returns_ref", $2).
                append($5).
                append($7).
                append($9);
        }
;

is_reference:
        /* empty */                                     { $$ = "false"; }
    |   '&'                                             { $$ = "true"; }
;

is_variadic:
        /* empty */                                     { $$ = "false"; }
    |   T_ELLIPSIS                                      { $$ = "true"; }
;

class_declaration_statement:
        class_modifiers T_CLASS T_STRING extends_from implements_list '{' class_statement_list '}'
            {
                $$ = Node("Class").
                    attribute("name", $3).
                    append($1).
                    append(Node("Extends").append($4)).
                    append(Node("Implements").append($5)).
                    append($7);
            }
    |   T_CLASS T_STRING extends_from implements_list '{' class_statement_list '}'
            {
                $$ = Node("Class").
                    attribute("name", $2).
                    append(Node("Extends").append($3)).
                    append(Node("Implements").append($4)).
                    append($6);
            }
;

class_modifiers:
        class_modifier                                  { $$ = Node("Class").attribute($1, "true") }
    |   class_modifiers class_modifier                  { $$ = $1.attribute($2, "true") }
;

class_modifier:
        T_ABSTRACT                                      { $$ = "abstract" }
    |   T_FINAL                                         { $$ = "final" }
;

trait_declaration_statement:
    T_TRAIT T_STRING '{' class_statement_list '}'       { $$ = Node("Trait").attribute("name", $2).append($4) }
;

interface_declaration_statement:
    T_INTERFACE T_STRING interface_extends_list '{' class_statement_list '}'
        {
            $$ = Node("Interface").
                attribute("name", $2).
                append(Node("Extends").append($3)).
                append($5);
        }
;

extends_from:
        /* empty */                                     { $$ = Node(""); }
    |   T_EXTENDS name                                  { $$ = $2; }
;

interface_extends_list:
        /* empty */                                     { $$ = Node("") }
    |   T_EXTENDS name_list                             { $$ = $2; }
;

implements_list:
        /* empty */                                     { $$ = Node(""); }
    |   T_IMPLEMENTS name_list                          { $$ = $2; }
;

foreach_variable:
        variable                                        { $$ = $1; }
    |   '&' variable                                    { $$ = Node("Ref").append($2); }
    |   T_LIST '(' array_pair_list ')'                  { $$ = Node("List").append($3) }
    |   '[' array_pair_list ']'                         { $$ = Node("ShortList").append($2) }
;

for_statement:
        statement                                       { $$ = $1; }
    |    ':' inner_statement_list T_ENDFOR ';'          { $$ = $2; }
;

foreach_statement:
        statement                                       { $$ = $1; }
    |   ':' inner_statement_list T_ENDFOREACH ';'       { $$ = $2; }
;

declare_statement:
        statement                                       { $$ = $1; }
    |   ':' inner_statement_list T_ENDDECLARE ';'       { $$ = $2; }
;

switch_case_list:
        '{' case_list '}'                               { $$ = $2; }
    |   '{' ';' case_list '}'                           { $$ = $3; }
    |   ':' case_list T_ENDSWITCH ';'                   { $$ = $2; }
    |   ':' ';' case_list T_ENDSWITCH ';'               { $$ = $3; }
;

case_list:
        /* empty */                                     { $$ = Node("CaseList") }
    |   case_list T_CASE expr case_separator inner_statement_list
            {
                $$ = $1.append(Node("Case").append(Node("expr").append($3)).append($5))
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                $$ = $1.append(Node("Default").append($4))
            }
;

case_separator:
        ':'
    |   ';'
;

while_statement:
        statement                                       { $$ = $1; }
    |   ':' inner_statement_list T_ENDWHILE ';'         { $$ = $2; }
;

if_stmt_without_else:
        T_IF '(' expr ')' statement
            {
                $$ = Node("If").append(Node("expr").append($3)).append(Node("stmt").append($5))
            }
    |   if_stmt_without_else T_ELSEIF '(' expr ')' statement
            { 
                $$ = $1.append(Node("ElseIf").append(Node("expr").append($4)).append(Node("stmt").append($6)))
            }
;

if_stmt:
        if_stmt_without_else %prec T_NOELSE             { $$ = $1; }
    |   if_stmt_without_else T_ELSE statement
            {
                $$ = $1.append(Node("Else").append(Node("stmt").append($3)))
            }
;

alt_if_stmt_without_else:
        T_IF '(' expr ')' ':' inner_statement_list
            { 
                $$ = Node("AltIf").append(Node("expr").append($3)).append(Node("stmt").append($6))
            }
    |   alt_if_stmt_without_else T_ELSEIF '(' expr ')' ':' inner_statement_list
            {
                $$ = $1.append(Node("AltElseIf").append(Node("expr").append($4)).append(Node("stmt").append($7)))
            }
;

alt_if_stmt:
        alt_if_stmt_without_else T_ENDIF ';'            { $$ = $1; }
    |   alt_if_stmt_without_else T_ELSE ':' inner_statement_list T_ENDIF ';'
            {
                $$ = $1.append(Node("AltElse").append(Node("stmt").append($4)))
            }
;

parameter_list:
        non_empty_parameter_list                        { $$ = $1; }
    |   /* empty */                                     { $$ = Node("Parameter list"); }
;

non_empty_parameter_list:
        parameter                                       { $$ = Node("Parameter list").append($1) }
    |   non_empty_parameter_list ',' parameter          { $$ = $1.append($3); }
;

parameter:
        optional_type is_reference is_variadic T_VARIABLE
            {
                $$ = Node("Parameter").
                    append($1).
                    attribute("is_reference", $2).
                    attribute("is_variadic", $3).
                    attribute("var", $4);
            }
    |   optional_type is_reference is_variadic T_VARIABLE '=' expr
            {
                $$ = Node("Parameter").
                    append($1).
                    attribute("is_reference", $2).
                    attribute("is_variadic", $3).
                    attribute("var", $4).
                    append($6);
            }
;

optional_type:
        /* empty */                                     { $$ = Node("No type") }
    |   type_expr                                       { $$ = $1; }
;

type_expr:
        type                                            { $$ = $1; }
    |   '?' type                                        { $$ = $2; $$.attribute("nullable", "true") }
;

type:
        T_ARRAY                                         { $$ = Node("array type"); }
    |   T_CALLABLE                                      { $$ = Node("callable type"); }
    |   name                                            { $$ = $1; }
;

return_type:
        /* empty */                                     { $$ = Node("No return type"); }
    |   ':' type_expr                                   { $$ = $2; }
;

argument_list:
        '(' ')'                                         { $$ = Node("ArgumentList") }
    |   '(' non_empty_argument_list possible_comma ')'  { $$ = $2; }
;

non_empty_argument_list:
        argument                                        { $$ = Node("ArgumentList").append($1) }
    |   non_empty_argument_list ',' argument            { $$ = $1.append($3) }
;

argument:
        expr                                            { $$ = $1; }
    |   T_ELLIPSIS expr                                 { $$ = Node("Unpack").append($2) }
;

global_var_list:
        global_var_list ',' global_var                  { $$ = $1.append($3); }
    |   global_var                                      { $$ = Node("GlobalVarList").append($1); }
;

global_var:
    simple_variable                                     { $$ = $1 }
;

static_var_list:
        static_var_list ',' static_var                  { $$ = $1.append($3); }
    |   static_var                                      { $$ = Node("StaticVarList").append($1); }
;

static_var:
        T_VARIABLE                                      { $$ = Node("StaticVariable").attribute("Name", $1); }
    |   T_VARIABLE '=' expr                             { $$ = Node("StaticVariable").attribute("Name", $1).append(Node("expr").append($3)); }
;

class_statement_list:
        class_statement_list class_statement            { $$ = $1.append($2) }
    |   /* empty */                                     { $$ = Node("Stmt") }
;

class_statement:
        variable_modifiers property_list ';'            { $$ = $2.append($1) }
    |   method_modifiers T_CONST class_const_list ';'   { $$ = $3.append($1); }
    |   T_USE name_list trait_adaptations               { $$ = Node("Use").append($2).append($3); }
    |   method_modifiers T_FUNCTION returns_ref identifier '(' parameter_list ')'
        return_type method_body
            {
                $$ = Node("Function").
                    append($1).
                    append(Node("name").append($4)).
                    attribute("returns_ref", $3).
                    append($6).
                    append($8).
                    append($9);
            }
;

name_list:
        name                                            { $$ = Node("NameList").append($1) }
    |   name_list ',' name                              { $$ = $1.append($3) }
;

trait_adaptations:
        ';'                                             { $$ = Node(""); }
    |   '{' '}'                                         { $$ = Node(""); }
    |   '{' trait_adaptation_list '}'                   { $$ = $2; }
;

trait_adaptation_list:
        trait_adaptation                                { $$ = Node("TraitAdaptionList").append($1) }
    |   trait_adaptation_list trait_adaptation          { $$ = $1.append($2) }
;

trait_adaptation:
        trait_precedence ';'                            { $$ = $1; }
    |   trait_alias ';'                                 { $$ = $1; }
;

trait_precedence:
    absolute_trait_method_reference T_INSTEADOF name_list
                                                        { $$ = Node("TraitPrecedence").append($1).append($3) }
;

trait_alias:
        trait_method_reference T_AS T_STRING            { $$ = $1.append(Node("as").attribute("value", $3)); }
    |   trait_method_reference T_AS reserved_non_modifiers
                                                        { $$ = $1.append(Node("as").append(Node("reservedNonModifiers")));  }
    |   trait_method_reference T_AS member_modifier identifier
                                                        { $$ = $1.append($3).append($4); }
    |   trait_method_reference T_AS member_modifier     { $$ = $1.append($3); }
;

trait_method_reference:
        identifier                                      { $$ = Node("TraitMethodRef").append($1); }
    |   absolute_trait_method_reference                 { $$ = $1; }
;

absolute_trait_method_reference:
    name T_PAAMAYIM_NEKUDOTAYIM identifier              { $$ = Node("TraitMethodRef").append($1).append($3) }
;

method_body:
        ';' /* abstract method */                       { $$ = Node(""); }
    |   '{' inner_statement_list '}'                    { $$ = $2; }
;

variable_modifiers:
        non_empty_member_modifiers                      { $$ = $1; }
    |   T_VAR                                           { $$ = Node("VarMemberModifier") }
;

method_modifiers:
        /* empty */                                     { $$ = Node("PublicMemberModifier"); }
    |   non_empty_member_modifiers                      { $$ = $1; }
;

non_empty_member_modifiers:
        member_modifier	                                { $$ = $1; }
    |   non_empty_member_modifiers member_modifier      { $$ = $1.append($2) }
;

member_modifier:
        T_PUBLIC                                        { $$ = Node("PublicMemberModifier"); }
    |   T_PROTECTED                                     { $$ = Node("ProtectedMemberModifier"); }
    |   T_PRIVATE                                       { $$ = Node("PrivateMemberModifier"); }
    |   T_STATIC                                        { $$ = Node("StaticMemberModifier"); }
    |   T_ABSTRACT                                      { $$ = Node("AbstractMemberModifier"); }
    |   T_FINAL                                         { $$ = Node("FinalMemberModifier"); }
;

property_list:
        property_list ',' property                      { $$ = $1.append($3) }
    |   property                                        { $$ = Node("PropertyList").append($1) }
;

property:
        T_VARIABLE                                      { $$ = Node("Property").attribute("name", $1) }
    |   T_VARIABLE '=' expr                             { $$ = Node("Property").attribute("name", $1).append(Node("Default").append($3)) }
;

class_const_list:
        class_const_list ',' class_const_decl           { $$ = $1.append($3) }
    |   class_const_decl                                { $$ = Node("ConstList").append($1) }
;

class_const_decl:
    identifier '=' expr                                 { $$ = Node("Const").append($3) }
;

const_decl:
    T_STRING '=' expr                                   { $$ = Node("Const").attribute("name", $1).append($3) }
;

echo_expr_list:
        echo_expr_list ',' echo_expr                    { $$ = $1.append($3) }
    |   echo_expr                                       { $$ = Node("EchoList").append($1) }
;

echo_expr:
    expr                                                { $$ = Node("Echo").append($1) }
;

for_exprs:
        /* empty */                                     { $$ = Node(""); }
    |   non_empty_for_exprs                             { $$ = $1; }
;
non_empty_for_exprs:
        non_empty_for_exprs ',' expr                    { $$ = $1.append($3) }
    |   expr                                            { $$ = Node("ExpressionList").append($1) }
;

anonymous_class:
    T_CLASS ctor_arguments extends_from implements_list '{' class_statement_list '}'
        {
            $$ = Node("AnonymousClass").
                attribute("name", $1).
                append($2).
                append($3).
                append($4).
                append($6);
        }
;

new_expr:
        T_NEW class_name_reference ctor_arguments       { $$ = Node("New").append($2).append($3) }
    |   T_NEW anonymous_class                           { $$ = Node("New").append($2) }
;

expr_without_variable:
        T_LIST '(' array_pair_list ')' '=' expr         { $$ = Node("Assign").append($3).append($6); }
    |   '[' array_pair_list ']' '=' expr                { $$ = Node("Assign").append($2).append($5); }
    |   variable '=' expr                               { $$ = Node("Assign").append($1).append($3); }
    |   variable '=' '&' expr                           { $$ = Node("AssignRef").append($1).append($4); }
    |   T_CLONE expr                                    { $$ = Node("Clone").append($2); }
    |   variable T_PLUS_EQUAL expr                      { $$ = Node("AssignAdd").append($1).append($3); }
    |   variable T_MINUS_EQUAL expr                     { $$ = Node("AssignSub").append($1).append($3); }
    |   variable T_MUL_EQUAL expr                       { $$ = Node("AssignMul").append($1).append($3); }
    |   variable T_POW_EQUAL expr                       { $$ = Node("AssignPow").append($1).append($3); }
    |   variable T_DIV_EQUAL expr                       { $$ = Node("AssignDiv").append($1).append($3); }
    |   variable T_CONCAT_EQUAL expr                    { $$ = Node("AssignConcat").append($1).append($3); }
    |   variable T_MOD_EQUAL expr                       { $$ = Node("AssignMod").append($1).append($3); }
    |   variable T_AND_EQUAL expr                       { $$ = Node("AssignAnd").append($1).append($3); }
    |   variable T_OR_EQUAL expr                        { $$ = Node("AssignOr").append($1).append($3); }
    |   variable T_XOR_EQUAL expr                       { $$ = Node("AssignXor").append($1).append($3); }
    |   variable T_SL_EQUAL expr                        { $$ = Node("AssignShiftLeft").append($1).append($3); }
    |   variable T_SR_EQUAL expr                        { $$ = Node("AssignShiftRight").append($1).append($3); }
    |   variable T_INC                                  { $$ = Node("PostIncrement").append($1) }
    |   T_INC variable                                  { $$ = Node("PreIncrement").append($2) }
    |   variable T_DEC                                  { $$ = Node("PostDecrement").append($1) }
    |   T_DEC variable                                  { $$ = Node("PreDecrement").append($2) }
    |   expr T_BOOLEAN_OR expr                          { $$ = Node("Or").append($1).append($3) }
    |   expr T_BOOLEAN_AND expr                         { $$ = Node("And").append($1).append($3) }
    |   expr T_LOGICAL_OR expr                          { $$ = Node("Or").append($1).append($3) }
    |   expr T_LOGICAL_AND expr                         { $$ = Node("And").append($1).append($3) }
    |   expr T_LOGICAL_XOR expr                         { $$ = Node("Xor").append($1).append($3) }
    |   expr '|' expr                                   { $$ = Node("BitwiseOr").append($1).append($3) }
    |   expr '&' expr                                   { $$ = Node("BitwiseAnd").append($1).append($3) }
    |   expr '^' expr                                   { $$ = Node("BitwiseXor").append($1).append($3) }
    |   expr '.' expr                                   { $$ = Node("Concat").append($1).append($3) }
    |   expr '+' expr                                   { $$ = Node("Add").append($1).append($3) }
    |   expr '-' expr                                   { $$ = Node("Sub").append($1).append($3) }
    |   expr '*' expr                                   { $$ = Node("Mul").append($1).append($3) }
    |   expr T_POW expr                                 { $$ = Node("Pow").append($1).append($3) }
    |   expr '/' expr                                   { $$ = Node("Div").append($1).append($3) }
    |   expr '%' expr                                   { $$ = Node("Mod").append($1).append($3) }
    |   expr T_SL expr                                  { $$ = Node("ShiftLeft").append($1).append($3) }
    |   expr T_SR expr                                  { $$ = Node("ShiftRight").append($1).append($3) }
    |   '+' expr %prec T_INC                            { $$ = Node("UnaryPlus").append($2) }
    |   '-' expr %prec T_INC                            { $$ = Node("UnaryMinus").append($2) }
    |   '!' expr                                        { $$ = Node("BooleanNot").append($2) }
    |   '~' expr                                        { $$ = Node("BitwiseNot").append($2) }
    |   expr T_IS_IDENTICAL expr                        { $$ = Node("Identical").append($1).append($3) }
    |   expr T_IS_NOT_IDENTICAL expr                    { $$ = Node("NotIdentical").append($1).append($3) }
    |   expr T_IS_EQUAL expr                            { $$ = Node("Equal").append($1).append($3) }
    |   expr T_IS_NOT_EQUAL expr                        { $$ = Node("NotEqual").append($1).append($3) }
    |   expr '<' expr                                   { $$ = Node("Smaller").append($1).append($3) }
    |   expr T_IS_SMALLER_OR_EQUAL expr                 { $$ = Node("SmallerOrEqual").append($1).append($3) }
    |   expr '>' expr                                   { $$ = Node("Greater").append($1).append($3) }
    |   expr T_IS_GREATER_OR_EQUAL expr                 { $$ = Node("GreaterOrEqual").append($1).append($3) }
    |   expr T_SPACESHIP expr                           { $$ = Node("Spaceship").append($1).append($3); }
    |   expr T_INSTANCEOF class_name_reference          { $$ = Node("InstanceOf").append($1).append($3) }
    |   '(' expr ')'                                    { $$ = $2; }
    |   new_expr                                        { $$ = $1; }
    |   expr '?' expr ':' expr                          { $$ = Node("Ternary").append($1).append($3).append($5); }
    |   expr '?' ':' expr                               { $$ = Node("Ternary").append($1).append($4); }
    |   expr T_COALESCE expr                            { $$ = Node("Coalesce").append($1).append($3); }
    |   internal_functions_in_yacc                      { $$ = $1}
    |   T_INT_CAST expr                                 { $$ = Node("CastInt").append($2); }
    |   T_DOUBLE_CAST expr                              { $$ = Node("CastDouble").append($2); }
    |   T_STRING_CAST expr                              { $$ = Node("CastString").append($2); }
    |   T_ARRAY_CAST expr                               { $$ = Node("CastArray").append($2); }
    |   T_OBJECT_CAST expr                              { $$ = Node("CastObject").append($2); }
    |   T_BOOL_CAST expr                                { $$ = Node("CastBool").append($2); }
    |   T_UNSET_CAST expr                               { $$ = Node("CastUnset").append($2); }
    |   T_EXIT exit_expr                                { $$ = Node("Exit").append($2); }
    |   '@' expr                                        { $$ = Node("Silence").append($2); }
    |   scalar                                          { $$ = $1; }
    |   '`' backticks_expr '`'                          { $$ = Node("ShellExec").append($2) }
    |   T_PRINT expr                                    { $$ = Node("Print").append($2); }
    |   T_YIELD                                         { $$ = Node("Yield"); }
    |   T_YIELD expr                                    { $$ = Node("Yield").append($2); }
    |   T_YIELD expr T_DOUBLE_ARROW expr                { $$ = Node("Yield").append($2).append($4); }
    |   T_YIELD_FROM expr                               { $$ = Node("YieldFrom").append($2); }
    |   T_FUNCTION returns_ref '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                $$ = Node("Closure").
                    attribute("returns_ref", $2).
                    append($4).
                    append($6).
                    append($7).
                    append($9);
            }
    |   T_STATIC T_FUNCTION returns_ref '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                $$ = Node("StaticClosure").
                    attribute("returns_ref", $3).
                    append($5).
                    append($7).
                    append($8).
                    append($10);
            }
;

returns_ref:
        /* empty */                                     { $$ = "false"; }
    |   '&'                                             { $$ = "true"; }
;

lexical_vars:
        /* empty */                                     { $$ = Node("") }
    |   T_USE '(' lexical_var_list ')'                  { $$ = $3; }
;

lexical_var_list:
        lexical_var_list ',' lexical_var                { $$ = $1.append($3) }
    |   lexical_var                                     { $$ = Node("ClosureUses").append($1) }
;

lexical_var:
        T_VARIABLE                                      { $$ = Node("Variable").attribute("value", $1) }
    |   '&' T_VARIABLE                                  { $$ = Node("Variable").attribute("value", $2).attribute("ref", "true") }
;

function_call:
        name argument_list                              { $$ = Node("FunctionCall").append($1).append($2) }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
                                                        { $$ = Node("StaticCall").append($1).append($3).append($4) }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
                                                        { $$ = Node("StaticCall").append($1).append($3).append($4) }
    |   callable_expr argument_list                     { $$ = Node("Call").append($1).append($2); }
;

class_name:
        T_STATIC                                        { $$ = Node("Static") }
    |   name                                            { $$ = $1; }
;

class_name_reference:
        class_name                                      { $$ = $1; }
    |   new_variable                                    { $$ = $1; }
;

exit_expr:
        /* empty */                                     { $$ = Node("") }
    |   '(' optional_expr ')'                           { $$ = $2; }
;

backticks_expr:
        /* empty */                                     { $$ = Node("EmptyBackticks") }
    |   T_ENCAPSED_AND_WHITESPACE                       { $$ = Node("String").attribute("value", $1) }
    |   encaps_list                                     { $$ = $1; }
;

ctor_arguments:
        /* empty */	                                    { $$ = Node("ArgumentList") }
    |   argument_list                                   { $$ = $1; }
;

dereferencable_scalar:
        T_ARRAY '(' array_pair_list ')'                 { $$ = $3; }
    |   '[' array_pair_list ']'                         { $$ = $2; }
    |   T_CONSTANT_ENCAPSED_STRING                      { $$ = Node("String").attribute("value", $1) }
;

scalar:
        T_LNUMBER                                       { $$ = Node("Scalar").append(Node("Lnumber").attribute("value", $1)) }
    |   T_DNUMBER                                       { $$ = Node("Scalar").append(Node("Dnumber").attribute("value", $1)) }
    |   T_LINE                                          { $$ = Node("Scalar").append(Node("__LINE__")) }
    |   T_FILE                                          { $$ = Node("Scalar").append(Node("__FILE__")) }
    |   T_DIR                                           { $$ = Node("Scalar").append(Node("__DIR__")) }
    |   T_TRAIT_C                                       { $$ = Node("Scalar").append(Node("__TRAIT__")) }
    |   T_METHOD_C                                      { $$ = Node("Scalar").append(Node("__METHOD__")); }
    |   T_FUNC_C                                        { $$ = Node("Scalar").append(Node("__FUNCTION__")); }
    |   T_NS_C                                          { $$ = Node("Scalar").append(Node("__NAMESPACE__")); }
    |   T_CLASS_C                                       { $$ = Node("Scalar").append(Node("__CLASS__")); }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC 
                                                        { $$ = Node("Scalar").append(Node("Heredoc").attribute("value", $2)) }
    |   T_START_HEREDOC T_END_HEREDOC
                                                        { $$ = Node("Scalar").append(Node("Heredoc")) }
    |   '"' encaps_list '"'                             { $$ = $2; }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC       { $$ = $2; }
    |   dereferencable_scalar                           { $$ = $1; }
    |   constant                                        { $$ = $1; }
;

constant:
        name                                            { $$ = Node("Const").append($1) }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM identifier    { $$ = Node("Const").append($1).append($3) }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM identifier
                                                        { $$ = Node("Const").append($1).append($3) }
;

expr:
        variable                                        { $$ = $1; }
    |   expr_without_variable                           { $$ = $1; }
;

optional_expr:
        /* empty */                                     { $$ = Node("") }
    |   expr                                            { $$ = $1; }
;

variable_class_name:
    dereferencable                                      { $$ = $1; }
;

dereferencable:
        variable                                        { $$ = $1; }
    |   '(' expr ')'                                    { $$ = $2; }
    |   dereferencable_scalar                           { $$ = $1; }
;

callable_expr:
        callable_variable                               { $$ = $1; }
    |   '(' expr ')'                                    { $$ = $2; }
    |   dereferencable_scalar                           { $$ = $1; }
;

callable_variable:
    simple_variable                                     { $$ = $1; }
    |   dereferencable '[' optional_expr ']'            { $$ = Node("Dim").append($1).append($3)}
    |   constant '[' optional_expr ']'                  { $$ = Node("Dim").append($1).append($3)}
    |   dereferencable '{' expr '}'                     { $$ = Node("Dim").append($1).append($3)}
    |   dereferencable T_OBJECT_OPERATOR property_name argument_list
                                                        { $$ = Node("MethodCall").append($1).append($3).append($4)}
    |   function_call                                   { $$ = $1; }
;

variable:
        callable_variable                               { $$ = $1; }
    |   static_member                                   { $$ = $1; }
    |   dereferencable T_OBJECT_OPERATOR property_name  { $$ = Node("Property").append($1).append($3) }
;

simple_variable:
        T_VARIABLE                                      { $$ = Node("Variable").attribute("name", $1); }
    |   '$' '{' expr '}'                                { $$ = Node("Variable").append($3); }
    |   '$' simple_variable                             { $$ = Node("Variable").append($2); }
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
                                                        { $$ = Node("StaticProp").append($1).append($3) }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
                                                        { $$ = Node("StaticProp").append($1).append($3) }
;

new_variable:
        simple_variable                                 { $$ = $1 }
    |   new_variable '[' optional_expr ']'              { $$ = Node("Dim").append($1).append($3) }
    |   new_variable '{' expr '}'                       { $$ = Node("Dim").append($1).append($3) }
    |   new_variable T_OBJECT_OPERATOR property_name    { $$ = Node("Property").append($1).append($3) }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
                                                        { $$ = Node("StaticProperty").append($1).append($3) }
    |   new_variable T_PAAMAYIM_NEKUDOTAYIM simple_variable
                                                        { $$ = Node("StaticProperty").append($1).append($3) }
;

member_name:
        identifier                                      { $$ = $1; }
    |   '{' expr '}'                                    { $$ = $2; }
    |   simple_variable                                 { $$ = $1 }
;

property_name:
        T_STRING                                        { $$ = Node("PropertyName").attribute("value", $1) }
    |   '{' expr '}'                                    { $$ = $2; }
    |   simple_variable                                 { $$ = $1 }
;

array_pair_list:
    non_empty_array_pair_list                           { /* TODO: allow single trailing comma */ $$ = $1 }
;

possible_array_pair:
        /* empty */                                     { $$ = Node(""); }
    |   array_pair                                      { $$ = $1; }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' possible_array_pair
                                                        { $$ = $1.append($3) }
    |   possible_array_pair                             { $$ = Node("ArrayPairList").append($1) }
;

array_pair:
        expr T_DOUBLE_ARROW expr                        { $$ = Node("ArrayElement").append($1).append($3) }
    |   expr                                            { $$ = Node("ArrayElement").append($1) }
    |   expr T_DOUBLE_ARROW '&' variable                { $$ = Node("ArrayElement").append($1).append(Node("Ref").append($4)) }
    |   '&' variable                                    { $$ = Node("ArrayElement").append(Node("Ref").append($2)) }
    |   expr T_DOUBLE_ARROW T_LIST '(' array_pair_list ')'
            { 
                $$ = Node("ArrayElement").
                    append($1).
                    append(Node("ArrayList").append($5))
            }
    |   T_LIST '(' array_pair_list ')'
            {
                $$ = Node("ArrayElement").
                    append(Node("ArrayList").append($3))
            }
;

encaps_list:
        encaps_list encaps_var                          { $$ = $1.append($2) }
    |   encaps_list T_ENCAPSED_AND_WHITESPACE           { $$ = $1.append(Node("String").attribute("value", $2)) }
    |   encaps_var                                      { $$ = Node("EncapsList").append($1) }
    |   T_ENCAPSED_AND_WHITESPACE encaps_var            { $$ = Node("EncapsList").append(Node("String").attribute("value", $1)).append($2) }
;

encaps_var:
        T_VARIABLE                                      { $$ = Node("Variable").attribute("value", $1) }
    |   T_VARIABLE '[' encaps_var_offset ']'            { $$ = Node("Variable").attribute("value", $1).append(Node("offset").append($3)) }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING           { $$ = Node("Variable").attribute("value", $1).append(Node("property").attribute("value", $3)) }
    |   T_DOLLAR_OPEN_CURLY_BRACES expr '}'             { $$ = Node("Variable").append(Node("expr").append($2)) }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}' { $$ = Node("Variable").attribute("value", $2) }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
                                                        { $$ = Node("Variable").attribute("value", $2).append(Node("offset").append($4)) }
    |   T_CURLY_OPEN variable '}'                       { $$ = $2; }
;
encaps_var_offset:
        T_STRING                                        { $$ = Node("OffsetString").attribute("value", $1) }
    |   T_NUM_STRING                                    { $$ = Node("OffsetNumString").attribute("value", $1) }
    |   '-' T_NUM_STRING                                { $$ = Node("OffsetNegateNumString").attribute("value", $2) }
    |   T_VARIABLE                                      { $$ = Node("OffsetVariable").attribute("value", $1) }
;

internal_functions_in_yacc:
        T_ISSET '(' isset_variables possible_comma ')'  { $$ = $3; }
    |   T_EMPTY '(' expr ')'                            { $$ = Node("Empty").append($3); }
    |   T_INCLUDE expr                                  { $$ = Node("Include").append($2); }
    |   T_INCLUDE_ONCE expr                             { $$ = Node("IncludeOnce").append($2); }
    |   T_EVAL '(' expr ')'                             { $$ = Node("Eval").append($3); }
    |   T_REQUIRE expr                                  { $$ = Node("Require").append($2); }
    |   T_REQUIRE_ONCE expr                             { $$ = Node("RequireOnce").append($2); }
;

isset_variables:
        isset_variable                                  { $$ = $1; }
    |   isset_variables ',' isset_variable              { $$ = Node("AndIsset").append($1).append($3); }
;

isset_variable:
    expr                                                { $$ = Node("Isset").append($1) }
;

/////////////////////////////////////////////////////////////////////////

%%

const src = `
<?php

namespace Test;

/**
 * Class foo
 */
class foo
{
    
}
`

func main() {
  yyDebug        = 0
  yyErrorVerbose = true
  l := newLexer(bytes.NewBufferString(src), os.Stdout, "file.name")
  yyParse(l)
}