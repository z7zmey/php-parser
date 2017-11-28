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
const T_NOELSE = 57395
const T_ELSEIF = 57396
const T_ELSE = 57397
const T_ENDIF = 57398
const T_STATIC = 57399
const T_ABSTRACT = 57400
const T_FINAL = 57401
const T_PRIVATE = 57402
const T_PROTECTED = 57403
const T_PUBLIC = 57404
const T_EXIT = 57405
const T_IF = 57406
const T_LNUMBER = 57407
const T_DNUMBER = 57408
const T_STRING = 57409
const T_STRING_VARNAME = 57410
const T_VARIABLE = 57411
const T_NUM_STRING = 57412
const T_INLINE_HTML = 57413
const T_CHARACTER = 57414
const T_BAD_CHARACTER = 57415
const T_ENCAPSED_AND_WHITESPACE = 57416
const T_CONSTANT_ENCAPSED_STRING = 57417
const T_ECHO = 57418
const T_DO = 57419
const T_WHILE = 57420
const T_ENDWHILE = 57421
const T_FOR = 57422
const T_ENDFOR = 57423
const T_FOREACH = 57424
const T_ENDFOREACH = 57425
const T_DECLARE = 57426
const T_ENDDECLARE = 57427
const T_AS = 57428
const T_SWITCH = 57429
const T_ENDSWITCH = 57430
const T_CASE = 57431
const T_DEFAULT = 57432
const T_BREAK = 57433
const T_CONTINUE = 57434
const T_GOTO = 57435
const T_FUNCTION = 57436
const T_CONST = 57437
const T_RETURN = 57438
const T_TRY = 57439
const T_CATCH = 57440
const T_FINALLY = 57441
const T_THROW = 57442
const T_USE = 57443
const T_INSTEADOF = 57444
const T_GLOBAL = 57445
const T_VAR = 57446
const T_UNSET = 57447
const T_ISSET = 57448
const T_EMPTY = 57449
const T_HALT_COMPILER = 57450
const T_CLASS = 57451
const T_TRAIT = 57452
const T_INTERFACE = 57453
const T_EXTENDS = 57454
const T_IMPLEMENTS = 57455
const T_OBJECT_OPERATOR = 57456
const T_LIST = 57457
const T_ARRAY = 57458
const T_CALLABLE = 57459
const T_CLASS_C = 57460
const T_TRAIT_C = 57461
const T_METHOD_C = 57462
const T_FUNC_C = 57463
const T_LINE = 57464
const T_FILE = 57465
const T_COMMENT = 57466
const T_DOC_COMMENT = 57467
const T_OPEN_TAG = 57468
const T_OPEN_TAG_WITH_ECHO = 57469
const T_CLOSE_TAG = 57470
const T_WHITESPACE = 57471
const T_START_HEREDOC = 57472
const T_END_HEREDOC = 57473
const T_DOLLAR_OPEN_CURLY_BRACES = 57474
const T_CURLY_OPEN = 57475
const T_PAAMAYIM_NEKUDOTAYIM = 57476
const T_NAMESPACE = 57477
const T_NS_C = 57478
const T_DIR = 57479
const T_NS_SEPARATOR = 57480
const T_ELLIPSIS = 57481

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
	"T_NOELSE",
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
	"'$'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:506
