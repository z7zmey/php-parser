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

%%

/////////////////////////////////////////////////////////////////////////

start:
    top_statement_list                                { fmt.Println($1) }
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
        T_STRING { $$ = Node("identifier") }
    |   semi_reserved  { $$ = Node("reserved") }
;

namespace_name:
        T_STRING                                      { $$ = Node("Namespace").append(Node($1)) }
    |   namespace_name T_NS_SEPARATOR T_STRING        { $$ = $1.append(Node($3)) }
;

top_statement_list:
        top_statement_list top_statement              { $$ = $1.append($2); }
    |   /* empty */                                   { $$ = Node("Statements") }
;

top_statement:
        statement                                     { $$ = $1 }
    |   function_declaration_statement                { $$ = $1 }
    |   T_INCLUDE identifier ';'                      { $$ = $2; /*TODO: identifier stub, refactor it*/ }
    |   T_NAMESPACE namespace_name ';'                { $$ = $2; }
    |   class_declaration_statement                   { $$ = $1; }
;

inner_statement_list:
        inner_statement_list inner_statement          { $$ = $1.append($2); }
    |   /* empty */                                   { $$ = Node("statement_list") }
;

inner_statement:
    statement                                         { $$ = $1; }
    |   class_declaration_statement                   { $$ = $1; }

statement:
    '{' inner_statement_list '}'                      { $$ = $2; }

class_declaration_statement:
        class_modifiers T_CLASS T_STRING '{' '}'      { $$ = $1.attribute("name", $3) }
    |   T_CLASS T_STRING '{' '}'                      { $$ = Node("Class").attribute("name", $2) }
;
class_modifiers:
        class_modifier                                { $$ = Node("Class").attribute($1, "true") }
    |   class_modifiers class_modifier                { $$ = $1.attribute($2, "true") }
;

class_modifier:
        T_ABSTRACT                                    { $$ = "abstract" }
    |   T_FINAL                                       { $$ = "final" }
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
        non_empty_parameter_list                      { $$ = $1; }
    |   /* empty */                                   { $$ = Node("Parameter list"); }
;
non_empty_parameter_list:
        parameter                                     { $$ = Node("Parameter list").append($1) }
    |   non_empty_parameter_list ',' parameter        { $$ = $1.append($3); }
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
;

optional_type:
        /* empty */                                   { $$ = Node("No type") }
    |   type_expr                                     { $$ = $1; }
;

returns_ref:
        /* empty */                                   { $$ = "false"; }
    |   '&'                                           { $$ = "true"; }
;

is_reference:
        /* empty */                                   { $$ = "false"; }
    |   '&'                                           { $$ = "true"; }
;

is_variadic:
        /* empty */                                   { $$ = "false"; }
    |   T_ELLIPSIS                                    { $$ = "true"; }
;

type_expr:
        type                                          { $$ = $1; }
    |   '?' type                                      { $$ = $2; $$.attribute("nullable", "true") }
;

type:
        T_ARRAY                                       { $$ = Node("array type"); }
    |   T_CALLABLE                                    { $$ = Node("callable type"); }
;

return_type:
        /* empty */                                   { $$ = Node("void"); }
    |   ':' type_expr                                 { $$ = $2; }
;

/////////////////////////////////////////////////////////////////////////

%%

const src = `<?
namespace foo\bar\test;

function test(array $a, array $b) {

}

`

func main() {
  yyDebug        = 0
  yyErrorVerbose = true
  l := newLexer(bytes.NewBufferString(src), os.Stdout, "file.name")
  yyParse(l)
}