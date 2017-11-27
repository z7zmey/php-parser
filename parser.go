//line parser.y:2
package main

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type node struct {
	name       string
	children   []node
	attributes map[string]string
}

func (n node) String() string {
	buf := new(bytes.Buffer)
	n.print(buf, " ")
	return buf.String()
}

func (n node) print(out io.Writer, indent string) {
	if len(n.attributes) > 0 {
		fmt.Fprintf(out, "\n%v%v %s", indent, n.name, n.attributes)
	} else {
		fmt.Fprintf(out, "\n%v%v", indent, n.name)
	}
	for _, nn := range n.children {
		nn.print(out, indent+"  ")
	}
}

func Node(name string) node {
	return node{name: name, attributes: make(map[string]string)}
}

func (n node) append(nn ...node) node {
	n.children = append(n.children, nn...)
	return n
}

func (n node) attribute(key string, value string) node {
	n.attributes[key] = value
	return n
}

//line parser.y:50
type yySymType struct {
	yys   int
	node  node
	token string
	value string
}

const T_INCLUDE = 57346
const T_INCLUDE_ONCE = 57347
const T_EVAL = 57348
const T_REQUIRE = 57349
const T_REQUIRE_ONCE = 57350
const T_LOGICAL_OR = 57351
const T_LOGICAL_XOR = 57352
const T_LOGICAL_AND = 57353
const T_PRINT = 57354
const T_YIELD = 57355
const T_DOUBLE_ARROW = 57356
const T_YIELD_FROM = 57357
const T_PLUS_EQUAL = 57358
const T_MINUS_EQUAL = 57359
const T_MUL_EQUAL = 57360
const T_DIV_EQUAL = 57361
const T_CONCAT_EQUAL = 57362
const T_MOD_EQUAL = 57363
const T_AND_EQUAL = 57364
const T_OR_EQUAL = 57365
const T_XOR_EQUAL = 57366
const T_SL_EQUAL = 57367
const T_SR_EQUAL = 57368
const T_POW_EQUAL = 57369
const T_COALESCE = 57370
const T_BOOLEAN_OR = 57371
const T_BOOLEAN_AND = 57372
const T_IS_EQUAL = 57373
const T_IS_NOT_EQUAL = 57374
const T_IS_IDENTICAL = 57375
const T_IS_NOT_IDENTICAL = 57376
const T_SPACESHIP = 57377
const T_IS_SMALLER_OR_EQUAL = 57378
const T_IS_GREATER_OR_EQUAL = 57379
const T_SL = 57380
const T_SR = 57381
const T_INSTANCEOF = 57382
const T_INC = 57383
const T_DEC = 57384
const T_INT_CAST = 57385
const T_DOUBLE_CAST = 57386
const T_STRING_CAST = 57387
const T_ARRAY_CAST = 57388
const T_OBJECT_CAST = 57389
const T_BOOL_CAST = 57390
const T_UNSET_CAST = 57391
const T_POW = 57392
const T_NEW = 57393
const T_CLONE = 57394
const T_ELSEIF = 57395
const T_ELSE = 57396
const T_ENDIF = 57397
const T_STATIC = 57398
const T_ABSTRACT = 57399
const T_FINAL = 57400
const T_PRIVATE = 57401
const T_PROTECTED = 57402
const T_PUBLIC = 57403
const T_EXIT = 57404
const T_IF = 57405
const T_LNUMBER = 57406
const T_DNUMBER = 57407
const T_STRING = 57408
const T_STRING_VARNAME = 57409
const T_VARIABLE = 57410
const T_NUM_STRING = 57411
const T_INLINE_HTML = 57412
const T_CHARACTER = 57413
const T_BAD_CHARACTER = 57414
const T_ENCAPSED_AND_WHITESPACE = 57415
const T_CONSTANT_ENCAPSED_STRING = 57416
const T_ECHO = 57417
const T_DO = 57418
const T_WHILE = 57419
const T_ENDWHILE = 57420
const T_FOR = 57421
const T_ENDFOR = 57422
const T_FOREACH = 57423
const T_ENDFOREACH = 57424
const T_DECLARE = 57425
const T_ENDDECLARE = 57426
const T_AS = 57427
const T_SWITCH = 57428
const T_ENDSWITCH = 57429
const T_CASE = 57430
const T_DEFAULT = 57431
const T_BREAK = 57432
const T_CONTINUE = 57433
const T_GOTO = 57434
const T_FUNCTION = 57435
const T_CONST = 57436
const T_RETURN = 57437
const T_TRY = 57438
const T_CATCH = 57439
const T_FINALLY = 57440
const T_THROW = 57441
const T_USE = 57442
const T_INSTEADOF = 57443
const T_GLOBAL = 57444
const T_VAR = 57445
const T_UNSET = 57446
const T_ISSET = 57447
const T_EMPTY = 57448
const T_HALT_COMPILER = 57449
const T_CLASS = 57450
const T_TRAIT = 57451
const T_INTERFACE = 57452
const T_EXTENDS = 57453
const T_IMPLEMENTS = 57454
const T_OBJECT_OPERATOR = 57455
const T_LIST = 57456
const T_ARRAY = 57457
const T_CALLABLE = 57458
const T_CLASS_C = 57459
const T_TRAIT_C = 57460
const T_METHOD_C = 57461
const T_FUNC_C = 57462
const T_LINE = 57463
const T_FILE = 57464
const T_COMMENT = 57465
const T_DOC_COMMENT = 57466
const T_OPEN_TAG = 57467
const T_OPEN_TAG_WITH_ECHO = 57468
const T_CLOSE_TAG = 57469
const T_WHITESPACE = 57470
const T_START_HEREDOC = 57471
const T_END_HEREDOC = 57472
const T_DOLLAR_OPEN_CURLY_BRACES = 57473
const T_CURLY_OPEN = 57474
const T_PAAMAYIM_NEKUDOTAYIM = 57475
const T_NAMESPACE = 57476
const T_NS_C = 57477
const T_DIR = 57478
const T_NS_SEPARATOR = 57479
const T_ELLIPSIS = 57480

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"T_INCLUDE",
	"T_INCLUDE_ONCE",
	"T_EVAL",
	"T_REQUIRE",
	"T_REQUIRE_ONCE",
	"','",
	"T_LOGICAL_OR",
	"T_LOGICAL_XOR",
	"T_LOGICAL_AND",
	"T_PRINT",
	"T_YIELD",
	"T_DOUBLE_ARROW",
	"T_YIELD_FROM",
	"'='",
	"T_PLUS_EQUAL",
	"T_MINUS_EQUAL",
	"T_MUL_EQUAL",
	"T_DIV_EQUAL",
	"T_CONCAT_EQUAL",
	"T_MOD_EQUAL",
	"T_AND_EQUAL",
	"T_OR_EQUAL",
	"T_XOR_EQUAL",
	"T_SL_EQUAL",
	"T_SR_EQUAL",
	"T_POW_EQUAL",
	"'?'",
	"':'",
	"T_COALESCE",
	"T_BOOLEAN_OR",
	"T_BOOLEAN_AND",
	"'|'",
	"'^'",
	"'&'",
	"T_IS_EQUAL",
	"T_IS_NOT_EQUAL",
	"T_IS_IDENTICAL",
	"T_IS_NOT_IDENTICAL",
	"T_SPACESHIP",
	"'<'",
	"T_IS_SMALLER_OR_EQUAL",
	"'>'",
	"T_IS_GREATER_OR_EQUAL",
	"T_SL",
	"T_SR",
	"'+'",
	"'-'",
	"'.'",
	"'*'",
	"'/'",
	"'%'",
	"'!'",
	"T_INSTANCEOF",
	"'~'",
	"T_INC",
	"T_DEC",
	"T_INT_CAST",
	"T_DOUBLE_CAST",
	"T_STRING_CAST",
	"T_ARRAY_CAST",
	"T_OBJECT_CAST",
	"T_BOOL_CAST",
	"T_UNSET_CAST",
	"'@'",
	"T_POW",
	"'['",
	"T_NEW",
	"T_CLONE",
	"T_ELSEIF",
	"T_ELSE",
	"T_ENDIF",
	"T_STATIC",
	"T_ABSTRACT",
	"T_FINAL",
	"T_PRIVATE",
	"T_PROTECTED",
	"T_PUBLIC",
	"T_EXIT",
	"T_IF",
	"T_LNUMBER",
	"T_DNUMBER",
	"T_STRING",
	"T_STRING_VARNAME",
	"T_VARIABLE",
	"T_NUM_STRING",
	"T_INLINE_HTML",
	"T_CHARACTER",
	"T_BAD_CHARACTER",
	"T_ENCAPSED_AND_WHITESPACE",
	"T_CONSTANT_ENCAPSED_STRING",
	"T_ECHO",
	"T_DO",
	"T_WHILE",
	"T_ENDWHILE",
	"T_FOR",
	"T_ENDFOR",
	"T_FOREACH",
	"T_ENDFOREACH",
	"T_DECLARE",
	"T_ENDDECLARE",
	"T_AS",
	"T_SWITCH",
	"T_ENDSWITCH",
	"T_CASE",
	"T_DEFAULT",
	"T_BREAK",
	"T_CONTINUE",
	"T_GOTO",
	"T_FUNCTION",
	"T_CONST",
	"T_RETURN",
	"T_TRY",
	"T_CATCH",
	"T_FINALLY",
	"T_THROW",
	"T_USE",
	"T_INSTEADOF",
	"T_GLOBAL",
	"T_VAR",
	"T_UNSET",
	"T_ISSET",
	"T_EMPTY",
	"T_HALT_COMPILER",
	"T_CLASS",
	"T_TRAIT",
	"T_INTERFACE",
	"T_EXTENDS",
	"T_IMPLEMENTS",
	"T_OBJECT_OPERATOR",
	"T_LIST",
	"T_ARRAY",
	"T_CALLABLE",
	"T_CLASS_C",
	"T_TRAIT_C",
	"T_METHOD_C",
	"T_FUNC_C",
	"T_LINE",
	"T_FILE",
	"T_COMMENT",
	"T_DOC_COMMENT",
	"T_OPEN_TAG",
	"T_OPEN_TAG_WITH_ECHO",
	"T_CLOSE_TAG",
	"T_WHITESPACE",
	"T_START_HEREDOC",
	"T_END_HEREDOC",
	"T_DOLLAR_OPEN_CURLY_BRACES",
	"T_CURLY_OPEN",
	"T_PAAMAYIM_NEKUDOTAYIM",
	"T_NAMESPACE",
	"T_NS_C",
	"T_DIR",
	"T_NS_SEPARATOR",
	"T_ELLIPSIS",
	"';'",
	"'{'",
	"'}'",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:335
