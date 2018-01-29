package scanner

import (
	"bufio"
	"bytes"
	"go/token"
	"io"
	"unicode"

	"github.com/cznic/golex/lex"
	"github.com/z7zmey/php-parser/comment"
	t "github.com/z7zmey/php-parser/token"
)

// Allocate Character classes anywhere in [0x80, 0xFF].
const (
	classUnicodeLeter = iota + 0x80
	classUnicodeDigit
	classUnicodeGraphic
	classOther
)

const T_INCLUDE = 57346
const T_INCLUDE_ONCE = 57347
const T_EXIT = 57348
const T_IF = 57349
const T_LNUMBER = 57350
const T_DNUMBER = 57351
const T_STRING = 57352
const T_STRING_VARNAME = 57353
const T_VARIABLE = 57354
const T_NUM_STRING = 57355
const T_INLINE_HTML = 57356
const T_CHARACTER = 57357
const T_BAD_CHARACTER = 57358
const T_ENCAPSED_AND_WHITESPACE = 57359
const T_CONSTANT_ENCAPSED_STRING = 57360
const T_ECHO = 57361
const T_DO = 57362
const T_WHILE = 57363
const T_ENDWHILE = 57364
const T_FOR = 57365
const T_ENDFOR = 57366
const T_FOREACH = 57367
const T_ENDFOREACH = 57368
const T_DECLARE = 57369
const T_ENDDECLARE = 57370
const T_AS = 57371
const T_SWITCH = 57372
const T_ENDSWITCH = 57373
const T_CASE = 57374
const T_DEFAULT = 57375
const T_BREAK = 57376
const T_CONTINUE = 57377
const T_GOTO = 57378
const T_FUNCTION = 57379
const T_CONST = 57380
const T_RETURN = 57381
const T_TRY = 57382
const T_CATCH = 57383
const T_FINALLY = 57384
const T_THROW = 57385
const T_USE = 57386
const T_INSTEADOF = 57387
const T_GLOBAL = 57388
const T_VAR = 57389
const T_UNSET = 57390
const T_ISSET = 57391
const T_EMPTY = 57392
const T_HALT_COMPILER = 57393
const T_CLASS = 57394
const T_TRAIT = 57395
const T_INTERFACE = 57396
const T_EXTENDS = 57397
const T_IMPLEMENTS = 57398
const T_OBJECT_OPERATOR = 57399
const T_DOUBLE_ARROW = 57400
const T_LIST = 57401
const T_ARRAY = 57402
const T_CALLABLE = 57403
const T_CLASS_C = 57404
const T_TRAIT_C = 57405
const T_METHOD_C = 57406
const T_FUNC_C = 57407
const T_LINE = 57408
const T_FILE = 57409
const T_COMMENT = 57410
const T_DOC_COMMENT = 57411
const T_OPEN_TAG = 57412
const T_OPEN_TAG_WITH_ECHO = 57413
const T_CLOSE_TAG = 57414
const T_WHITESPACE = 57415
const T_START_HEREDOC = 57416
const T_END_HEREDOC = 57417
const T_DOLLAR_OPEN_CURLY_BRACES = 57418
const T_CURLY_OPEN = 57419
const T_PAAMAYIM_NEKUDOTAYIM = 57420
const T_NAMESPACE = 57421
const T_NS_C = 57422
const T_DIR = 57423
const T_NS_SEPARATOR = 57424
const T_ELLIPSIS = 57425
const T_EVAL = 57426
const T_REQUIRE = 57427
const T_REQUIRE_ONCE = 57428
const T_LOGICAL_OR = 57429
const T_LOGICAL_XOR = 57430
const T_LOGICAL_AND = 57431
const T_INSTANCEOF = 57432
const T_NEW = 57433
const T_CLONE = 57434
const T_ELSEIF = 57435
const T_ELSE = 57436
const T_ENDIF = 57437
const T_PRINT = 57438
const T_YIELD = 57439
const T_STATIC = 57440
const T_ABSTRACT = 57441
const T_FINAL = 57442
const T_PRIVATE = 57443
const T_PROTECTED = 57444
const T_PUBLIC = 57445
const T_INC = 57446
const T_DEC = 57447
const T_YIELD_FROM = 57448
const T_INT_CAST = 57449
const T_DOUBLE_CAST = 57450
const T_STRING_CAST = 57451
const T_ARRAY_CAST = 57452
const T_OBJECT_CAST = 57453
const T_BOOL_CAST = 57454
const T_UNSET_CAST = 57455
const T_COALESCE = 57456
const T_SPACESHIP = 57457
const T_NOELSE = 57458
const T_PLUS_EQUAL = 57459
const T_MINUS_EQUAL = 57460
const T_MUL_EQUAL = 57461
const T_DIV_EQUAL = 57462
const T_CONCAT_EQUAL = 57463
const T_MOD_EQUAL = 57464
const T_AND_EQUAL = 57465
const T_OR_EQUAL = 57466
const T_XOR_EQUAL = 57467
const T_SL_EQUAL = 57468
const T_SR_EQUAL = 57469
const T_POW_EQUAL = 57470
const T_BOOLEAN_OR = 57471
const T_BOOLEAN_AND = 57472
const T_IS_EQUAL = 57473
const T_IS_NOT_EQUAL = 57474
const T_IS_IDENTICAL = 57475
const T_IS_NOT_IDENTICAL = 57476
const T_IS_SMALLER_OR_EQUAL = 57477
const T_IS_GREATER_OR_EQUAL = 57478
const T_SL = 57479
const T_SR = 57480
const T_POW = 57481

