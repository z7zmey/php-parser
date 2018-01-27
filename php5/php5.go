//line php5/php5.y:2
package php5

import __yyfmt__ "fmt"

//line php5/php5.y:2
import (
	//    "strings"
	//    "strconv"
	"github.com/z7zmey/php-parser/token" //    "github.com/z7zmey/php-parser/node"
	//    "github.com/z7zmey/php-parser/node/scalar"
	//    "github.com/z7zmey/php-parser/node/name"
	//    "github.com/z7zmey/php-parser/node/stmt"
	//    "github.com/z7zmey/php-parser/node/expr"
//    "github.com/z7zmey/php-parser/node/expr/assign_op"
//    "github.com/z7zmey/php-parser/node/expr/binary_op"
//    "github.com/z7zmey/php-parser/node/expr/cast"
)

//line php5/php5.y:21
type yySymType struct {
	yys   int
	token token.Token
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
const T_PLUS_EQUAL = 57356
const T_MINUS_EQUAL = 57357
const T_MUL_EQUAL = 57358
const T_DIV_EQUAL = 57359
const T_CONCAT_EQUAL = 57360
const T_MOD_EQUAL = 57361
const T_AND_EQUAL = 57362
const T_OR_EQUAL = 57363
const T_XOR_EQUAL = 57364
const T_SL_EQUAL = 57365
const T_SR_EQUAL = 57366
const T_POW_EQUAL = 57367
const T_BOOLEAN_OR = 57368
const T_BOOLEAN_AND = 57369
const T_IS_EQUAL = 57370
const T_IS_NOT_EQUAL = 57371
const T_IS_IDENTICAL = 57372
const T_IS_NOT_IDENTICAL = 57373
const T_IS_SMALLER_OR_EQUAL = 57374
const T_IS_GREATER_OR_EQUAL = 57375
const T_SL = 57376
const T_SR = 57377
const T_INSTANCEOF = 57378
const T_INC = 57379
const T_DEC = 57380
const T_INT_CAST = 57381
const T_DOUBLE_CAST = 57382
const T_STRING_CAST = 57383
const T_ARRAY_CAST = 57384
const T_OBJECT_CAST = 57385
const T_BOOL_CAST = 57386
const T_UNSET_CAST = 57387
const T_POW = 57388
const T_NEW = 57389
const T_CLONE = 57390
const T_ELSEIF = 57391
const T_ELSE = 57392
const T_ENDIF = 57393
const T_STATIC = 57394
const T_ABSTRACT = 57395
const T_FINAL = 57396
const T_PRIVATE = 57397
const T_PROTECTED = 57398
const T_PUBLIC = 57399
const T_EXIT = 57400
const T_IF = 57401
const T_LNUMBER = 57402
const T_DNUMBER = 57403
const T_STRING = 57404
const T_STRING_VARNAME = 57405
const T_VARIABLE = 57406
const T_NUM_STRING = 57407
const T_INLINE_HTML = 57408
const T_CHARACTER = 57409
const T_BAD_CHARACTER = 57410
const T_ENCAPSED_AND_WHITESPACE = 57411
const T_CONSTANT_ENCAPSED_STRING = 57412
const T_ECHO = 57413
const T_DO = 57414
const T_WHILE = 57415
const T_ENDWHILE = 57416
const T_FOR = 57417
const T_ENDFOR = 57418
const T_FOREACH = 57419
const T_ENDFOREACH = 57420
const T_DECLARE = 57421
const T_ENDDECLARE = 57422
const T_AS = 57423
const T_SWITCH = 57424
const T_ENDSWITCH = 57425
const T_CASE = 57426
const T_DEFAULT = 57427
const T_BREAK = 57428
const T_CONTINUE = 57429
const T_GOTO = 57430
const T_FUNCTION = 57431
const T_CONST = 57432
const T_RETURN = 57433
const T_TRY = 57434
const T_CATCH = 57435
const T_FINALLY = 57436
const T_THROW = 57437
const T_USE = 57438
const T_INSTEADOF = 57439
const T_GLOBAL = 57440
const T_VAR = 57441
const T_UNSET = 57442
const T_ISSET = 57443
const T_EMPTY = 57444
const T_HALT_COMPILER = 57445
const T_CLASS = 57446
const T_TRAIT = 57447
const T_INTERFACE = 57448
const T_EXTENDS = 57449
const T_IMPLEMENTS = 57450
const T_OBJECT_OPERATOR = 57451
const T_DOUBLE_ARROW = 57452
const T_LIST = 57453
const T_ARRAY = 57454
const T_CALLABLE = 57455
const T_CLASS_C = 57456
const T_TRAIT_C = 57457
const T_METHOD_C = 57458
const T_FUNC_C = 57459
const T_LINE = 57460
const T_FILE = 57461
const T_COMMENT = 57462
const T_DOC_COMMENT = 57463
const T_OPEN_TAG = 57464
const T_OPEN_TAG_WITH_ECHO = 57465
const T_CLOSE_TAG = 57466
const T_WHITESPACE = 57467
const T_START_HEREDOC = 57468
const T_END_HEREDOC = 57469
const T_DOLLAR_OPEN_CURLY_BRACES = 57470
const T_CURLY_OPEN = 57471
const T_PAAMAYIM_NEKUDOTAYIM = 57472
const T_NAMESPACE = 57473
const T_NS_C = 57474
const T_DIR = 57475
const T_NS_SEPARATOR = 57476
const T_ELLIPSIS = 57477
const T_YIELD_FROM = 57478

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
	"T_BOOLEAN_OR",
	"T_BOOLEAN_AND",
	"'|'",
	"'^'",
	"'&'",
	"T_IS_EQUAL",
	"T_IS_NOT_EQUAL",
	"T_IS_IDENTICAL",
	"T_IS_NOT_IDENTICAL",
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
	"T_DOUBLE_ARROW",
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
	"T_YIELD_FROM",
	"'\"'",
	"'`'",
	"'{'",
	"'}'",
	"';'",
	"'('",
	"')'",
	"']'",
	"'$'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line php5/php5.y:1262

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 2

var yyAct = [...]int{

	2, 1,
}
var yyPact = [...]int{

	-1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 3, 3, 4, 4, 4, 4,
	4, 8, 4, 9, 4, 4, 4, 4, 4, 10,
	10, 14, 14, 14, 14, 11, 11, 15, 15, 15,
	15, 12, 12, 16, 16, 16, 16, 13, 13, 18,
	18, 19, 19, 19, 19, 5, 5, 20, 22, 23,
	20, 26, 27, 20, 30, 31, 20, 33, 34, 20,
	36, 37, 38, 20, 40, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	50, 53, 20, 55, 56, 20, 57, 20, 20, 60,
	62, 20, 20, 20, 61, 64, 66, 67, 68, 61,
	63, 70, 63, 69, 69, 71, 71, 73, 74, 72,
	49, 49, 75, 6, 7, 78, 78, 79, 79, 81,
	76, 85, 77, 89, 77, 83, 83, 83, 83, 84,
	84, 88, 90, 90, 86, 86, 91, 91, 52, 52,
	51, 51, 92, 51, 39, 39, 54, 54, 59, 59,
	58, 58, 41, 41, 41, 41, 94, 96, 94, 97,
	94, 95, 95, 32, 32, 24, 98, 24, 28, 99,
	28, 25, 25, 29, 29, 82, 82, 100, 100, 101,
	101, 102, 102, 102, 102, 103, 103, 103, 104, 104,
	105, 105, 105, 105, 46, 46, 107, 107, 107, 47,
	47, 47, 47, 87, 87, 111, 109, 109, 109, 116,
	109, 114, 118, 118, 119, 119, 120, 120, 121, 121,
	122, 122, 123, 126, 126, 127, 127, 125, 124, 124,
	128, 128, 117, 117, 110, 110, 115, 115, 130, 130,
	129, 129, 129, 129, 129, 129, 112, 112, 112, 112,
	113, 113, 48, 48, 35, 35, 132, 131, 131, 133,
	133, 135, 135, 138, 137, 137, 137, 139, 140, 139,
	143, 141, 145, 43, 43, 43, 146, 43, 43, 43,
	43, 43, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 43, 43, 43, 43, 148, 43, 149, 43, 150,
	43, 151, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 43, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 43, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 152, 43, 153, 154, 43, 155, 43, 43, 43,
	43, 43, 43, 43, 43, 43, 43, 158, 43, 43,
	43, 43, 43, 43, 43, 163, 43, 165, 43, 45,
	45, 45, 45, 160, 160, 160, 160, 161, 161, 80,
	164, 164, 168, 168, 168, 168, 170, 169, 171, 169,
	172, 169, 175, 169, 177, 169, 179, 169, 180, 169,
	181, 169, 173, 173, 173, 173, 65, 65, 65, 142,
	142, 184, 186, 182, 182, 187, 187, 188, 157, 157,
	157, 162, 162, 162, 144, 144, 190, 190, 190, 190,
	190, 190, 190, 190, 190, 190, 190, 190, 191, 17,
	192, 192, 192, 192, 192, 192, 192, 192, 192, 192,
	195, 195, 195, 195, 195, 195, 195, 195, 195, 195,
	195, 195, 195, 195, 195, 195, 195, 195, 195, 195,
	195, 195, 195, 195, 195, 195, 195, 195, 195, 195,
	195, 195, 195, 166, 166, 166, 166, 159, 159, 159,
	159, 159, 159, 159, 194, 194, 199, 199, 198, 198,
	198, 198, 42, 42, 21, 21, 108, 106, 147, 201,
	202, 44, 44, 204, 204, 205, 134, 206, 206, 208,
	207, 203, 203, 203, 176, 176, 211, 211, 178, 212,
	213, 212, 200, 200, 200, 183, 183, 183, 209, 209,
	209, 214, 214, 136, 136, 185, 185, 215, 215, 215,
	174, 174, 210, 210, 93, 93, 216, 217, 216, 216,
	167, 167, 218, 218, 218, 218, 218, 218, 218, 218,
	189, 189, 189, 189, 219, 220, 219, 219, 219, 219,
	219, 221, 221, 221, 156, 156, 156, 156, 156, 156,
	156, 156, 222, 224, 222, 223, 223, 196, 196, 193,
	197,
}
var yyR2 = [...]int{

	0, 1, 1, 0, 1, 3, 1, 1, 1, 4,
	3, 0, 6, 0, 5, 3, 4, 4, 2, 3,
	1, 1, 3, 2, 4, 3, 1, 1, 3, 2,
	4, 3, 1, 1, 3, 2, 4, 5, 4, 1,
	0, 1, 1, 1, 4, 1, 2, 3, 0, 0,
	7, 0, 0, 10, 0, 0, 5, 0, 0, 7,
	0, 0, 0, 12, 0, 4, 2, 3, 2, 3,
	2, 3, 3, 2, 3, 3, 3, 1, 2, 5,
	0, 0, 10, 0, 0, 10, 0, 6, 1, 0,
	0, 8, 3, 3, 0, 0, 0, 0, 0, 13,
	0, 0, 5, 1, 0, 1, 2, 0, 0, 10,
	1, 3, 1, 1, 1, 0, 1, 0, 1, 0,
	10, 0, 8, 0, 7, 1, 2, 1, 2, 0,
	2, 1, 0, 2, 0, 2, 1, 3, 0, 2,
	1, 2, 0, 5, 1, 4, 1, 4, 1, 4,
	3, 5, 3, 4, 4, 5, 0, 0, 6, 0,
	5, 1, 1, 1, 4, 0, 0, 5, 0, 0,
	6, 0, 2, 0, 3, 1, 0, 1, 3, 4,
	6, 0, 1, 1, 1, 2, 3, 3, 1, 3,
	1, 1, 2, 2, 3, 1, 1, 2, 4, 3,
	5, 1, 3, 2, 0, 0, 4, 2, 1, 0,
	9, 3, 1, 3, 1, 3, 0, 1, 1, 2,
	2, 2, 3, 1, 3, 1, 1, 3, 4, 3,
	0, 1, 1, 3, 1, 1, 0, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 3, 5, 1, 3,
	5, 4, 3, 1, 0, 1, 0, 4, 1, 2,
	1, 4, 3, 0, 3, 1, 1, 0, 0, 2,
	0, 4, 0, 7, 3, 4, 0, 7, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 0, 4, 0, 4, 0,
	4, 0, 4, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 1,
	1, 0, 5, 0, 0, 7, 0, 5, 1, 2,
	2, 2, 2, 2, 2, 2, 2, 0, 3, 1,
	1, 1, 3, 2, 1, 0, 10, 0, 11, 2,
	2, 4, 4, 4, 4, 4, 4, 4, 3, 1,
	0, 4, 3, 4, 1, 2, 0, 3, 0, 5,
	0, 4, 0, 5, 0, 5, 0, 5, 0, 5,
	0, 3, 1, 1, 3, 2, 1, 3, 2, 1,
	1, 0, 0, 6, 1, 2, 0, 2, 0, 2,
	1, 0, 1, 1, 0, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 2, 3, 1,
	1, 1, 1, 3, 2, 4, 3, 1, 1, 1,
	4, 3, 3, 3, 3, 3, 3, 2, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 5,
	2, 2, 3, 1, 1, 3, 2, 1, 1, 1,
	1, 3, 3, 1, 0, 2, 0, 1, 5, 3,
	3, 1, 1, 1, 3, 3, 1, 1, 1, 0,
	0, 7, 1, 2, 0, 0, 4, 4, 4, 0,
	2, 1, 1, 0, 1, 2, 3, 3, 1, 4,
	0, 5, 1, 1, 1, 1, 2, 1, 4, 4,
	1, 1, 4, 0, 1, 1, 1, 4, 4, 1,
	1, 3, 1, 2, 3, 1, 1, 0, 5, 0,
	0, 2, 5, 3, 3, 1, 6, 4, 4, 2,
	2, 2, 1, 2, 1, 0, 5, 3, 3, 6,
	3, 1, 1, 1, 4, 4, 4, 2, 2, 4,
	2, 2, 1, 0, 4, 1, 1, 3, 3, 3,
	3,
}
var yyChk = [...]int{

	-1000, -1, -2,
}
var yyDef = [...]int{

	3, -2, 1,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 51, 156, 3, 164, 50, 34, 3,
	161, 162, 48, 45, 9, 46, 47, 49, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 29, 160,
	39, 15, 41, 28, 63, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 65, 3, 163, 33, 3, 157, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 158, 32, 159, 53,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 10, 11, 12,
	13, 14, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 30, 31, 35, 36, 37, 38,
	40, 42, 43, 44, 52, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 64, 66, 67, 68, 69, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
	121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
	151, 152, 153, 154, 155,
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
		//line php5/php5.y:188
		{
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:192
		{
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:197
		{
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:198
		{
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:202
		{
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:203
		{
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:204
		{
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:205
		{
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:206
		{
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:207
		{
		}
	case 12:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line php5/php5.y:208
		{
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:209
		{
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:210
		{
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:211
		{
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:212
		{
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:213
		{
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:214
		{
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:223
		{
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:224
		{
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:225
		{
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:226
		{
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:235
		{
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:236
		{
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:237
		{
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:238
		{
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:247
		{
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:248
		{
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:249
		{
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:250
		{
		}
	case 37:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:254
		{
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:255
		{
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:259
		{
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:268
		{
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:273
		{
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:274
		{
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:279
		{
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:279
		{
		}
	case 50:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line php5/php5.y:279
		{
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:280
		{
		}
	case 52:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:280
		{
		}
	case 53:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line php5/php5.y:280
		{
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:281
		{
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:281
		{
		}
	case 56:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:281
		{
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:282
		{
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:282
		{
		}
	case 59:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line php5/php5.y:282
		{
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:286
		{
		}
	case 61:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line php5/php5.y:288
		{
		}
	case 62:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line php5/php5.y:290
		{
		}
	case 63:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line php5/php5.y:291
		{
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:292
		{
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:292
		{
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:293
		{
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:294
		{
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:295
		{
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:296
		{
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:297
		{
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:298
		{
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:299
		{
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:300
		{
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:304
		{
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:305
		{
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:308
		{
		}
	case 81:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line php5/php5.y:309
		{
		}
	case 82:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line php5/php5.y:310
		{
		}
	case 83:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:312
		{
		}
	case 84:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line php5/php5.y:313
		{
		}
	case 85:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line php5/php5.y:314
		{
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:315
		{
		}
	case 87:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line php5/php5.y:315
		{
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:317
		{
		}
	case 90:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line php5/php5.y:318
		{
		}
	case 91:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line php5/php5.y:319
		{
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:320
		{
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:321
		{
		}
	case 94:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:325
		{
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:326
		{
		}
	case 96:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:327
		{
		}
	case 97:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line php5/php5.y:328
		{
		}
	case 98:
		yyDollar = yyS[yypt-11 : yypt+1]
		//line php5/php5.y:329
		{
		}
	case 99:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line php5/php5.y:330
		{
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:333
		{
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:334
		{
		}
	case 102:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:334
		{
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:338
		{
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:339
		{
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:343
		{
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:344
		{
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:348
		{
		}
	case 108:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line php5/php5.y:348
		{
		}
	case 109:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line php5/php5.y:348
		{
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:357
		{
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:361
		{
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:365
		{
		}
	case 115:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:369
		{
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:370
		{
		}
	case 117:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:374
		{
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:375
		{
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:379
		{
		}
	case 120:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line php5/php5.y:381
		{
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:386
		{
		}
	case 122:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line php5/php5.y:390
		{
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:392
		{
		}
	case 124:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line php5/php5.y:396
		{
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:401
		{
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:402
		{
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:403
		{
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:404
		{
		}
	case 129:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:408
		{
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:409
		{
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:413
		{
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:427
		{
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:428
		{
		}
	case 138:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:432
		{
		}
	case 139:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:433
		{
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:437
		{
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:438
		{
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:439
		{
		}
	case 143:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:439
		{
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:461
		{
		}
	case 151:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:462
		{
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:467
		{
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:468
		{
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:469
		{
		}
	case 155:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:470
		{
		}
	case 156:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:475
		{
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:476
		{
		}
	case 158:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line php5/php5.y:476
		{
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:477
		{
		}
	case 160:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:477
		{
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:496
		{
		}
	case 167:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:496
		{
		}
	case 169:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:502
		{
		}
	case 170:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line php5/php5.y:502
		{
		}
	case 179:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:531
		{
		}
	case 180:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line php5/php5.y:533
		{
		}
	case 181:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:538
		{
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:539
		{
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:540
		{
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:541
		{
		}
	case 185:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:546
		{
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:547
		{
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:548
		{
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:558
		{
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:559
		{
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:560
		{
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:561
		{
		}
	case 194:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:565
		{
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:566
		{
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:571
		{
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:572
		{
		}
	case 198:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:573
		{
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:578
		{
		}
	case 200:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:579
		{
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:580
		{
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:581
		{
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:593
		{
		}
	case 209:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:596
		{
		}
	case 210:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line php5/php5.y:598
		{
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:606
		{
		}
	case 213:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:607
		{
		}
	case 222:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:631
		{
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:635
		{
		}
	case 224:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:636
		{
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:640
		{
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:641
		{
		}
	case 227:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:645
		{
		}
	case 228:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:649
		{
		}
	case 229:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:650
		{
		}
	case 230:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:654
		{
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:655
		{
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:659
		{
		}
	case 233:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:660
		{
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:664
		{
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:665
		{
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:669
		{
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:670
		{
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:674
		{
		}
	case 239:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:675
		{
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:679
		{
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:680
		{
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:681
		{
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:682
		{
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:683
		{
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:684
		{
		}
	case 246:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:688
		{
		}
	case 247:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:689
		{
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:690
		{
		}
	case 249:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:691
		{
		}
	case 250:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:695
		{
		}
	case 251:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:696
		{
		}
	case 252:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:700
		{
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:701
		{
		}
	case 254:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:706
		{
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:707
		{
		}
	case 256:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:711
		{
		}
	case 257:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:711
		{
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:712
		{
		}
	case 259:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:716
		{
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:717
		{
		}
	case 261:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:721
		{
		}
	case 262:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:722
		{
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:726
		{
		}
	case 264:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:726
		{
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:727
		{
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:728
		{
		}
	case 267:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:732
		{
		}
	case 268:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:733
		{
		}
	case 269:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:734
		{
		}
	case 270:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:738
		{
		}
	case 271:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:738
		{
		}
	case 272:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:742
		{
		}
	case 273:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line php5/php5.y:742
		{
		}
	case 274:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:743
		{
		}
	case 275:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:744
		{
		}
	case 276:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:745
		{
		}
	case 277:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line php5/php5.y:745
		{
		}
	case 278:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:746
		{
		}
	case 279:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:747
		{
		}
	case 280:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:748
		{
		}
	case 281:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:749
		{
		}
	case 282:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:750
		{
		}
	case 283:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:751
		{
		}
	case 284:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:752
		{
		}
	case 285:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:753
		{
		}
	case 286:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:754
		{
		}
	case 287:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:755
		{
		}
	case 288:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:756
		{
		}
	case 289:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:757
		{
		}
	case 290:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:758
		{
		}
	case 291:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:759
		{
		}
	case 292:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:760
		{
		}
	case 293:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:761
		{
		}
	case 294:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:762
		{
		}
	case 295:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:763
		{
		}
	case 296:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:763
		{
		}
	case 297:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:764
		{
		}
	case 298:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:764
		{
		}
	case 299:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:765
		{
		}
	case 300:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:765
		{
		}
	case 301:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:766
		{
		}
	case 302:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:766
		{
		}
	case 303:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:767
		{
		}
	case 304:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:768
		{
		}
	case 305:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:769
		{
		}
	case 306:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:770
		{
		}
	case 307:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:771
		{
		}
	case 308:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:772
		{
		}
	case 309:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:773
		{
		}
	case 310:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:774
		{
		}
	case 311:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:775
		{
		}
	case 312:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:776
		{
		}
	case 313:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:777
		{
		}
	case 314:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:778
		{
		}
	case 315:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:779
		{
		}
	case 316:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:780
		{
		}
	case 317:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:781
		{
		}
	case 318:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:782
		{
		}
	case 319:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:783
		{
		}
	case 320:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:784
		{
		}
	case 321:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:785
		{
		}
	case 322:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:786
		{
		}
	case 323:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:787
		{
		}
	case 324:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:788
		{
		}
	case 325:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:789
		{
		}
	case 326:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:790
		{
		}
	case 327:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:791
		{
		}
	case 328:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:792
		{
		}
	case 329:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:793
		{
		}
	case 330:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:794
		{
		}
	case 331:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:795
		{
		}
	case 332:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:795
		{
		}
	case 333:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:796
		{
		}
	case 334:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:797
		{
		}
	case 335:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line php5/php5.y:798
		{
		}
	case 336:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:799
		{
		}
	case 337:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:800
		{
		}
	case 338:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:801
		{
		}
	case 339:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:802
		{
		}
	case 340:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:803
		{
		}
	case 341:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:804
		{
		}
	case 342:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:805
		{
		}
	case 343:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:806
		{
		}
	case 344:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:807
		{
		}
	case 345:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:808
		{
		}
	case 346:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:809
		{
		}
	case 347:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:810
		{
		}
	case 348:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:810
		{
		}
	case 349:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:811
		{
		}
	case 350:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:812
		{
		}
	case 351:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:813
		{
		}
	case 352:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:814
		{
		}
	case 353:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:815
		{
		}
	case 354:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:816
		{
		}
	case 355:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:817
		{
		}
	case 356:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line php5/php5.y:819
		{
		}
	case 357:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:820
		{
		}
	case 358:
		yyDollar = yyS[yypt-11 : yypt+1]
		//line php5/php5.y:822
		{
		}
	case 359:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:826
		{
		}
	case 360:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:827
		{
		}
	case 361:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:828
		{
		}
	case 362:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:829
		{
		}
	case 363:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:833
		{
		}
	case 364:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:834
		{
		}
	case 365:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:835
		{
		}
	case 366:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:836
		{
		}
	case 367:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:840
		{
		}
	case 368:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:841
		{
		}
	case 369:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:845
		{
		}
	case 372:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:854
		{
		}
	case 373:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:855
		{
		}
	case 374:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:856
		{
		}
	case 375:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:857
		{
		}
	case 376:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:861
		{
		}
	case 377:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:862
		{
		}
	case 378:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:863
		{
		}
	case 379:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:864
		{
		}
	case 380:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:865
		{
		}
	case 381:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:866
		{
		}
	case 382:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:867
		{
		}
	case 383:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:868
		{
		}
	case 384:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:869
		{
		}
	case 385:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:870
		{
		}
	case 386:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:871
		{
		}
	case 387:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:872
		{
		}
	case 388:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:873
		{
		}
	case 389:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:874
		{
		}
	case 390:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:875
		{
		}
	case 391:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:876
		{
		}
	case 392:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:880
		{
		}
	case 393:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:881
		{
		}
	case 394:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:882
		{
		}
	case 395:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:883
		{
		}
	case 396:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:887
		{
		}
	case 397:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:888
		{
		}
	case 398:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:889
		{
		}
	case 399:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:895
		{
		}
	case 400:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:896
		{
		}
	case 401:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:901
		{
		}
	case 402:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:902
		{
		}
	case 403:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line php5/php5.y:903
		{
		}
	case 404:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:904
		{
		}
	case 407:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:915
		{
		}
	case 408:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:919
		{
		}
	case 409:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:920
		{
		}
	case 410:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:921
		{
		}
	case 411:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:925
		{
		}
	case 412:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:926
		{
		}
	case 413:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:927
		{
		}
	case 414:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:932
		{
		}
	case 415:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:933
		{
		}
	case 416:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:938
		{
		}
	case 417:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:939
		{
		}
	case 418:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:940
		{
		}
	case 419:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:941
		{
		}
	case 420:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:942
		{
		}
	case 421:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:943
		{
		}
	case 422:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:944
		{
		}
	case 423:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:945
		{
		}
	case 424:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:946
		{
		}
	case 425:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:947
		{
		}
	case 426:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:948
		{
		}
	case 427:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:949
		{
		}
	case 428:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:953
		{
		}
	case 429:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:957
		{
		}
	case 430:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:961
		{
		}
	case 431:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:962
		{
		}
	case 432:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:963
		{
		}
	case 433:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:964
		{
		}
	case 434:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:965
		{
		}
	case 435:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:966
		{
		}
	case 436:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:967
		{
		}
	case 437:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:968
		{
		}
	case 438:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:969
		{
		}
	case 439:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:970
		{
		}
	case 440:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:974
		{
		}
	case 441:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:975
		{
		}
	case 442:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:976
		{
		}
	case 443:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:977
		{
		}
	case 444:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:978
		{
		}
	case 445:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:979
		{
		}
	case 446:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:980
		{
		}
	case 447:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:981
		{
		}
	case 448:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:982
		{
		}
	case 449:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:983
		{
		}
	case 450:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:984
		{
		}
	case 451:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:985
		{
		}
	case 452:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:986
		{
		}
	case 453:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:987
		{
		}
	case 454:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:988
		{
		}
	case 455:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:989
		{
		}
	case 456:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:990
		{
		}
	case 457:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:991
		{
		}
	case 458:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:992
		{
		}
	case 459:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:993
		{
		}
	case 460:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:994
		{
		}
	case 461:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:995
		{
		}
	case 462:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:996
		{
		}
	case 463:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:997
		{
		}
	case 464:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:998
		{
		}
	case 465:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:999
		{
		}
	case 466:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1000
		{
		}
	case 467:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1001
		{
		}
	case 468:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1002
		{
		}
	case 469:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:1003
		{
		}
	case 470:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1004
		{
		}
	case 471:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1005
		{
		}
	case 472:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1006
		{
		}
	case 473:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1010
		{
		}
	case 474:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1011
		{
		}
	case 475:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1012
		{
		}
	case 476:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1013
		{
		}
	case 477:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1017
		{
		}
	case 478:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1018
		{
		}
	case 479:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1019
		{
		}
	case 480:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1020
		{
		}
	case 481:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1021
		{
		}
	case 482:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1022
		{
		}
	case 483:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1023
		{
		}
	case 484:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:1027
		{
		}
	case 485:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1028
		{
		}
	case 488:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:1037
		{
		}
	case 489:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1038
		{
		}
	case 490:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1039
		{
		}
	case 491:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1040
		{
		}
	case 492:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1044
		{
		}
	case 493:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1045
		{
		}
	case 494:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1049
		{
		}
	case 495:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1050
		{
		}
	case 496:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1055
		{
		}
	case 497:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1060
		{
		}
	case 498:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1064
		{
		}
	case 499:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1068
		{
		}
	case 500:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1069
		{
		}
	case 501:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line php5/php5.y:1070
		{
		}
	case 502:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1071
		{
		}
	case 503:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1075
		{
		}
	case 504:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:1076
		{
		}
	case 505:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1081
		{
		}
	case 506:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1081
		{
		}
	case 507:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1085
		{
		}
	case 508:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1086
		{
		}
	case 509:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:1090
		{
		}
	case 510:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1091
		{
		}
	case 511:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1095
		{
		}
	case 512:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1096
		{
		}
	case 513:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:1097
		{
		}
	case 514:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1101
		{
		}
	case 515:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1102
		{
		}
	case 516:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1106
		{
		}
	case 517:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1107
		{
		}
	case 518:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1112
		{
		}
	case 519:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1116
		{
		}
	case 520:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1117
		{
		}
	case 521:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:1118
		{
		}
	case 522:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1122
		{
		}
	case 523:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1123
		{
		}
	case 524:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1124
		{
		}
	case 525:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1129
		{
		}
	case 526:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1130
		{
		}
	case 527:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1131
		{
		}
	case 528:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1135
		{
		}
	case 529:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1136
		{
		}
	case 530:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1137
		{
		}
	case 531:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1142
		{
		}
	case 532:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1143
		{
		}
	case 533:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:1147
		{
		}
	case 534:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1148
		{
		}
	case 535:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1153
		{
		}
	case 536:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1154
		{
		}
	case 537:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1158
		{
		}
	case 538:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1159
		{
		}
	case 539:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1160
		{
		}
	case 540:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1164
		{
		}
	case 541:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1165
		{
		}
	case 542:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1169
		{
		}
	case 543:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1170
		{
		}
	case 546:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1180
		{
		}
	case 547:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1181
		{
		}
	case 548:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:1181
		{
		}
	case 549:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:1182
		{
		}
	case 550:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line php5/php5.y:1187
		{
		}
	case 551:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1188
		{
		}
	case 552:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:1192
		{
		}
	case 553:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1193
		{
		}
	case 554:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1194
		{
		}
	case 555:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1195
		{
		}
	case 556:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line php5/php5.y:1196
		{
		}
	case 557:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1197
		{
		}
	case 558:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1198
		{
		}
	case 559:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1199
		{
		}
	case 560:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1203
		{
		}
	case 561:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1204
		{
		}
	case 562:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1205
		{
		}
	case 563:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1206
		{
		}
	case 564:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1212
		{
		}
	case 565:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1213
		{
		}
	case 566:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line php5/php5.y:1213
		{
		}
	case 567:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1214
		{
		}
	case 568:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1215
		{
		}
	case 569:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line php5/php5.y:1216
		{
		}
	case 570:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1217
		{
		}
	case 571:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1222
		{
		}
	case 572:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1223
		{
		}
	case 573:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1224
		{
		}
	case 574:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1229
		{
		}
	case 575:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1230
		{
		}
	case 576:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1231
		{
		}
	case 577:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1232
		{
		}
	case 578:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1233
		{
		}
	case 579:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1234
		{
		}
	case 580:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1235
		{
		}
	case 581:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1236
		{
		}
	case 582:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1240
		{
		}
	case 583:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line php5/php5.y:1241
		{
		}
	case 584:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line php5/php5.y:1241
		{
		}
	case 585:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1245
		{
		}
	case 586:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line php5/php5.y:1246
		{
		}
	case 587:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1250
		{
		}
	case 588:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1251
		{
		}
	case 589:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1255
		{
		}
	case 590:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line php5/php5.y:1259
		{
		}
	}
	goto yystack /* stack new state and value */
}