const src = `<?

if ($a > $b) : 
    $b=$c;
elseif ($a === $b) :

else :

endif;

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
	-1, 72,
	159, 46,
	-2, 207,
	-1, 244,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 179,
	-1, 245,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 180,
	-1, 246,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 181,
	-1, 247,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 182,
	-1, 248,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 183,
	-1, 249,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 184,
	-1, 250,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 185,
	-1, 251,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 186,
	-1, 252,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 187,
	-1, 288,
	163, 117,
	-2, 122,
}

const yyPrivate = 57344

const yyLast = 1995

var yyAct = [...]int{

	13, 134, 306, 304, 312, 307, 322, 54, 308, 288,
	285, 261, 131, 259, 202, 200, 191, 315, 290, 341,
	223, 289, 258, 4, 51, 192, 333, 51, 195, 196,
	197, 198, 199, 135, 201, 331, 203, 204, 205, 206,
	207, 208, 209, 210, 211, 212, 213, 214, 215, 262,
	219, 218, 133, 338, 140, 142, 141, 328, 220, 48,
	49, 226, 343, 133, 133, 192, 284, 54, 201, 257,
	203, 204, 213, 214, 164, 256, 165, 138, 139, 143,
	145, 144, 157, 158, 155, 156, 159, 160, 161, 162,
	163, 153, 154, 147, 148, 146, 149, 151, 152, 216,
	52, 310, 311, 52, 170, 173, 175, 174, 171, 172,
	168, 150, 150, 310, 311, 149, 151, 152, 325, 167,
	314, 335, 330, 313, 263, 345, 323, 224, 56, 57,
	5, 150, 314, 225, 1, 313, 8, 54, 50, 227,
	228, 229, 230, 231, 232, 233, 234, 235, 236, 237,
	238, 239, 240, 241, 242, 243, 244, 245, 246, 247,
	248, 249, 250, 251, 252, 253, 255, 11, 18, 153,
	154, 147, 148, 146, 149, 151, 152, 264, 266, 267,
	268, 269, 270, 271, 272, 273, 274, 275, 276, 277,
	150, 217, 278, 260, 147, 148, 146, 149, 151, 152,
	19, 280, 21, 281, 10, 17, 23, 321, 20, 140,
	142, 141, 303, 150, 302, 334, 305, 283, 169, 15,
	221, 2, 309, 132, 3, 53, 193, 194, 166, 164,
	337, 165, 138, 139, 143, 145, 144, 157, 158, 155,
	156, 159, 160, 161, 162, 163, 153, 154, 147, 148,
	146, 149, 151, 152, 324, 287, 0, 0, 0, 0,
	291, 0, 292, 0, 0, 293, 294, 150, 160, 161,
	162, 163, 153, 154, 147, 148, 146, 149, 151, 152,
	0, 0, 0, 298, 0, 0, 300, 301, 0, 0,
	0, 0, 0, 150, 0, 0, 66, 67, 68, 69,
	70, 0, 73, 74, 75, 71, 72, 0, 47, 0,
	0, 0, 0, 0, 326, 0, 319, 0, 327, 0,
	0, 0, 332, 0, 0, 0, 0, 336, 0, 0,
	0, 0, 340, 339, 0, 0, 0, 329, 342, 0,
	0, 27, 28, 344, 0, 0, 347, 29, 76, 30,
	25, 26, 37, 38, 39, 40, 41, 42, 43, 44,
	299, 0, 77, 64, 0, 80, 81, 82, 58, 59,
	60, 61, 62, 63, 78, 79, 0, 0, 55, 0,
	51, 0, 0, 0, 0, 0, 0, 83, 84, 85,
	86, 87, 88, 89, 90, 91, 92, 93, 110, 111,
	112, 113, 114, 104, 105, 106, 107, 108, 94, 95,
	96, 97, 98, 99, 100, 101, 102, 103, 65, 0,
	122, 120, 121, 117, 118, 0, 109, 115, 116, 123,
	124, 126, 125, 127, 128, 0, 0, 0, 0, 0,
	136, 33, 34, 35, 36, 0, 119, 130, 129, 45,
	46, 0, 47, 0, 31, 0, 52, 0, 139, 143,
	145, 144, 157, 158, 155, 156, 159, 160, 161, 162,
	163, 153, 154, 147, 148, 146, 149, 151, 152, 0,
	0, 0, 0, 0, 0, 27, 28, 0, 0, 0,
	0, 29, 150, 30, 25, 26, 37, 38, 39, 40,
	41, 42, 43, 44, 0, 0, 0, 24, 0, 0,
	0, 0, 0, 48, 49, 0, 0, 0, 0, 22,
	0, 0, 0, 0, 51, 0, 0, 0, 0, 0,
	0, 0, 12, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 14,
	140, 142, 141, 0, 0, 0, 136, 33, 34, 35,
	36, 0, 32, 0, 16, 45, 46, 0, 47, 0,
	164, 0, 165, 138, 139, 143, 145, 144, 157, 158,
	155, 156, 159, 160, 161, 162, 163, 153, 154, 147,
	148, 146, 149, 151, 152, 0, 9, 346, 31, 0,
	52, 27, 28, 0, 0, 0, 0, 29, 150, 30,
	25, 26, 37, 38, 39, 40, 41, 42, 43, 44,
	0, 0, 0, 24, 0, 0, 0, 318, 0, 48,
	49, 0, 0, 0, 0, 22, 0, 0, 0, 0,
	51, 0, 0, 0, 0, 0, 0, 0, 12, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 14, 0, 0, 0, 0,
	0, 0, 136, 33, 34, 35, 36, 0, 32, 0,
	16, 45, 46, 0, 47, 157, 158, 155, 156, 159,
	160, 161, 162, 163, 153, 154, 147, 148, 146, 149,
	151, 152, 0, 317, 0, 0, 0, 0, 0, 0,
	0, 0, 9, 0, 31, 150, 52, 27, 28, 0,
	0, 0, 0, 29, 0, 30, 25, 26, 37, 38,
	39, 40, 41, 42, 43, 44, 0, 0, 0, 24,
	0, 0, 0, 0, 0, 48, 49, 0, 0, 0,
	0, 22, 0, 0, 0, 0, 51, 0, 0, 0,
	0, 0, 0, 0, 12, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 14, 140, 142, 141, 0, 0, 0, 6, 33,
	34, 35, 36, 0, 32, 0, 16, 45, 46, 0,
	47, 0, 164, 0, 165, 138, 139, 143, 145, 144,
	157, 158, 155, 156, 159, 160, 161, 162, 163, 153,
	154, 147, 148, 146, 149, 151, 152, 0, 9, 222,
	31, 0, 52, 27, 28, 0, 0, 0, 0, 29,
	150, 30, 25, 26, 37, 38, 39, 40, 41, 42,
	43, 44, 0, 0, 0, 24, 0, 0, 0, 0,
	0, 48, 49, 0, 0, 0, 0, 22, 0, 0,
	0, 0, 51, 0, 0, 0, 0, 0, 0, 0,
	12, 176, 177, 178, 179, 181, 182, 183, 184, 185,
	186, 187, 188, 180, 0, 0, 0, 14, 0, 0,
	0, 0, 0, 0, 136, 33, 34, 35, 36, 0,
	32, 0, 16, 45, 46, 0, 47, 0, 0, 0,
	0, 0, 189, 190, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 316, 0, 0, 7, 0,
	0, 0, 0, 0, 9, 0, 31, 0, 52, 27,
	28, 0, 0, 0, 0, 29, 0, 30, 25, 26,
	37, 38, 39, 40, 41, 42, 43, 44, 0, 0,
	0, 24, 0, 0, 0, 0, 0, 48, 49, 0,
	0, 0, 0, 22, 0, 0, 0, 0, 51, 136,
	33, 34, 35, 36, 0, 0, 12, 0, 45, 46,
	0, 47, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 14, 0, 0, 320, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 32, 0, 16, 136,
	33, 34, 35, 36, 27, 28, 0, 0, 45, 46,
	29, 47, 30, 25, 26, 37, 38, 39, 40, 41,
	42, 43, 44, 0, 0, 0, 24, 0, 0, 0,
	9, 0, 31, 0, 52, 0, 0, 0, 22, 0,
	0, 0, 0, 51, 27, 28, 0, 0, 0, 0,
	29, 12, 30, 25, 26, 37, 38, 39, 40, 41,
	42, 43, 44, 0, 0, 0, 24, 0, 0, 0,
	140, 142, 141, 0, 0, 0, 0, 0, 22, 0,
	0, 32, 0, 51, 0, 0, 0, 0, 0, 0,
	164, 12, 165, 138, 139, 143, 145, 144, 157, 158,
	155, 156, 159, 160, 161, 162, 163, 153, 154, 147,
	148, 146, 149, 151, 152, 9, 0, 31, 0, 52,
	0, 32, 0, 0, 0, 0, 0, 0, 150, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 140, 142,
	141, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 9, 0, 31, 164, 52,
	165, 138, 139, 143, 145, 144, 157, 158, 155, 156,
	159, 160, 161, 162, 163, 153, 154, 147, 148, 146,
	149, 151, 152, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 150, 0, 0, 0,
	0, 0, 140, 142, 141, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 164, 297, 165, 138, 139, 143, 145, 144,
	157, 158, 155, 156, 159, 160, 161, 162, 163, 153,
	154, 147, 148, 146, 149, 151, 152, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	150, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	140, 142, 141, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	164, 296, 165, 138, 139, 143, 145, 144, 157, 158,
	155, 156, 159, 160, 161, 162, 163, 153, 154, 147,
	148, 146, 149, 151, 152, 136, 33, 34, 35, 36,
	0, 0, 0, 0, 45, 46, 0, 47, 150, 0,
	0, 0, 0, 0, 0, 136, 33, 34, 35, 36,
	0, 0, 0, 0, 45, 46, 0, 47, 265, 0,
	0, 0, 0, 0, 0, 295, 0, 0, 0, 0,
	27, 28, 254, 0, 0, 0, 29, 0, 30, 25,
	26, 37, 38, 39, 40, 41, 42, 43, 44, 0,
	27, 28, 24, 0, 0, 0, 29, 0, 30, 25,
	26, 37, 38, 39, 40, 41, 42, 43, 44, 51,
	0, 0, 24, 0, 0, 136, 33, 34, 35, 36,
	0, 0, 0, 0, 45, 46, 0, 47, 0, 51,
	0, 0, 0, 279, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 32, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	27, 28, 0, 0, 0, 0, 29, 32, 30, 25,
	26, 37, 38, 39, 40, 41, 42, 43, 44, 0,
	0, 0, 24, 31, 0, 52, 140, 142, 141, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 51,
	0, 0, 0, 31, 0, 52, 164, 0, 165, 138,
	139, 143, 145, 144, 157, 158, 155, 156, 159, 160,
	161, 162, 163, 153, 154, 147, 148, 146, 149, 151,
	152, 0, 0, 0, 0, 0, 0, 32, 0, 0,
	0, 0, 0, 0, 150, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 140, 142, 141, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 31, 0, 52, 164, 286, 165, 138,
	139, 143, 145, 144, 157, 158, 155, 156, 159, 160,
	161, 162, 163, 153, 154, 147, 148, 146, 149, 151,
	152, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 150, 0, 140, 142, 141, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 137, 164, 0, 165, 138,
	139, 143, 145, 144, 157, 158, 155, 156, 159, 160,
	161, 162, 163, 153, 154, 147, 148, 146, 149, 151,
	152, 142, 141, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 150, 0, 0, 0, 0, 0,
	164, 0, 165, 138, 139, 143, 145, 144, 157, 158,
	155, 156, 159, 160, 161, 162, 163, 153, 154, 147,
	148, 146, 149, 151, 152, 141, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 150, 0,
	0, 0, 0, 164, 0, 165, 138, 139, 143, 145,
	144, 157, 158, 155, 156, 159, 160, 161, 162, 163,
	153, 154, 147, 148, 146, 149, 151, 152, 282, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 150, 0, 164, 0, 165, 138, 139, 143, 145,
	144, 157, 158, 155, 156, 159, 160, 161, 162, 163,
	153, 154, 147, 148, 146, 149, 151, 152, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	164, 150, 165, 138, 139, 143, 145, 144, 157, 158,
	155, 156, 159, 160, 161, 162, 163, 153, 154, 147,
	148, 146, 149, 151, 152, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 150, 165,
	138, 139, 143, 145, 144, 157, 158, 155, 156, 159,
	160, 161, 162, 163, 153, 154, 147, 148, 146, 149,
	151, 152, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 150, 143, 145, 144, 157,
	158, 155, 156, 159, 160, 161, 162, 163, 153, 154,
	147, 148, 146, 149, 151, 152, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 150,
	145, 144, 157, 158, 155, 156, 159, 160, 161, 162,
	163, 153, 154, 147, 148, 146, 149, 151, 152, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 150, 144, 157, 158, 155, 156, 159, 160,
	161, 162, 163, 153, 154, 147, 148, 146, 149, 151,
	152, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 150,
}
var yyPact = [...]int{

	-1000, -1000, 784, -1000, -1000, -1000, 292, -23, -1000, -1000,
	-1000, -1000, 1025, 1496, 82, -18, 18, 35, 32, 864,
	-1000, -1000, -146, -1000, 1431, -64, -64, 1431, 1431, 1431,
	1431, 1431, -147, 1431, -148, 1431, 1431, 1431, 1431, 1431,
	1431, 1431, 1431, 1431, 1431, 1431, 1431, 1431, -1000, -1000,
	-1000, -1000, -61, -108, 1626, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 1431, -147, 1431, 1431, -148, 1431,
	1431, 1431, 1431, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -109, -99, -1000, 668, -36, 1431, -1000, 1431, 1431,
	1431, 1431, 1431, 1431, 1431, 1431, 1431, 1431, 1431, 1431,
	1431, 1431, 1431, 1431, 1431, 1431, 1431, 1431, 1431, 1431,
	1431, 1431, 1431, 1431, 1361, 1431, -11, -1000, -17, -1000,
	-138, -149, 1025, -151, -110, 93, 1341, 1431, 1431, 1431,
	1431, 1431, 1431, 1431, 1431, 1431, 1431, 1431, 1431, -1000,
	-1000, 1431, -1000, -1000, -1000, 43, 43, 43, 43, 1290,
	1431, 1626, 1431, 1626, 1626, 43, 43, 43, 43, 43,
	43, 43, 43, 1790, 1753, 1790, 1431, -1000, -1000, -1000,
	-20, -1000, -1000, -1000, -1000, -1000, -152, 424, 1861, 1670,
	1790, 1713, 1894, 647, 1926, 63, 63, 63, 43, 43,
	43, 43, 145, 145, 225, 225, 225, 225, 225, 122,
	122, 122, 122, 1566, 1431, 1827, -153, -139, -143, 1431,
	-1000, 1431, -1000, -1000, 1790, 1431, 1790, 1790, 1790, 1790,
	1790, 1790, 1790, 1790, 1790, 1790, 1790, 1790, 1222, -1000,
	1158, 1090, 1431, 199, -1000, 1431, 1431, 1827, -22, -144,
	-1000, 772, 540, 552, 647, 985, -1000, -1000, 1790, -1000,
	44, 1827, -157, 117, -1000, 81, -1000, -1000, -34, -1000,
	-1000, -1000, -1000, -23, -100, -1000, 1025, 91, -124, -1000,
	-1000, -133, 90, -22, -105, -1000, -1000, -1000, -23, -1000,
	-1000, -1000, 900, -1000, -141, -22, -1000, -26, -1000, -1000,
	900, -1000, -1000, 108, 436, 1431, -1000, 1626,
}
var yyPgo = [...]int{

	0, 202, 254, 230, 228, 225, 224, 4, 223, 222,
	221, 20, 220, 1, 219, 133, 127, 216, 215, 2,
	5, 214, 212, 3, 0, 208, 206, 200, 138, 205,
	204, 168, 167, 134, 129, 128,
}
var yyR1 = [...]int{

	0, 33, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 35,
	35, 35, 35, 35, 35, 35, 5, 5, 8, 8,
	7, 9, 9, 9, 10, 10, 6, 6, 6, 6,
	6, 13, 13, 12, 12, 12, 11, 11, 11, 11,
	11, 29, 29, 30, 30, 31, 31, 32, 32, 15,
	15, 14, 14, 1, 1, 16, 21, 21, 22, 22,
	23, 23, 17, 17, 4, 4, 2, 2, 3, 3,
	19, 19, 20, 20, 20, 18, 18, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 24, 24, 26, 27, 28, 28, 28,
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
	1, 1, 2, 3, 2, 0, 1, 1, 3, 3,
	1, 2, 0, 1, 1, 1, 3, 1, 1, 7,
	2, 5, 6, 1, 3, 6, 7, 3, 6, 5,
	4, 1, 2, 1, 1, 10, 1, 0, 1, 3,
	4, 6, 0, 1, 0, 1, 0, 1, 0, 1,
	1, 2, 1, 1, 1, 0, 2, 3, 4, 4,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	4, 3, 4, 2, 2, 4, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 1, 2, 4,
	2, 1, 1, 1, 1, 1, 4, 2,
}
var yyChk = [...]int{

	-1000, -33, -10, -6, -11, -16, 4, 154, -15, 160,
	-30, -32, 96, -24, 113, -14, 128, -29, -31, -27,
	-25, -1, 83, -26, 71, 58, 59, 49, 50, 55,
	57, 162, 126, 5, 6, 7, 8, 60, 61, 62,
	63, 64, 65, 66, 67, 13, 14, 16, 77, 78,
	-28, 88, 164, -5, -24, 86, -35, -34, 76, 77,
	78, 79, 80, 81, 71, 126, 4, 5, 6, 7,
	8, 13, 14, 10, 11, 12, 56, 70, 82, 83,
	73, 74, 75, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 116, 117, 118, 119, 120, 121,
	122, 123, 124, 125, 111, 112, 113, 114, 115, 134,
	106, 107, 108, 109, 110, 135, 136, 131, 132, 154,
	129, 130, 128, 137, 138, 140, 139, 141, 142, 156,
	155, -7, -8, 86, -13, -11, 4, 159, 33, 34,
	10, 12, 11, 35, 37, 36, 51, 49, 50, 52,
	68, 53, 54, 47, 48, 40, 41, 38, 39, 42,
	43, 44, 45, 46, 30, 32, -4, 37, 128, -1,
	86, 73, 74, 73, 75, 74, 17, 18, 19, 20,
	29, 21, 22, 23, 24, 25, 26, 27, 28, 58,
	59, 162, -24, -27, -27, -24, -24, -24, -24, -24,
	162, -24, 162, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, 160, -28, 159, 159,
	157, -12, 161, -11, -16, -15, 97, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, 31, -24, 86, 86, 160, 162,
	-11, 162, 159, 31, -24, 37, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, 163,
	-24, -24, 15, -24, 86, 162, 31, -24, 162, 160,
	161, -24, -24, -13, -24, 163, 163, 163, -24, 161,
	-24, -24, -21, -22, -23, -17, -19, -20, 30, -9,
	135, 136, -7, 157, 154, 161, 163, 163, 75, -11,
	31, 163, 163, 9, -2, 37, -20, -7, 157, -11,
	31, 159, -13, 159, -18, 31, -23, -3, 158, -7,
	-13, 160, -19, 88, -13, 17, 161, -24,
}
var yyDef = [...]int{

	85, -2, 1, 84, 86, 87, 0, 0, 90, 92,
	97, 98, 0, 0, 124, 0, 0, 103, 0, 211,
	212, 111, 0, 214, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 207, 0, 113, 114,
	213, 215, 0, 0, 193, 76, 77, 69, 70, 71,
	72, 73, 74, 75, 12, 39, 2, 3, 4, 5,
	6, 45, -2, 7, 8, 9, 10, 11, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
	25, 26, 27, 28, 29, 30, 31, 32, 33, 34,
	35, 36, 37, 38, 40, 41, 42, 43, 44, 47,
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 0, 80, 78, 0, 0, 0, 100, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 125, 0, 112,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 154,
	156, 0, 140, 155, 157, 175, 176, 177, 178, 0,
	0, 194, 0, 196, 197, 198, 199, 200, 201, 202,
	203, 204, 205, 206, 208, 210, 0, 217, 88, 89,
	0, 91, 96, 93, 94, 95, 0, 158, 159, 160,
	161, 162, 163, 164, 165, 166, 167, 168, 169, 170,
	171, 172, 173, 174, -2, -2, -2, -2, -2, -2,
	-2, -2, -2, 0, 0, 191, 0, 0, 0, 0,
	104, 0, 107, 92, 137, 0, 141, 142, 143, 144,
	145, 146, 147, 148, 150, 151, 152, 153, 0, 188,
	0, 0, 0, 0, 79, 0, 0, 190, -2, 0,
	110, 0, 0, 0, 138, 0, 192, 195, 209, 216,
	0, 189, 0, 116, 118, 126, 123, 130, 0, 132,
	133, 134, 81, 0, 0, 109, 0, 0, 0, 101,
	92, 0, 135, 122, 128, 127, 131, 82, 0, 102,
	92, 108, 105, 99, 0, 0, 119, 0, 129, 83,
	106, 92, 136, 120, 0, 0, 115, 121,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 3, 3, 164, 54, 37, 3,
	162, 163, 52, 49, 9, 50, 51, 53, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 31, 159,
	43, 17, 45, 30, 67, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 69, 3, 3, 36, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 160, 35, 161, 57,
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
	151, 152, 153, 154, 155, 156, 157, 158,
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
		//line parser.y:211
		{
			fmt.Println(yyDollar[1].node)
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:230
		{
			yyVAL.node = Node("identifier")
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:231
		{
			yyVAL.node = Node("reserved")
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:235
		{
			yyVAL.node = Node("NamespaceParts").append(Node(yyDollar[1].token))
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:236
		{
			yyVAL.node = yyDollar[1].node.append(Node(yyDollar[3].token))
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:240
		{
			yyVAL.node = yyDollar[1].node
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:244
		{
			yyVAL.node = Node("Name").append(yyDollar[1].node)
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:245
		{
			yyVAL.node = Node("Name").append(yyDollar[2].node).attribute("FullyQualified", "true")
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:246
		{
			yyVAL.node = Node("Name").append(yyDollar[3].node).attribute("Relative", "true")
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:250
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[2].node)
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:251
		{
			yyVAL.node = Node("Statements")
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:255
		{
			yyVAL.node = yyDollar[1].node
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:256
		{
			yyVAL.node = yyDollar[1].node
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:257
		{
			yyVAL.node = yyDollar[2].node /*TODO: identifier stub, refactor it*/
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:258
		{
			yyVAL.node = Node("Namespace").append(yyDollar[2].node)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:259
		{
			yyVAL.node = yyDollar[1].node
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:263
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[2].node)
		}
	case 92:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:264
		{
			yyVAL.node = Node("statement_list")
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:268
		{
			yyVAL.node = yyDollar[1].node
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:269
		{
			yyVAL.node = yyDollar[1].node
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:270
		{
			yyVAL.node = yyDollar[1].node
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:273
		{
			yyVAL.node = yyDollar[2].node
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:274
		{
			yyVAL.node = yyDollar[1].node
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:275
		{
			yyVAL.node = yyDollar[1].node
		}
	case 99:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:276
		{
			yyVAL.node = Node("Do While").append(yyDollar[2].node).append(yyDollar[5].node)
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:277
		{
			yyVAL.node = yyDollar[1].node
		}
	case 101:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:281
		{
			yyVAL.node = Node("If").append(Node("expr").append(yyDollar[3].node)).append(Node("stmt").append(yyDollar[5].node))
		}
	case 102:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:285
		{
			yyVAL.node = yyDollar[1].node.append(Node("ElseIf").append(Node("expr").append(yyDollar[4].node)).append(Node("stmt").append(yyDollar[6].node)))
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:291
		{
			yyVAL.node = yyDollar[1].node
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:293
		{
			yyVAL.node = yyDollar[1].node.append(Node("Else").append(Node("stmt").append(yyDollar[3].node)))
		}
	case 105:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:300
		{
			yyVAL.node = Node("AltIf").append(Node("expr").append(yyDollar[3].node)).append(Node("stmt").append(yyDollar[6].node))
		}
	case 106:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:304
		{
			yyVAL.node = yyDollar[1].node.append(Node("AltElseIf").append(Node("expr").append(yyDollar[4].node)).append(Node("stmt").append(yyDollar[7].node)))
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:310
		{
			yyVAL.node = yyDollar[1].node
		}
	case 108:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:312
		{
			yyVAL.node = yyDollar[1].node.append(Node("AltElse").append(Node("stmt").append(yyDollar[4].node)))
		}
	case 109:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:318
		{
			yyVAL.node = yyDollar[1].node.attribute("name", yyDollar[3].token)
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:319
		{
			yyVAL.node = Node("Class").attribute("name", yyDollar[2].token)
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:322
		{
			yyVAL.node = Node("Class").attribute(yyDollar[1].value, "true")
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:323
		{
			yyVAL.node = yyDollar[1].node.attribute(yyDollar[2].value, "true")
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:327
		{
			yyVAL.value = "abstract"
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:328
		{
			yyVAL.value = "final"
		}
	case 115:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line parser.y:333
		{
			yyVAL.node = Node("Function").
				attribute("name", yyDollar[3].token).
				attribute("returns_ref", yyDollar[2].value).
				append(yyDollar[5].node).
				append(yyDollar[7].node).
				append(yyDollar[9].node)
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:344
		{
			yyVAL.node = yyDollar[1].node
		}
	case 117:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:345
		{
			yyVAL.node = Node("Parameter list")
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:348
		{
			yyVAL.node = Node("Parameter list").append(yyDollar[1].node)
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:349
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[3].node)
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:353
		{
			yyVAL.node = Node("Parameter").
				append(yyDollar[1].node).
				attribute("is_reference", yyDollar[2].value).
				attribute("is_variadic", yyDollar[3].value).
				attribute("var", yyDollar[4].token)
		}
	case 121:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:361
		{
			yyVAL.node = Node("Parameter").
				append(yyDollar[1].node).
				attribute("is_reference", yyDollar[2].value).
				attribute("is_variadic", yyDollar[3].value).
				attribute("var", yyDollar[4].token).
				append(yyDollar[6].node)
		}
	case 122:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:372
		{
			yyVAL.node = Node("No type")
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:373
		{
			yyVAL.node = yyDollar[1].node
		}
	case 124:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:377
		{
			yyVAL.value = "false"
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:378
		{
			yyVAL.value = "true"
		}
	case 126:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:382
		{
			yyVAL.value = "false"
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:383
		{
			yyVAL.value = "true"
		}
	case 128:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:387
		{
			yyVAL.value = "false"
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:388
		{
			yyVAL.value = "true"
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:392
		{
			yyVAL.node = yyDollar[1].node
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:393
		{
			yyVAL.node = yyDollar[2].node
			yyVAL.node.attribute("nullable", "true")
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:397
		{
			yyVAL.node = yyDollar[1].node
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:398
		{
			yyVAL.node = Node("array type")
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:399
		{
			yyVAL.node = Node("callable type")
		}
	case 135:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:403
		{
			yyVAL.node = Node("void")
		}
	case 136:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:404
		{
			yyVAL.node = yyDollar[2].node
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:408
		{
			yyVAL.node = Node("Assign").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 138:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:409
		{
			yyVAL.node = Node("AssignRef").append(yyDollar[1].node).append(yyDollar[4].node)
		}
	case 139:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:410
		{
			yyVAL.node = Node("AssignRef").append(yyDollar[1].node).append(yyDollar[4].node)
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:411
		{
			yyVAL.node = Node("Clone").append(yyDollar[2].node)
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:412
		{
			yyVAL.node = Node("AssignAdd").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:413
		{
			yyVAL.node = Node("AssignSub").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:414
		{
			yyVAL.node = Node("AssignMul").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:415
		{
			yyVAL.node = Node("AssignPow").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:416
		{
			yyVAL.node = Node("AssignDiv").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:417
		{
			yyVAL.node = Node("AssignConcat").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:418
		{
			yyVAL.node = Node("AssignMod").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:419
		{
			yyVAL.node = Node("AssignAnd").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:420
		{
			yyVAL.node = Node("AssignAnd").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:421
		{
			yyVAL.node = Node("AssignOr").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:422
		{
			yyVAL.node = Node("AssignXor").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:423
		{
			yyVAL.node = Node("AssignShiftLeft").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:424
		{
			yyVAL.node = Node("AssignShiftRight").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:425
		{
			yyVAL.node = Node("PostIncrement").append(yyDollar[1].node)
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:426
		{
			yyVAL.node = Node("PreIncrement").append(yyDollar[2].node)
		}
	case 156:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:427
		{
			yyVAL.node = Node("PostDecrement").append(yyDollar[1].node)
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:428
		{
			yyVAL.node = Node("PreDecrement").append(yyDollar[2].node)
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:429
		{
			yyVAL.node = Node("Or").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:430
		{
			yyVAL.node = Node("And").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:431
		{
			yyVAL.node = Node("Or").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:432
		{
			yyVAL.node = Node("And").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:433
		{
			yyVAL.node = Node("Xor").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:434
		{
			yyVAL.node = Node("BitwiseOr").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:435
		{
			yyVAL.node = Node("BitwiseAnd").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:436
		{
			yyVAL.node = Node("BitwiseXor").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:437
		{
			yyVAL.node = Node("Concat").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:438
		{
			yyVAL.node = Node("Add").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:439
		{
			yyVAL.node = Node("Sub").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:440
		{
			yyVAL.node = Node("Mul").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:441
		{
			yyVAL.node = Node("Pow").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:442
		{
			yyVAL.node = Node("Div").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:443
		{
			yyVAL.node = Node("Mod").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:444
		{
			yyVAL.node = Node("ShiftLeft").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:445
		{
			yyVAL.node = Node("ShiftRight").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 175:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:446
		{
			yyVAL.node = Node("UnaryPlus").append(yyDollar[2].node)
		}
	case 176:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:447
		{
			yyVAL.node = Node("UnaryMinus").append(yyDollar[2].node)
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:448
		{
			yyVAL.node = Node("BooleanNot").append(yyDollar[2].node)
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:449
		{
			yyVAL.node = Node("BitwiseNot").append(yyDollar[2].node)
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:450
		{
			yyVAL.node = Node("Identical").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:451
		{
			yyVAL.node = Node("NotIdentical").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:452
		{
			yyVAL.node = Node("Equal").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:453
		{
			yyVAL.node = Node("NotEqual").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:454
		{
			yyVAL.node = Node("Spaceship").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:455
		{
			yyVAL.node = Node("Smaller").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:456
		{
			yyVAL.node = Node("SmallerOrEqual").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:457
		{
			yyVAL.node = Node("Greater").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:458
		{
			yyVAL.node = Node("GreaterOrEqual").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:459
		{
			yyVAL.node = yyDollar[2].node
		}
	case 189:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:460
		{
			yyVAL.node = Node("Ternary").append(yyDollar[1].node).append(yyDollar[3].node).append(yyDollar[5].node)
		}
	case 190:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:461
		{
			yyVAL.node = Node("Ternary").append(yyDollar[1].node).append(yyDollar[4].node)
		}
	case 191:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:462
		{
			yyVAL.node = Node("Coalesce").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 192:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:463
		{
			yyVAL.node = Node("Empty").append(yyDollar[3].node)
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:464
		{
			yyVAL.node = Node("Include").append(yyDollar[2].node)
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:465
		{
			yyVAL.node = Node("IncludeOnce").append(yyDollar[2].node)
		}
	case 195:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:466
		{
			yyVAL.node = Node("Eval").append(yyDollar[3].node)
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:467
		{
			yyVAL.node = Node("Require").append(yyDollar[2].node)
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:468
		{
			yyVAL.node = Node("RequireOnce").append(yyDollar[2].node)
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:469
		{
			yyVAL.node = Node("CastInt").append(yyDollar[2].node)
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:470
		{
			yyVAL.node = Node("CastDouble").append(yyDollar[2].node)
		}
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:471
		{
			yyVAL.node = Node("CastString").append(yyDollar[2].node)
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:472
		{
			yyVAL.node = Node("CastArray").append(yyDollar[2].node)
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:473
		{
			yyVAL.node = Node("CastObject").append(yyDollar[2].node)
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:474
		{
			yyVAL.node = Node("CastBool").append(yyDollar[2].node)
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:475
		{
			yyVAL.node = Node("CastUnset").append(yyDollar[2].node)
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:476
		{
			yyVAL.node = Node("Silence").append(yyDollar[2].node)
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:478
		{
			yyVAL.node = Node("Print").append(yyDollar[2].node)
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:479
		{
			yyVAL.node = Node("Yield")
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:480
		{
			yyVAL.node = Node("Yield").append(yyDollar[2].node)
		}
	case 209:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:481
		{
			yyVAL.node = Node("Yield").append(yyDollar[2].node).append(yyDollar[4].node)
		}
	case 210:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:482
		{
			yyVAL.node = Node("YieldFrom").append(yyDollar[2].node)
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:486
		{
			yyVAL.node = yyDollar[1].node
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:487
		{
			yyVAL.node = yyDollar[1].node
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:491
		{
			yyVAL.node = yyDollar[1].node
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:495
		{
			yyVAL.node = yyDollar[1].node
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:499
		{
			yyVAL.node = Node("Variable").attribute("name", yyDollar[1].token)
		}
	case 216:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:500
		{
			yyVAL.node = yyDollar[3].node
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:501
		{
			yyVAL.node = Node("Variable").append(yyDollar[2].node)
		}
	}
	goto yystack /* stack new state and value */
}