type Lval interface {
	Token(tkn t.Token)
}

type Lexer struct {
	*lex.Lexer
	StateStack    []int
	PhpDocComment string
	Comments      []comment.Comment
}

func Rune2Class(r rune) int {
	if r >= 0 && r < 0x80 { // Keep ASCII as it is.
		return int(r)
	}
	if unicode.IsLetter(r) {
		return classUnicodeLeter
	}
	if unicode.IsDigit(r) {
		return classUnicodeDigit
	}
	if unicode.IsGraphic(r) {
		return classUnicodeGraphic
	}
	// return classOther
	return -1
}

func NewLexer(src io.Reader, fName string) *Lexer {
	file := token.NewFileSet().AddFile(fName, -1, 1<<31-1)
	lx, err := lex.New(file, bufio.NewReader(src), lex.RuneClass(Rune2Class))
	if err != nil {
		panic(err)
	}
	return &Lexer{lx, []int{0}, "", nil}
}

func (l *Lexer) ungetChars(n int) []lex.Char {
	l.Unget(l.Lookahead())

	chars := l.Token()

	for i := 1; i <= n; i++ {
		char := chars[len(chars)-i]
		l.Unget(char)
	}

	buf := l.Token()
	buf = buf[:len(buf)-n]

	return buf
}

func (l *Lexer) pushState(state int) {
	l.StateStack = append(l.StateStack, state)
}

func (l *Lexer) popState() {
	len := len(l.StateStack)
	if len <= 1 {
		return
	}

	l.StateStack = l.StateStack[:len-1]
}

func (l *Lexer) begin(state int) {
	len := len(l.StateStack)
	l.StateStack = l.StateStack[:len-1]
	l.StateStack = append(l.StateStack, state)
}

func (l *Lexer) getCurrentState() int {
	return l.StateStack[len(l.StateStack)-1]
}

func (l *Lexer) newToken(chars []lex.Char) t.Token {
	firstChar := chars[0]
	lastChar := chars[len(chars)-1]

	startLine := l.File.Line(firstChar.Pos())
	endLine := l.File.Line(lastChar.Pos())
	startPos := int(firstChar.Pos())
	endPos := int(lastChar.Pos())

	return t.NewToken(l.charsToBytes(chars), startLine, endLine, startPos, endPos).SetComments(l.Comments)
}

func (l *Lexer) addComment(c comment.Comment) {
	l.Comments = append(l.Comments, c)
}

func (l *Lexer) charsToBytes(chars []lex.Char) []byte {
	bytesBuf := bytes.Buffer{}

	for _, c := range chars {
		bytesBuf.WriteRune(c.Rune)
	}

	return bytesBuf.Bytes()
}
