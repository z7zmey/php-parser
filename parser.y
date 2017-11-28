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

%type <value> class_modifier
%type <value> is_reference
%type <value> is_variadic
%type <value> returns_ref

%type <node> identifier
%type <node> top_statement
%type <node> namespace_name
%type <node> namespace_name_parts
%type <node> name
%type <node> top_statement_list
%type <node> statement
%type <node> inner_statement
%type <node> inner_statement_list
%type <node> class_modifiers
%type <node> class_declaration_statement
%type <node> function_declaration_statement
%type <node> optional_type
%type <node> return_type
%type <node> type_expr
%type <node> type
%type <node> parameter_list
%type <node> non_empty_parameter_list
%type <node> parameter
%type <node> expr
%type <node> expr_without_variable
%type <node> callable_variable
%type <node> variable
%type <node> simple_variable
%type <node> if_stmt_without_else
%type <node> if_stmt
%type <node> alt_if_stmt_without_else
%type <node> alt_if_stmt
%type <node> while_statement
%type <node> for_exprs
%type <node> non_empty_for_exprs
%type <node> for_statement
%type <node> switch_case_list
%type <node> case_list
%type <node> optional_expr

%%

/////////////////////////////////////////////////////////////////////////

start:
    top_statement_list                                  { fmt.Println($1) }
;

reserved_non_modifiers:
      T_INCLUDE | T_INCLUDE_ONCE | T_EVAL | T_REQUIRE | T_REQUIRE_ONCE | T_LOGICAL_OR | T_LOGICAL_XOR | T_LOGICAL_AND
    | T_INSTANCEOF | T_NEW | T_CLONE | T_EXIT | T_IF | T_ELSEIF | T_ELSE | T_ENDIF | T_ECHO | T_DO | T_WHILE | T_ENDWHILE
    | T_FOR | T_ENDFOR | T_FOREACH | T_ENDFOREACH | T_DECLARE | T_ENDDECLARE | T_AS | T_TRY | T_CATCH | T_FINALLY
    | T_THROW | T_USE | T_INSTEADOF | T_GLOBAL | T_VAR | T_UNSET | T_ISSET | T_EMPTY | T_CONTINUE | T_GOTO
    | T_FUNCTION | T_CONST | T_RETURN | T_PRINT | T_YIELD | T_LIST | T_SWITCH | T_ENDSWITCH | T_CASE | T_DEFAULT | T_BREAK
    | T_ARRAY | T_CALLABLE | T_EXTENDS | T_IMPLEMENTS | T_NAMESPACE | T_TRAIT | T_INTERFACE | T_CLASS
    | T_CLASS_C | T_TRAIT_C | T_FUNC_C | T_METHOD_C | T_LINE | T_FILE | T_DIR | T_NS_C
;

semi_reserved:
        reserved_non_modifiers
    | T_STATIC | T_ABSTRACT | T_FINAL | T_PRIVATE | T_PROTECTED | T_PUBLIC
;

identifier:
        T_STRING                                        { $$ = Node("identifier") }
    |   semi_reserved                                   { $$ = Node("reserved") }
;

namespace_name_parts:
        T_STRING                                        { $$ = Node("NamespaceParts").append(Node($1)) }
    |   namespace_name_parts T_NS_SEPARATOR T_STRING    { $$ = $1.append(Node($3)) }
;

namespace_name:
      namespace_name_parts                              { $$ = $1; }
;

name:
      namespace_name                                    { $$ = Node("Name").append($1); }
    | T_NS_SEPARATOR namespace_name                     { $$ = Node("Name").append($2).attribute("FullyQualified", "true"); }
    | T_NAMESPACE T_NS_SEPARATOR namespace_name         { $$ = Node("Name").append($3).attribute("Relative", "true"); }
;

top_statement_list:
        top_statement_list top_statement                { $$ = $1.append($2); }
    |   /* empty */                                     { $$ = Node("Statements") }
;

top_statement:
        statement                                       { $$ = $1 }
    |   function_declaration_statement                  { $$ = $1 }
    |   T_INCLUDE identifier ';'                        { $$ = $2; /*TODO: identifier stub, refactor it*/ }
    |   T_NAMESPACE namespace_name ';'                  { $$ = Node("Namespace").append($2); }
    |   class_declaration_statement                     { $$ = $1; }
;

inner_statement_list:
        inner_statement_list inner_statement            { $$ = $1.append($2); }
    |   /* empty */                                     { $$ = Node("stmt") }
;

inner_statement:
    statement                                           { $$ = $1; }
    |   function_declaration_statement                  { $$ = $1; }
    |   class_declaration_statement                     { $$ = $1; }

statement:
    '{' inner_statement_list '}'                        { $$ = $2; }
    |   if_stmt                                         { $$ = $1; }
    |   alt_if_stmt                                     { $$ = $1; }
    |   T_WHILE '(' expr ')' while_statement            { $$ = Node("While").append(Node("expr").append($3)).append(Node("stmt").append($5)); }
    |   T_DO statement T_WHILE '(' expr ')' ';'         { $$ = Node("DoWhile").append(Node("expr").append($5)).append(Node("stmt").append($2)); }
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
    |   expr ';'                                        { $$ = $1; }

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

while_statement:
        statement                                       { $$ = $1; }
    |   ':' inner_statement_list T_ENDWHILE ';'         { $$ = $2; }
;

for_exprs:
        /* empty */                                     { $$ = Node("null"); }
    |   non_empty_for_exprs                             { $$ = $1; }