const src = `<?
namespace foo\bar\test;

function test(array $a, array $b) {

}

`

func main() {
	yyDebug = 0
	yyErrorVerbose = true
	l := newLexer(bytes.NewBufferString(src), os.Stdout, "file.name")
	yyParse(l)
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 112,
	162, 100,
	-2, 104,
}

const yyPrivate = 57344

const yyLast = 236

var yyAct = [...]int{

	26, 27, 28, 29, 30, 125, 31, 32, 33, 69,
	70, 112, 124, 114, 135, 113, 110, 102, 101, 103,
	134, 122, 123, 95, 119, 120, 137, 117, 14, 15,
	111, 109, 108, 100, 94, 128, 97, 121, 13, 131,
	126, 107, 106, 18, 8, 4, 19, 1, 116, 115,
	99, 130, 34, 118, 5, 11, 104, 2, 93, 3,
	16, 96, 133, 127, 0, 0, 35, 36, 39, 40,
	41, 20, 21, 22, 23, 24, 25, 37, 38, 98,
	6, 17, 0, 0, 0, 0, 0, 14, 15, 0,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 72, 73, 74, 75, 76, 64, 65, 66, 67,
	68, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 0, 84, 82, 83, 79, 80, 0, 71,
	77, 78, 85, 86, 88, 87, 89, 90, 12, 14,
	15, 122, 123, 0, 0, 0, 0, 129, 0, 81,
	92, 91, 14, 15, 132, 0, 136, 0, 0, 138,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	9, 139, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 10, 0,
	12, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 12, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 9, 105, 0, 0, 0, 0, 0, 7,
	0, 0, 0, 0, 0, 9,
}
var yyPact = [...]int{

	-1000, -1000, 76, -1000, -1000, -1000, -4, -51, -1000, -1000,
	-1, -48, -52, -1000, -1000, -1000, -140, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -139, -1000, 63, -53, -1000, -54, -1000,
	-143, -1000, -55, -1000, -1000, -1000, -1000, -1000, -150, -144,
	-147, -1000, 7, -148, -1000, -157, 31, -1000, -2, -1000,
	-1000, -113, -1000, -1000, -1000, 8, 7, -137, -1000, -1000,
	-145, 7, -1000, -61, -1000, -1000, -1000, -1000, 11, -1000,
}
var yyPgo = [...]int{

	0, 38, 63, 62, 61, 60, 59, 58, 57, 42,
	56, 23, 55, 41, 54, 53, 51, 24, 25, 49,
	48, 27, 47, 46, 43,
}
var yyR1 = [...]int{

	0, 22, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 24,
	24, 24, 24, 24, 24, 24, 5, 5, 7, 7,
	8, 8, 6, 6, 6, 6, 6, 11, 11, 10,
	10, 9, 13, 13, 12, 12, 1, 1, 14, 19,
	19, 20, 20, 21, 15, 15, 4, 4, 2, 2,
	3, 3, 17, 17, 18, 18, 16, 16,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	2, 0, 1, 1, 3, 3, 1, 2, 0, 1,
	1, 3, 5, 4, 1, 2, 1, 1, 10, 1,
	0, 1, 3, 4, 0, 1, 0, 1, 0, 1,
	0, 1, 1, 2, 1, 1, 0, 2,
}
var yyChk = [...]int{

	-1000, -22, -8, -6, -9, -14, 4, 153, -13, 159,
	112, -12, 127, -1, 76, 77, -5, 85, -24, -23,
	75, 76, 77, 78, 79, 80, 4, 5, 6, 7,
	8, 10, 11, 12, 56, 70, 71, 81, 82, 72,
	73, 74, 94, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 115, 116, 117, 118, 119, 120, 121,
	122, 123, 124, 125, 110, 111, 112, 113, 114, 13,
	14, 133, 105, 106, 107, 108, 109, 134, 135, 130,
	131, 153, 128, 129, 127, 136, 137, 139, 138, 140,
	141, 155, 154, -7, 85, -11, -4, 37, 127, -1,
	85, 158, 156, 158, -10, 160, -9, -13, 85, 85,
	159, 85, 161, 159, 160, -19, -20, -21, -15, -17,
	-18, 30, 134, 135, 160, 162, 9, -2, 37, -18,
	-16, 31, -21, -3, 157, 159, -17, 87, -11, 160,
}
var yyDef = [...]int{

	81, -2, 1, 80, 82, 83, 0, 0, 86, 88,
	106, 0, 0, 94, 96, 97, 0, 76, 77, 69,
	70, 71, 72, 73, 74, 75, 2, 3, 4, 5,
	6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 49, 50, 51, 52, 53, 54, 55,
	56, 57, 58, 59, 60, 61, 62, 63, 64, 65,
	66, 67, 68, 0, 78, 0, 0, 107, 0, 95,
	0, 84, 0, 85, 87, 91, 89, 90, 0, 0,
	0, 79, -2, 0, 93, 0, 99, 101, 108, 105,
	112, 0, 114, 115, 92, 116, 104, 110, 109, 113,
	0, 0, 102, 0, 111, 88, 117, 103, 0, 98,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 3, 3, 3, 54, 37, 3,
	161, 162, 52, 49, 9, 50, 51, 53, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 31, 158,
	43, 17, 45, 30, 67, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 69, 3, 3, 36, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 159, 35, 160, 57,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 10, 11, 12,
	13, 14, 15, 16, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 32, 33, 34, 38,
	39, 40, 41, 42, 44, 46, 47, 48, 56, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 68, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
	121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
	151, 152, 153, 154, 155, 156, 157,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:199
		{
			fmt.Println(yyDollar[1].node)
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:218
		{
			yyVAL.node = Node("identifier")
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:219
		{
			yyVAL.node = Node("reserved")
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:223
		{
			yyVAL.node = Node("Namespace").append(Node(yyDollar[1].token))
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:224
		{
			yyVAL.node = yyDollar[1].node.append(Node(yyDollar[3].token))
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:228
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[2].node)
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:229
		{
			yyVAL.node = Node("Statements")
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:233
		{
			yyVAL.node = yyDollar[1].node
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:234
		{
			yyVAL.node = yyDollar[1].node
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:235
		{
			yyVAL.node = yyDollar[2].node /*TODO: identifier stub, refactor it*/
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:236
		{
			yyVAL.node = yyDollar[2].node
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:237
		{
			yyVAL.node = yyDollar[1].node
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:241
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[2].node)
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:242
		{
			yyVAL.node = Node("statement_list")
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:246
		{
			yyVAL.node = yyDollar[1].node
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:247
		{
			yyVAL.node = yyDollar[1].node
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:250
		{
			yyVAL.node = yyDollar[2].node
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:253
		{
			yyVAL.node = yyDollar[1].node.attribute("name", yyDollar[3].token)
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:254
		{
			yyVAL.node = Node("Class").attribute("name", yyDollar[2].token)
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:257
		{
			yyVAL.node = Node("Class").attribute(yyDollar[1].value, "true")
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:258
		{
			yyVAL.node = yyDollar[1].node.attribute(yyDollar[2].value, "true")
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:262
		{
			yyVAL.value = "abstract"
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:263
		{
			yyVAL.value = "final"
		}
	case 98:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line parser.y:268
		{
			fmt.Println("adsfadsf")
			yyVAL.node = Node("Function").
				attribute("name", yyDollar[3].token).
				attribute("returns_ref", yyDollar[2].value).
				append(yyDollar[5].node).
				append(yyDollar[7].node).
				append(yyDollar[9].node)
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:280
		{
			yyVAL.node = yyDollar[1].node
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:281
		{
			yyVAL.node = Node("Parameter list")
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:284
		{
			yyVAL.node = Node("Parameter list").append(yyDollar[1].node)
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:285
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[3].node)
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:289
		{
			yyVAL.node = Node("Parameter").
				append(yyDollar[1].node).
				attribute("is_reference", yyDollar[2].value).
				attribute("is_variadic", yyDollar[3].value).
				attribute("var", yyDollar[4].token)
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:299
		{
			yyVAL.node = Node("No type")
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:300
		{
			yyVAL.node = yyDollar[1].node
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:304
		{
			yyVAL.value = "false"
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:305
		{
			yyVAL.value = "true"
		}
	case 108:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:309
		{
			yyVAL.value = "false"
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:310
		{
			yyVAL.value = "true"
		}
	case 110:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:314
		{
			yyVAL.value = "false"
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:315
		{
			yyVAL.value = "true"
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:319
		{
			yyVAL.node = yyDollar[1].node
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:320
		{
			yyVAL.node = yyDollar[2].node
			yyVAL.node.attribute("nullable", "true")
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:324
		{
			yyVAL.node = Node("array type")
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:325
		{
			yyVAL.node = Node("callable type")
		}
	case 116:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:329
		{
			yyVAL.node = Node("void")
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:330
		{
			yyVAL.node = yyDollar[2].node
		}
	}
	goto yystack /* stack new state and value */
}
