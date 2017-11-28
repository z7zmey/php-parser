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

//line parser.y:574
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
	-1, 78,
	159, 46,
	-2, 230,
	-1, 265,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 202,
	-1, 266,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 203,
	-1, 267,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 204,
	-1, 268,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 205,
	-1, 269,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 206,
	-1, 270,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 207,
	-1, 271,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 208,
	-1, 272,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 209,
	-1, 273,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 210,
	-1, 313,
	163, 140,
	-2, 145,
}

const yyPrivate = 57344

const yyLast = 2592

var yyAct = [...]int{

	19, 140, 339, 391, 345, 357, 337, 60, 241, 389,
	340, 361, 137, 376, 377, 376, 377, 147, 147, 147,
	313, 307, 282, 280, 215, 236, 213, 204, 4, 145,
	57, 205, 144, 57, 208, 209, 210, 211, 212, 142,
	214, 141, 216, 217, 218, 219, 220, 221, 222, 223,
	224, 225, 226, 227, 228, 348, 333, 315, 139, 397,
	153, 155, 154, 314, 279, 413, 394, 405, 375, 395,
	388, 205, 373, 60, 214, 392, 216, 217, 226, 227,
	177, 370, 178, 151, 152, 156, 158, 157, 170, 171,
	168, 169, 172, 173, 174, 175, 176, 166, 167, 160,
	161, 159, 162, 164, 165, 229, 58, 343, 344, 58,
	360, 358, 356, 308, 283, 247, 246, 245, 163, 232,
	231, 385, 367, 233, 240, 399, 347, 139, 305, 346,
	54, 55, 278, 396, 376, 377, 379, 376, 377, 163,
	277, 183, 239, 364, 60, 243, 244, 186, 188, 187,
	184, 185, 248, 249, 250, 251, 252, 253, 254, 255,
	256, 257, 258, 259, 260, 261, 262, 263, 264, 265,
	266, 267, 268, 269, 270, 271, 272, 273, 274, 276,
	180, 181, 382, 369, 284, 332, 407, 362, 25, 309,
	285, 287, 288, 289, 290, 291, 292, 293, 294, 295,
	296, 297, 298, 393, 56, 299, 162, 164, 165, 341,
	27, 281, 62, 355, 301, 146, 302, 153, 155, 154,
	206, 207, 163, 237, 238, 63, 5, 8, 1, 331,
	304, 400, 182, 148, 149, 242, 325, 177, 11, 178,
	151, 152, 156, 158, 157, 170, 171, 168, 169, 172,
	173, 174, 175, 176, 166, 167, 160, 161, 159, 162,
	164, 165, 24, 230, 10, 139, 160, 161, 159, 162,
	164, 165, 23, 29, 26, 163, 312, 336, 335, 381,
	338, 316, 21, 317, 234, 163, 318, 319, 173, 174,
	175, 176, 166, 167, 160, 161, 159, 162, 164, 165,
	2, 342, 138, 3, 323, 59, 179, 384, 328, 243,
	330, 363, 334, 163, 343, 344, 0, 329, 0, 166,
	167, 160, 161, 159, 162, 164, 165, 0, 0, 354,
	0, 0, 326, 347, 0, 0, 346, 0, 0, 359,
	163, 0, 0, 0, 0, 0, 352, 0, 0, 0,
	0, 366, 365, 0, 0, 371, 0, 243, 0, 0,
	0, 0, 0, 0, 378, 374, 380, 0, 0, 383,
	350, 387, 386, 0, 0, 368, 0, 390, 0, 0,
	72, 73, 74, 75, 76, 398, 79, 80, 81, 77,
	78, 0, 53, 404, 403, 0, 0, 0, 0, 406,
	0, 0, 0, 0, 408, 409, 0, 0, 411, 0,
	0, 0, 0, 0, 0, 401, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 33, 34, 0, 0, 0,
	0, 35, 82, 36, 31, 32, 43, 44, 45, 46,
	47, 48, 49, 50, 0, 0, 83, 70, 0, 86,
	87, 88, 64, 65, 66, 67, 68, 69, 84, 85,
	0, 0, 61, 0, 57, 0, 0, 0, 0, 0,
	0, 89, 90, 91, 92, 93, 94, 95, 96, 97,
	98, 99, 116, 117, 118, 119, 120, 110, 111, 112,
	113, 114, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 71, 0, 128, 126, 127, 123, 124, 0,
	115, 121, 122, 129, 130, 132, 131, 133, 134, 0,
	0, 0, 0, 0, 143, 39, 40, 41, 42, 0,
	125, 136, 135, 51, 52, 0, 53, 0, 37, 0,
	58, 0, 152, 156, 158, 157, 170, 171, 168, 169,
	172, 173, 174, 175, 176, 166, 167, 160, 161, 159,
	162, 164, 165, 0, 0, 0, 0, 0, 0, 33,
	34, 0, 0, 0, 0, 35, 163, 36, 31, 32,
	43, 44, 45, 46, 47, 48, 49, 50, 0, 0,
	0, 30, 0, 0, 0, 0, 0, 54, 55, 0,
	0, 0, 0, 28, 0, 0, 0, 0, 57, 0,
	0, 0, 0, 0, 0, 0, 13, 12, 0, 14,
	412, 0, 0, 0, 0, 0, 15, 0, 0, 0,
	16, 17, 0, 20, 0, 18, 154, 0, 0, 0,
	143, 39, 40, 41, 42, 0, 38, 0, 22, 51,
	52, 0, 53, 0, 177, 0, 178, 151, 152, 156,
	158, 157, 170, 171, 168, 169, 172, 173, 174, 175,
	176, 166, 167, 160, 161, 159, 162, 164, 165, 0,
	9, 0, 37, 0, 58, 33, 34, 0, 0, 0,
	0, 35, 163, 36, 31, 32, 43, 44, 45, 46,
	47, 48, 49, 50, 0, 0, 0, 30, 0, 0,
	0, 0, 0, 54, 55, 0, 0, 0, 0, 28,
	0, 0, 0, 0, 57, 0, 0, 0, 0, 0,
	0, 0, 13, 12, 0, 14, 0, 0, 0, 0,
	0, 0, 15, 0, 0, 0, 16, 17, 0, 20,
	0, 18, 303, 0, 0, 0, 143, 39, 40, 41,
	42, 0, 38, 0, 22, 51, 52, 177, 53, 178,
	151, 152, 156, 158, 157, 170, 171, 168, 169, 172,
	173, 174, 175, 176, 166, 167, 160, 161, 159, 162,
	164, 165, 0, 0, 0, 0, 9, 410, 37, 0,
	58, 33, 34, 0, 0, 163, 0, 35, 0, 36,
	31, 32, 43, 44, 45, 46, 47, 48, 49, 50,
	0, 0, 0, 30, 0, 0, 0, 0, 0, 54,
	55, 0, 0, 0, 0, 28, 0, 0, 0, 0,
	57, 0, 0, 0, 0, 0, 0, 0, 13, 12,
	372, 14, 0, 0, 0, 0, 0, 0, 15, 0,
	0, 0, 16, 17, 0, 20, 0, 18, 0, 0,
	0, 0, 143, 39, 40, 41, 42, 0, 38, 0,
	22, 51, 52, 177, 53, 178, 151, 152, 156, 158,
	157, 170, 171, 168, 169, 172, 173, 174, 175, 176,
	166, 167, 160, 161, 159, 162, 164, 165, 0, 0,
	0, 0, 9, 0, 37, 0, 58, 33, 34, 0,
	0, 163, 0, 35, 0, 36, 31, 32, 43, 44,
	45, 46, 47, 48, 49, 50, 0, 0, 0, 30,
	0, 0, 0, 351, 0, 54, 55, 0, 0, 0,
	0, 28, 0, 0, 0, 0, 57, 0, 0, 0,
	0, 0, 0, 0, 13, 12, 0, 14, 0, 0,
	0, 0, 0, 0, 15, 0, 0, 0, 16, 17,
	0, 20, 0, 18, 0, 0, 0, 0, 143, 39,
	40, 41, 42, 0, 38, 0, 22, 51, 52, 0,
	53, 178, 151, 152, 156, 158, 157, 170, 171, 168,
	169, 172, 173, 174, 175, 176, 166, 167, 160, 161,
	159, 162, 164, 165, 0, 0, 0, 0, 9, 0,
	37, 0, 58, 33, 34, 0, 0, 163, 0, 35,
	0, 36, 31, 32, 43, 44, 45, 46, 47, 48,
	49, 50, 0, 0, 0, 30, 0, 0, 0, 0,
	0, 54, 55, 0, 0, 0, 0, 28, 0, 0,
	0, 0, 57, 0, 0, 0, 0, 0, 0, 0,
	13, 12, 0, 14, 0, 0, 0, 0, 0, 0,
	15, 0, 0, 0, 16, 17, 0, 20, 0, 18,
	0, 0, 0, 0, 6, 39, 40, 41, 42, 0,
	38, 0, 22, 51, 52, 0, 53, 0, 156, 158,
	157, 170, 171, 168, 169, 172, 173, 174, 175, 176,
	166, 167, 160, 161, 159, 162, 164, 165, 0, 0,
	0, 0, 0, 0, 9, 235, 37, 0, 58, 33,
	34, 163, 0, 0, 0, 35, 0, 36, 31, 32,
	43, 44, 45, 46, 47, 48, 49, 50, 0, 0,
	0, 30, 0, 0, 0, 0, 0, 54, 55, 0,
	0, 0, 0, 28, 0, 0, 0, 0, 57, 0,
	0, 0, 0, 0, 0, 0, 13, 12, 0, 14,
	0, 0, 0, 0, 0, 0, 15, 0, 0, 0,
	16, 17, 0, 20, 0, 18, 0, 0, 0, 0,
	143, 39, 40, 41, 42, 0, 38, 0, 22, 51,
	52, 0, 53, 0, 0, 158, 157, 170, 171, 168,
	169, 172, 173, 174, 175, 176, 166, 167, 160, 161,
	159, 162, 164, 165, 7, 0, 0, 0, 0, 0,
	9, 0, 37, 0, 58, 33, 34, 163, 0, 0,
	0, 35, 0, 36, 31, 32, 43, 44, 45, 46,
	47, 48, 49, 50, 0, 0, 0, 30, 0, 0,
	0, 0, 0, 54, 55, 0, 0, 0, 0, 28,
	0, 0, 0, 0, 57, 0, 0, 143, 39, 40,
	41, 42, 13, 12, 0, 14, 51, 52, 0, 53,
	0, 0, 15, 0, 0, 0, 16, 17, 0, 20,
	0, 18, 0, 0, 402, 0, 0, 0, 0, 0,
	0, 0, 38, 0, 22, 0, 0, 0, 0, 0,
	0, 0, 33, 34, 0, 0, 0, 0, 35, 0,
	36, 31, 32, 43, 44, 45, 46, 47, 48, 49,
	50, 0, 0, 0, 30, 0, 9, 0, 37, 0,
	58, 0, 0, 0, 0, 0, 28, 0, 0, 0,
	0, 57, 0, 0, 143, 39, 40, 41, 42, 13,
	12, 0, 14, 51, 52, 0, 53, 0, 0, 15,
	0, 0, 0, 16, 17, 0, 0, 0, 18, 0,
	0, 353, 0, 0, 0, 0, 0, 0, 0, 38,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 33,
	34, 0, 0, 0, 0, 35, 0, 36, 31, 32,
	43, 44, 45, 46, 47, 48, 49, 50, 0, 0,
	0, 30, 0, 9, 0, 37, 0, 58, 0, 0,
	0, 0, 0, 28, 0, 0, 0, 0, 57, 0,
	0, 143, 39, 40, 41, 42, 13, 12, 0, 14,
	51, 52, 0, 53, 0, 0, 15, 0, 0, 0,
	16, 17, 0, 0, 0, 18, 0, 0, 327, 0,
	0, 0, 0, 0, 0, 0, 38, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 33, 34, 0, 0,
	0, 0, 35, 0, 36, 31, 32, 43, 44, 45,
	46, 47, 48, 49, 50, 0, 0, 0, 30, 0,
	9, 0, 37, 0, 58, 0, 0, 0, 0, 0,
	28, 0, 0, 0, 0, 57, 0, 0, 143, 39,
	40, 41, 42, 13, 12, 0, 14, 51, 52, 0,
	53, 0, 0, 15, 0, 0, 0, 16, 17, 0,
	0, 0, 18, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 38, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 33, 34, 0, 0, 0, 0, 35,
	0, 36, 31, 32, 43, 44, 45, 46, 47, 48,
	49, 50, 0, 0, 0, 30, 0, 9, 0, 37,
	0, 58, 0, 0, 0, 0, 0, 28, 0, 0,
	0, 0, 57, 0, 0, 0, 0, 0, 0, 0,
	13, 12, 0, 14, 0, 0, 0, 0, 0, 0,
	15, 153, 155, 154, 16, 17, 0, 0, 0, 18,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	38, 177, 0, 178, 151, 152, 156, 158, 157, 170,
	171, 168, 169, 172, 173, 174, 175, 176, 166, 167,
	160, 161, 159, 162, 164, 165, 153, 155, 154, 0,
	0, 0, 0, 0, 9, 0, 37, 0, 58, 163,
	0, 0, 0, 0, 0, 0, 177, 0, 178, 151,
	152, 156, 158, 157, 170, 171, 168, 169, 172, 173,
	174, 175, 176, 166, 167, 160, 161, 159, 162, 164,
	165, 153, 155, 154, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 163, 0, 0, 0, 0, 0,
	0, 177, 0, 178, 151, 152, 156, 158, 157, 170,
	171, 168, 169, 172, 173, 174, 175, 176, 166, 167,
	160, 161, 159, 162, 164, 165, 153, 155, 154, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 163,
	0, 0, 0, 0, 349, 0, 177, 0, 178, 151,
	152, 156, 158, 157, 170, 171, 168, 169, 172, 173,
	174, 175, 176, 166, 167, 160, 161, 159, 162, 164,
	165, 153, 155, 154, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 163, 0, 0, 0, 0, 322,
	0, 177, 0, 178, 151, 152, 156, 158, 157, 170,
	171, 168, 169, 172, 173, 174, 175, 176, 166, 167,
	160, 161, 159, 162, 164, 165, 153, 155, 154, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 163,
	0, 0, 0, 0, 321, 0, 177, 0, 178, 151,
	152, 156, 158, 157, 170, 171, 168, 169, 172, 173,
	174, 175, 176, 166, 167, 160, 161, 159, 162, 164,
	165, 153, 155, 154, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 163, 0, 0, 0, 0, 320,
	0, 177, 0, 178, 151, 152, 156, 158, 157, 170,
	171, 168, 169, 172, 173, 174, 175, 176, 166, 167,
	160, 161, 159, 162, 164, 165, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 163,
	0, 0, 0, 0, 310, 143, 39, 40, 41, 42,
	0, 0, 0, 0, 51, 52, 0, 53, 0, 0,
	0, 0, 0, 0, 0, 143, 39, 40, 41, 42,
	0, 0, 0, 0, 51, 52, 0, 53, 286, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 306,
	33, 34, 275, 0, 0, 0, 35, 0, 36, 31,
	32, 43, 44, 45, 46, 47, 48, 49, 50, 0,
	33, 34, 30, 0, 0, 0, 35, 0, 36, 31,
	32, 43, 44, 45, 46, 47, 48, 49, 50, 57,
	0, 0, 30, 0, 300, 143, 39, 40, 41, 42,
	0, 0, 0, 0, 51, 52, 0, 53, 0, 57,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 38, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	33, 34, 0, 0, 0, 0, 35, 38, 36, 31,
	32, 43, 44, 45, 46, 47, 48, 49, 50, 0,
	0, 0, 30, 37, 0, 58, 153, 155, 154, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 57,
	0, 0, 0, 37, 0, 58, 177, 0, 178, 151,
	152, 156, 158, 157, 170, 171, 168, 169, 172, 173,
	174, 175, 176, 166, 167, 160, 161, 159, 162, 164,
	165, 0, 0, 0, 0, 0, 0, 38, 0, 0,
	0, 0, 0, 0, 163, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 153, 155, 154, 0,
	189, 190, 191, 192, 194, 195, 196, 197, 198, 199,
	200, 201, 193, 37, 0, 58, 177, 392, 178, 151,
	152, 156, 158, 157, 170, 171, 168, 169, 172, 173,
	174, 175, 176, 166, 167, 160, 161, 159, 162, 164,
	165, 202, 203, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 163, 0, 153, 155, 154, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 177, 324, 178, 151,
	152, 156, 158, 157, 170, 171, 168, 169, 172, 173,
	174, 175, 176, 166, 167, 160, 161, 159, 162, 164,
	165, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 163, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 153, 155, 154, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 393, 177, 311, 178, 151,
	152, 156, 158, 157, 170, 171, 168, 169, 172, 173,
	174, 175, 176, 166, 167, 160, 161, 159, 162, 164,
	165, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 163, 0, 153, 155, 154, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 150, 177, 0, 178, 151,
	152, 156, 158, 157, 170, 171, 168, 169, 172, 173,
	174, 175, 176, 166, 167, 160, 161, 159, 162, 164,
	165, 155, 154, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 163, 0, 0, 0, 0, 0,
	177, 0, 178, 151, 152, 156, 158, 157, 170, 171,
	168, 169, 172, 173, 174, 175, 176, 166, 167, 160,
	161, 159, 162, 164, 165, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 163, 157,
	170, 171, 168, 169, 172, 173, 174, 175, 176, 166,
	167, 160, 161, 159, 162, 164, 165, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	163, 170, 171, 168, 169, 172, 173, 174, 175, 176,
	166, 167, 160, 161, 159, 162, 164, 165, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 163,
}
var yyPact = [...]int{

	-1000, -1000, 1100, -1000, -1000, -1000, 376, 41, -1000, -1000,
	-1000, -1000, -121, 1564, -130, -133, 2091, 2091, 2091, 2286,
	143, 53, 55, 77, 74, 2223, -1000, -1000, -135, -1000,
	2091, -58, -58, 2091, 2091, 2091, 2091, 2091, -136, 2091,
	-138, 2091, 2091, 2091, 2091, 2091, 2091, 2091, 2091, 2091,
	2091, 2091, 2091, 2091, -1000, -1000, -1000, -1000, -55, -39,
	2416, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	2091, -136, 2091, 2091, -138, 2091, 2091, 2091, 2091, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -40, -34, -1000,
	984, 2091, 27, 2091, 2091, 2091, -42, 2416, -43, -44,
	-1000, 2091, 2091, 2091, 2091, 2091, 2091, 2091, 2091, 2091,
	2091, 2091, 2091, 2091, 2091, 2091, 2091, 2091, 2091, 2091,
	2091, 2091, 2091, 2091, 2091, 2091, 2091, 2021, 2091, 54,
	-1000, 46, -1000, -96, -139, 1564, -140, -45, 153, 2001,
	2091, 2091, 2091, 2091, 2091, 2091, 2091, 2091, 2091, 2091,
	2091, 2091, -1000, -1000, 2091, -1000, -1000, -1000, 71, 71,
	71, 71, 1931, 2091, 2416, 2091, 2416, 2416, 71, 71,
	71, 71, 71, 71, 71, 71, 853, 737, 853, 2091,
	-1000, -1000, -1000, 42, -1000, -1000, -1000, -1000, -1000, 1886,
	-141, -46, 180, 2416, 1841, -1000, -1000, -1000, 508, 1083,
	2460, 853, 624, 1199, 2523, 2492, 154, 154, 154, 71,
	71, 71, 71, 217, 217, 245, 245, 245, 245, 245,
	272, 272, 272, 272, 2356, 2091, 969, -142, -97, -104,
	2091, -1000, 2091, -1000, -1000, 853, 2091, 853, 853, 853,
	853, 853, 853, 853, 853, 853, 853, 853, 853, 1796,
	-1000, 1751, 1706, 2091, 2156, -1000, 1477, 2091, 2091, 2091,
	25, 2091, 969, 179, -106, -1000, 1661, 207, 868, 2523,
	1390, -1000, -1000, 853, -1000, -1000, -1000, -1000, 50, -47,
	2416, -1000, -48, -49, 969, -152, 178, -1000, 106, -1000,
	-1000, -28, -1000, -1000, -1000, -1000, 41, -35, -1000, 1564,
	152, -78, -1000, -1000, 752, -87, 2091, -93, -1000, 29,
	-1000, 151, 179, -37, -1000, -1000, -1000, 41, -1000, -1000,
	-1000, 1216, -89, -1000, -154, -1000, 2091, 44, -95, -90,
	26, -101, 179, -1000, 37, -1000, -1000, 1216, -1000, 1303,
	2226, -1000, -1000, -1000, -1000, -1000, -92, -1000, -1000, 169,
	-1000, -1000, -1000, -1000, 1216, -1000, 636, 2091, 520, 1216,
	-1000, 2416, -94, -1000,
}
var yyPgo = [...]int{

	0, 210, 311, 307, 306, 305, 303, 4, 302, 301,
	300, 25, 284, 1, 282, 224, 223, 280, 279, 2,
	10, 278, 277, 6, 0, 274, 273, 188, 204, 272,
	264, 262, 238, 236, 8, 235, 231, 229, 5, 215,
	228, 225, 212, 3,
}
var yyR1 = [...]int{

	0, 40, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 42,
	42, 42, 42, 42, 42, 42, 5, 5, 8, 8,
	7, 9, 9, 9, 10, 10, 6, 6, 6, 6,
	6, 13, 13, 12, 12, 12, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 29, 29, 30,
	30, 31, 31, 32, 32, 33, 33, 34, 34, 35,
	35, 37, 37, 37, 37, 38, 38, 38, 43, 43,
	36, 36, 15, 15, 14, 14, 1, 1, 16, 21,
	21, 22, 22, 23, 23, 17, 17, 4, 4, 2,
	2, 3, 3, 19, 19, 20, 20, 20, 18, 18,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 24, 24, 39, 39, 26, 27,
	28, 28, 28,
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
	1, 2, 0, 1, 1, 1, 3, 1, 1, 5,
	7, 9, 5, 3, 3, 3, 2, 5, 6, 1,
	3, 6, 7, 3, 6, 1, 4, 0, 1, 3,
	1, 3, 4, 4, 5, 0, 5, 4, 1, 1,
	1, 4, 5, 4, 1, 2, 1, 1, 10, 1,
	0, 1, 3, 4, 6, 0, 1, 0, 1, 0,
	1, 0, 1, 1, 2, 1, 1, 1, 0, 2,
	3, 4, 4, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 4, 3, 4, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	1, 2, 4, 2, 1, 1, 0, 1, 1, 1,
	1, 4, 2,
}
var yyChk = [...]int{

	-1000, -40, -10, -6, -11, -16, 4, 154, -15, 160,
	-30, -32, 97, 96, 99, 106, 110, 111, 115, -24,
	113, -14, 128, -29, -31, -27, -25, -1, 83, -26,
	71, 58, 59, 49, 50, 55, 57, 162, 126, 5,
	6, 7, 8, 60, 61, 62, 63, 64, 65, 66,
	67, 13, 14, 16, 77, 78, -28, 88, 164, -5,
	-24, 86, -42, -41, 76, 77, 78, 79, 80, 81,
	71, 126, 4, 5, 6, 7, 8, 13, 14, 10,
	11, 12, 56, 70, 82, 83, 73, 74, 75, 95,
	96, 97, 98, 99, 100, 101, 102, 103, 104, 105,
	116, 117, 118, 119, 120, 121, 122, 123, 124, 125,
	111, 112, 113, 114, 115, 134, 106, 107, 108, 109,
	110, 135, 136, 131, 132, 154, 129, 130, 128, 137,
	138, 140, 139, 141, 142, 156, 155, -7, -8, 86,
	-13, 162, -11, 4, 162, 162, -39, -24, -39, -39,
	159, 33, 34, 10, 12, 11, 35, 37, 36, 51,
	49, 50, 52, 68, 53, 54, 47, 48, 40, 41,
	38, 39, 42, 43, 44, 45, 46, 30, 32, -4,
	37, 128, -1, 86, 73, 74, 73, 75, 74, 17,
	18, 19, 20, 29, 21, 22, 23, 24, 25, 26,
	27, 28, 58, 59, 162, -24, -27, -27, -24, -24,
	-24, -24, -24, 162, -24, 162, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, 160,
	-28, 159, 159, 157, -12, 161, -11, -16, -15, -24,
	97, -34, -35, -24, -24, 159, 159, 159, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, 31, -24, 86, 86, 160,
	162, -11, 162, 159, 31, -24, 37, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	163, -24, -24, 15, -24, 86, 163, 162, 159, 9,
	163, 31, -24, 162, 160, 161, -24, -24, -13, -24,
	163, 163, 163, -24, 161, -33, -11, 31, -24, -34,
	-24, -37, 160, 31, -24, -21, -22, -23, -17, -19,
	-20, 30, -9, 135, 136, -7, 157, 154, 161, 163,
	163, 75, -11, 31, -13, 163, 159, -38, 159, -38,
	159, 163, 9, -2, 37, -20, -7, 157, -11, 31,
	159, -13, 98, 159, -34, 161, 108, 109, -38, 107,
	-38, -18, 31, -23, -3, 158, -7, -13, 159, 163,
	-24, -43, 31, 159, 161, 159, 107, 160, -19, 88,
	-36, -11, 31, -43, -13, 159, -13, 17, -13, -13,
	161, -24, 100, 159,
}
var yyDef = [...]int{

	85, -2, 1, 84, 86, 87, 0, 0, 90, 92,
	97, 98, 0, 0, 0, 0, 236, 236, 236, 0,
	147, 0, 0, 109, 0, 234, 235, 134, 0, 239,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 230, 0, 136, 137, 238, 240, 0, 0,
	216, 76, 77, 69, 70, 71, 72, 73, 74, 75,
	12, 39, 2, 3, 4, 5, 6, 45, -2, 7,
	8, 9, 10, 11, 13, 14, 15, 16, 17, 18,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
	40, 41, 42, 43, 44, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 0, 80, 78,
	0, 0, 0, 0, 117, 0, 0, 237, 0, 0,
	106, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	148, 0, 135, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 177, 179, 0, 163, 178, 180, 198, 199,
	200, 201, 0, 0, 217, 0, 219, 220, 221, 222,
	223, 224, 225, 226, 227, 228, 229, 231, 233, 0,
	242, 88, 89, 0, 91, 96, 93, 94, 95, 0,
	0, 0, 118, 120, 0, 103, 104, 105, 181, 182,
	183, 184, 185, 186, 187, 188, 189, 190, 191, 192,
	193, 194, 195, 196, 197, -2, -2, -2, -2, -2,
	-2, -2, -2, -2, 0, 0, 214, 0, 0, 0,
	0, 110, 0, 113, 92, 160, 0, 164, 165, 166,
	167, 168, 169, 170, 171, 173, 174, 175, 176, 0,
	211, 0, 0, 0, 0, 79, 0, 0, 117, 0,
	0, 0, 213, -2, 0, 133, 0, 0, 0, 161,
	0, 215, 218, 232, 241, 99, 115, 92, 0, 0,
	119, 102, 125, 125, 212, 0, 139, 141, 149, 146,
	153, 0, 155, 156, 157, 81, 0, 0, 132, 0,
	0, 0, 107, 92, 0, 0, 117, 0, 125, 0,
	125, 158, 145, 151, 150, 154, 82, 0, 108, 92,
	114, 111, 0, 100, 0, 121, 0, 0, 0, 0,
	0, 0, 0, 142, 0, 152, 83, 112, 116, 0,
	0, 92, 128, 129, 122, 123, 0, 92, 159, 143,
	101, 130, 92, 92, 127, 124, 0, 0, 0, 126,
	138, 144, 0, 131,
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
		//line parser.y:218
		{
			fmt.Println(yyDollar[1].node)
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:237
		{
			yyVAL.node = Node("identifier")
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:238
		{
			yyVAL.node = Node("reserved")
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:242
		{
			yyVAL.node = Node("NamespaceParts").append(Node(yyDollar[1].token))
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:243
		{
			yyVAL.node = yyDollar[1].node.append(Node(yyDollar[3].token))
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:247
		{
			yyVAL.node = yyDollar[1].node
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:251
		{
			yyVAL.node = Node("Name").append(yyDollar[1].node)
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:252
		{
			yyVAL.node = Node("Name").append(yyDollar[2].node).attribute("FullyQualified", "true")
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:253
		{
			yyVAL.node = Node("Name").append(yyDollar[3].node).attribute("Relative", "true")
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:257
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[2].node)
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:258
		{
			yyVAL.node = Node("Statements")
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:262
		{
			yyVAL.node = yyDollar[1].node
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:263
		{
			yyVAL.node = yyDollar[1].node
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:264
		{
			yyVAL.node = yyDollar[2].node /*TODO: identifier stub, refactor it*/
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:265
		{
			yyVAL.node = Node("Namespace").append(yyDollar[2].node)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:266
		{
			yyVAL.node = yyDollar[1].node
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:270
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[2].node)
		}
	case 92:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:271
		{
			yyVAL.node = Node("stmt")
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:275
		{
			yyVAL.node = yyDollar[1].node
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:276
		{
			yyVAL.node = yyDollar[1].node
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:277
		{
			yyVAL.node = yyDollar[1].node
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:280
		{
			yyVAL.node = yyDollar[2].node
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:281
		{
			yyVAL.node = yyDollar[1].node
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:282
		{
			yyVAL.node = yyDollar[1].node
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:283
		{
			yyVAL.node = Node("While").append(Node("expr").append(yyDollar[3].node)).append(Node("stmt").append(yyDollar[5].node))
		}
	case 100:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:284
		{
			yyVAL.node = Node("DoWhile").append(Node("expr").append(yyDollar[5].node)).append(Node("stmt").append(yyDollar[2].node))
		}
	case 101:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.y:286
		{
			yyVAL.node = Node("For").
				append(Node("expr1").append(yyDollar[3].node)).
				append(Node("expr2").append(yyDollar[5].node)).
				append(Node("expr3").append(yyDollar[7].node)).
				append(Node("stmt").append(yyDollar[9].node))
		}
	case 102:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:293
		{
			yyVAL.node = Node("Switch").append(Node("expr").append(yyDollar[3].node)).append(yyDollar[5].node)
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:294
		{
			yyVAL.node = Node("Break").append(yyDollar[2].node)
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:295
		{
			yyVAL.node = Node("Continue").append(yyDollar[2].node)
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:296
		{
			yyVAL.node = Node("Return").append(yyDollar[2].node)
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:297
		{
			yyVAL.node = yyDollar[1].node
		}
	case 107:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:301
		{
			yyVAL.node = Node("If").append(Node("expr").append(yyDollar[3].node)).append(Node("stmt").append(yyDollar[5].node))
		}
	case 108:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:305
		{
			yyVAL.node = yyDollar[1].node.append(Node("ElseIf").append(Node("expr").append(yyDollar[4].node)).append(Node("stmt").append(yyDollar[6].node)))
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:311
		{
			yyVAL.node = yyDollar[1].node
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:313
		{
			yyVAL.node = yyDollar[1].node.append(Node("Else").append(Node("stmt").append(yyDollar[3].node)))
		}
	case 111:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:320
		{
			yyVAL.node = Node("AltIf").append(Node("expr").append(yyDollar[3].node)).append(Node("stmt").append(yyDollar[6].node))
		}
	case 112:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:324
		{
			yyVAL.node = yyDollar[1].node.append(Node("AltElseIf").append(Node("expr").append(yyDollar[4].node)).append(Node("stmt").append(yyDollar[7].node)))
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:330
		{
			yyVAL.node = yyDollar[1].node
		}
	case 114:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:332
		{
			yyVAL.node = yyDollar[1].node.append(Node("AltElse").append(Node("stmt").append(yyDollar[4].node)))
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:338
		{
			yyVAL.node = yyDollar[1].node
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:339
		{
			yyVAL.node = yyDollar[2].node
		}
	case 117:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:343
		{
			yyVAL.node = Node("null")
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:344
		{
			yyVAL.node = yyDollar[1].node
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:347
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[3].node)
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:348
		{
			yyVAL.node = Node("ExpressionList").append(yyDollar[1].node)
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:352
		{
			yyVAL.node = yyDollar[2].node
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:353
		{
			yyVAL.node = yyDollar[3].node
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:354
		{
			yyVAL.node = yyDollar[2].node
		}
	case 124:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:355
		{
			yyVAL.node = yyDollar[3].node
		}
	case 125:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:359
		{
			yyVAL.node = Node("CaseList")
		}
	case 126:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:361
		{
			yyVAL.node = yyDollar[1].node.append(Node("Case").append(Node("expr").append(yyDollar[3].node)).append(yyDollar[5].node))
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:365
		{
			yyVAL.node = yyDollar[1].node.append(Node("Default").append(yyDollar[4].node))
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:376
		{
			yyVAL.node = yyDollar[1].node
		}
	case 131:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:377
		{
			yyVAL.node = yyDollar[2].node
		}
	case 132:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:381
		{
			yyVAL.node = yyDollar[1].node.attribute("name", yyDollar[3].token)
		}
	case 133:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:382
		{
			yyVAL.node = Node("Class").attribute("name", yyDollar[2].token)
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:385
		{
			yyVAL.node = Node("Class").attribute(yyDollar[1].value, "true")
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:386
		{
			yyVAL.node = yyDollar[1].node.attribute(yyDollar[2].value, "true")
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:390
		{
			yyVAL.value = "abstract"
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:391
		{
			yyVAL.value = "final"
		}
	case 138:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line parser.y:396
		{
			yyVAL.node = Node("Function").
				attribute("name", yyDollar[3].token).
				attribute("returns_ref", yyDollar[2].value).
				append(yyDollar[5].node).
				append(yyDollar[7].node).
				append(yyDollar[9].node)
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:407
		{
			yyVAL.node = yyDollar[1].node
		}
	case 140:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:408
		{
			yyVAL.node = Node("Parameter list")
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:411
		{
			yyVAL.node = Node("Parameter list").append(yyDollar[1].node)
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:412
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[3].node)
		}
	case 143:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:416
		{
			yyVAL.node = Node("Parameter").
				append(yyDollar[1].node).
				attribute("is_reference", yyDollar[2].value).
				attribute("is_variadic", yyDollar[3].value).
				attribute("var", yyDollar[4].token)
		}
	case 144:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:424
		{
			yyVAL.node = Node("Parameter").
				append(yyDollar[1].node).
				attribute("is_reference", yyDollar[2].value).
				attribute("is_variadic", yyDollar[3].value).
				attribute("var", yyDollar[4].token).
				append(yyDollar[6].node)
		}
	case 145:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:435
		{
			yyVAL.node = Node("No type")
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:436
		{
			yyVAL.node = yyDollar[1].node
		}
	case 147:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:440
		{
			yyVAL.value = "false"
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:441
		{
			yyVAL.value = "true"
		}
	case 149:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:445
		{
			yyVAL.value = "false"
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:446
		{
			yyVAL.value = "true"
		}
	case 151:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:450
		{
			yyVAL.value = "false"
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:451
		{
			yyVAL.value = "true"
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:455
		{
			yyVAL.node = yyDollar[1].node
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:456
		{
			yyVAL.node = yyDollar[2].node
			yyVAL.node.attribute("nullable", "true")
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:460
		{
			yyVAL.node = yyDollar[1].node
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:461
		{
			yyVAL.node = Node("array type")
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:462
		{
			yyVAL.node = Node("callable type")
		}
	case 158:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:466
		{
			yyVAL.node = Node("void")
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:467
		{
			yyVAL.node = yyDollar[2].node
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:471
		{
			yyVAL.node = Node("Assign").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 161:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:472
		{
			yyVAL.node = Node("AssignRef").append(yyDollar[1].node).append(yyDollar[4].node)
		}
	case 162:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:473
		{
			yyVAL.node = Node("AssignRef").append(yyDollar[1].node).append(yyDollar[4].node)
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:474
		{
			yyVAL.node = Node("Clone").append(yyDollar[2].node)
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:475
		{
			yyVAL.node = Node("AssignAdd").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:476
		{
			yyVAL.node = Node("AssignSub").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:477
		{
			yyVAL.node = Node("AssignMul").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:478
		{
			yyVAL.node = Node("AssignPow").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:479
		{
			yyVAL.node = Node("AssignDiv").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:480
		{
			yyVAL.node = Node("AssignConcat").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:481
		{
			yyVAL.node = Node("AssignMod").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:482
		{
			yyVAL.node = Node("AssignAnd").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:483
		{
			yyVAL.node = Node("AssignAnd").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:484
		{
			yyVAL.node = Node("AssignOr").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:485
		{
			yyVAL.node = Node("AssignXor").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:486
		{
			yyVAL.node = Node("AssignShiftLeft").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:487
		{
			yyVAL.node = Node("AssignShiftRight").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:488
		{
			yyVAL.node = Node("PostIncrement").append(yyDollar[1].node)
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:489
		{
			yyVAL.node = Node("PreIncrement").append(yyDollar[2].node)
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:490
		{
			yyVAL.node = Node("PostDecrement").append(yyDollar[1].node)
		}
	case 180:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:491
		{
			yyVAL.node = Node("PreDecrement").append(yyDollar[2].node)
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:492
		{
			yyVAL.node = Node("Or").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:493
		{
			yyVAL.node = Node("And").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:494
		{
			yyVAL.node = Node("Or").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:495
		{
			yyVAL.node = Node("And").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:496
		{
			yyVAL.node = Node("Xor").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:497
		{
			yyVAL.node = Node("BitwiseOr").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:498
		{
			yyVAL.node = Node("BitwiseAnd").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:499
		{
			yyVAL.node = Node("BitwiseXor").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 189:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:500
		{
			yyVAL.node = Node("Concat").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:501
		{
			yyVAL.node = Node("Add").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 191:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:502
		{
			yyVAL.node = Node("Sub").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:503
		{
			yyVAL.node = Node("Mul").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 193:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:504
		{
			yyVAL.node = Node("Pow").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 194:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:505
		{
			yyVAL.node = Node("Div").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:506
		{
			yyVAL.node = Node("Mod").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:507
		{
			yyVAL.node = Node("ShiftLeft").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:508
		{
			yyVAL.node = Node("ShiftRight").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:509
		{
			yyVAL.node = Node("UnaryPlus").append(yyDollar[2].node)
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:510
		{
			yyVAL.node = Node("UnaryMinus").append(yyDollar[2].node)
		}
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:511
		{
			yyVAL.node = Node("BooleanNot").append(yyDollar[2].node)
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:512
		{
			yyVAL.node = Node("BitwiseNot").append(yyDollar[2].node)
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:513
		{
			yyVAL.node = Node("Identical").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:514
		{
			yyVAL.node = Node("NotIdentical").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:515
		{
			yyVAL.node = Node("Equal").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:516
		{
			yyVAL.node = Node("NotEqual").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 206:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:517
		{
			yyVAL.node = Node("Spaceship").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:518
		{
			yyVAL.node = Node("Smaller").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:519
		{
			yyVAL.node = Node("SmallerOrEqual").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:520
		{
			yyVAL.node = Node("Greater").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:521
		{
			yyVAL.node = Node("GreaterOrEqual").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:522
		{
			yyVAL.node = yyDollar[2].node
		}
	case 212:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:523
		{
			yyVAL.node = Node("Ternary").append(yyDollar[1].node).append(yyDollar[3].node).append(yyDollar[5].node)
		}
	case 213:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:524
		{
			yyVAL.node = Node("Ternary").append(yyDollar[1].node).append(yyDollar[4].node)
		}
	case 214:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:525
		{
			yyVAL.node = Node("Coalesce").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 215:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:526
		{
			yyVAL.node = Node("Empty").append(yyDollar[3].node)
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:527
		{
			yyVAL.node = Node("Include").append(yyDollar[2].node)
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:528
		{
			yyVAL.node = Node("IncludeOnce").append(yyDollar[2].node)
		}
	case 218:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:529
		{
			yyVAL.node = Node("Eval").append(yyDollar[3].node)
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:530
		{
			yyVAL.node = Node("Require").append(yyDollar[2].node)
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:531
		{
			yyVAL.node = Node("RequireOnce").append(yyDollar[2].node)
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:532
		{
			yyVAL.node = Node("CastInt").append(yyDollar[2].node)
		}
	case 222:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:533
		{
			yyVAL.node = Node("CastDouble").append(yyDollar[2].node)
		}
	case 223:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:534
		{
			yyVAL.node = Node("CastString").append(yyDollar[2].node)
		}
	case 224:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:535
		{
			yyVAL.node = Node("CastArray").append(yyDollar[2].node)
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:536
		{
			yyVAL.node = Node("CastObject").append(yyDollar[2].node)
		}
	case 226:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:537
		{
			yyVAL.node = Node("CastBool").append(yyDollar[2].node)
		}
	case 227:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:538
		{
			yyVAL.node = Node("CastUnset").append(yyDollar[2].node)
		}
	case 228:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:539
		{
			yyVAL.node = Node("Silence").append(yyDollar[2].node)
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:541
		{
			yyVAL.node = Node("Print").append(yyDollar[2].node)
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:542
		{
			yyVAL.node = Node("Yield")
		}
	case 231:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:543
		{
			yyVAL.node = Node("Yield").append(yyDollar[2].node)
		}
	case 232:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:544
		{
			yyVAL.node = Node("Yield").append(yyDollar[2].node).append(yyDollar[4].node)
		}
	case 233:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:545
		{
			yyVAL.node = Node("YieldFrom").append(yyDollar[2].node)
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:549
		{
			yyVAL.node = yyDollar[1].node
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:550
		{
			yyVAL.node = yyDollar[1].node
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:554
		{
			yyVAL.node = Node("null")
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:555
		{
			yyVAL.node = yyDollar[1].node
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:559
		{
			yyVAL.node = yyDollar[1].node
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:563
		{
			yyVAL.node = yyDollar[1].node
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:567
		{
			yyVAL.node = Node("Variable").attribute("name", yyDollar[1].token)
		}
	case 241:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:568
		{
			yyVAL.node = yyDollar[3].node
		}
	case 242:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:569
		{
			yyVAL.node = Node("Variable").append(yyDollar[2].node)
		}
	}
	goto yystack /* stack new state and value */
}