;
non_empty_for_exprs:
        non_empty_for_exprs ',' expr                    { $$ = $1.append($3) }
    |   expr                                            { $$ = Node("ExpressionList").append($1) }
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

for_statement:
        statement                                       { $$ = $1; }
    |    ':' inner_statement_list T_ENDFOR ';'          { $$ = $2; }
;

class_declaration_statement:
        class_modifiers T_CLASS T_STRING '{' '}'        { $$ = $1.attribute("name", $3) }
    |   T_CLASS T_STRING '{' '}'                        { $$ = Node("Class").attribute("name", $2) }
;
class_modifiers:
        class_modifier                                  { $$ = Node("Class").attribute($1, "true") }
    |   class_modifiers class_modifier                  { $$ = $1.attribute($2, "true") }
;

class_modifier:
        T_ABSTRACT                                      { $$ = "abstract" }
    |   T_FINAL                                         { $$ = "final" }
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

returns_ref:
        /* empty */                                     { $$ = "false"; }
    |   '&'                                             { $$ = "true"; }
;

is_reference:
        /* empty */                                     { $$ = "false"; }
    |   '&'                                             { $$ = "true"; }
;

is_variadic:
        /* empty */                                     { $$ = "false"; }
    |   T_ELLIPSIS                                      { $$ = "true"; }
;

type_expr:
        type                                            { $$ = $1; }
    |   '?' type                                        { $$ = $2; $$.attribute("nullable", "true") }
;

type:
        name                                            { $$ = $1; }
    |   T_ARRAY                                         { $$ = Node("array type"); }
    |   T_CALLABLE                                      { $$ = Node("callable type"); }
;

return_type:
        /* empty */                                     { $$ = Node("void"); }
    |   ':' type_expr                                   { $$ = $2; }
;

expr_without_variable:
    variable '=' expr                                   { $$ = Node("Assign").append($1).append($3); }
    |   variable '=' '&' expr                           { $$ = Node("AssignRef").append($1).append($4); }
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
    |   expr T_SPACESHIP expr                           { $$ = Node("Spaceship").append($1).append($3) }
    |   expr '<' expr                                   { $$ = Node("Smaller").append($1).append($3) }
    |   expr T_IS_SMALLER_OR_EQUAL expr                 { $$ = Node("SmallerOrEqual").append($1).append($3) }
    |   expr '>' expr                                   { $$ = Node("Greater").append($1).append($3) }
    |   expr T_IS_GREATER_OR_EQUAL expr                 { $$ = Node("GreaterOrEqual").append($1).append($3) }
    |   '(' expr ')'                                    { $$ = $2; }
    |   expr '?' expr ':' expr                          { $$ = Node("Ternary").append($1).append($3).append($5); }
    |   expr '?' ':' expr                               { $$ = Node("Ternary").append($1).append($4); }
    |   expr T_COALESCE expr                            { $$ = Node("Coalesce").append($1).append($3); }
        |   T_EMPTY '(' expr ')'                            { $$ = Node("Empty").append($3); }
        |   T_INCLUDE expr                                  { $$ = Node("Include").append($2); }
        |   T_INCLUDE_ONCE expr                             { $$ = Node("IncludeOnce").append($2); }
        |   T_EVAL '(' expr ')'                             { $$ = Node("Eval").append($3); }
        |   T_REQUIRE expr                                  { $$ = Node("Require").append($2); }
        |   T_REQUIRE_ONCE expr                             { $$ = Node("RequireOnce").append($2); }
    |   T_INT_CAST expr                                 { $$ = Node("CastInt").append($2); }
    |   T_DOUBLE_CAST expr                              { $$ = Node("CastDouble").append($2); }
    |   T_STRING_CAST expr                              { $$ = Node("CastString").append($2); }
    |   T_ARRAY_CAST expr                               { $$ = Node("CastArray").append($2); }
    |   T_OBJECT_CAST expr                              { $$ = Node("CastObject").append($2); }
    |   T_BOOL_CAST expr                                { $$ = Node("CastBool").append($2); }
    |   T_UNSET_CAST expr                               { $$ = Node("CastUnset").append($2); }
    |   '@' expr                                        { $$ = Node("Silence").append($2); }

    |   T_PRINT expr                                    { $$ = Node("Print").append($2); }
    |   T_YIELD                                         { $$ = Node("Yield"); }
    |   T_YIELD expr                                    { $$ = Node("Yield").append($2); }
    |   T_YIELD expr T_DOUBLE_ARROW expr                { $$ = Node("Yield").append($2).append($4); }
    |   T_YIELD_FROM expr                               { $$ = Node("YieldFrom").append($2); }
;

expr:
        variable                                        { $$ = $1; }
    |   expr_without_variable                           { $$ = $1; }
;

optional_expr:
        /* empty */                                     { $$ = Node("null") }
    |   expr                                            { $$ = $1; }
;

callable_variable:
    simple_variable                                     { $$ = $1; }
;

variable:
    callable_variable                                   { $$ = $1; }
;

simple_variable:
        T_VARIABLE                                      { $$ = Node("Variable").attribute("name", $1); }
    |   '$' '{' expr '}'                                { $$ = $3; }
    |   '$' simple_variable                             { $$ = Node("Variable").append($2); }
;

/////////////////////////////////////////////////////////////////////////

%%

const src = `<?

switch($a) :;
    case $b; 
        $b = $a;
        break;
    default; break;
endswitch;

return $a;

`

func main() {
  yyDebug        = 0
  yyErrorVerbose = true
  l := newLexer(bytes.NewBufferString(src), os.Stdout, "file.name")
  yyParse(l)
}