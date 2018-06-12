// Package scanner transforms an input string into a stream of PHP tokens.
package scanner

import (
	"bufio"
	"bytes"
	"go/token"
	"io"
	"sync"
	"unicode"

	"github.com/z7zmey/php-parser/position"

	"github.com/cznic/golex/lex"
	"github.com/z7zmey/php-parser/comment"
)

// Allocate Character classes anywhere in [0x80, 0xFF].
const (
	classUnicodeLeter = iota + 0x80
	classUnicodeDigit
	classUnicodeGraphic
	classOther
)

// T_INCLUDE token
const T_INCLUDE = 57346

// T_INCLUDE_ONCE token
const T_INCLUDE_ONCE = 57347

// T_EXIT token
const T_EXIT = 57348

// T_IF token
const T_IF = 57349

// T_LNUMBER token
const T_LNUMBER = 57350

// T_DNUMBER token
const T_DNUMBER = 57351

// T_STRING token
const T_STRING = 57352

// T_STRING_VARNAME token
const T_STRING_VARNAME = 57353

// T_VARIABLE token
const T_VARIABLE = 57354

// T_NUM_STRING token
const T_NUM_STRING = 57355

// T_INLINE_HTML token
const T_INLINE_HTML = 57356

// T_CHARACTER token
const T_CHARACTER = 57357

// T_BAD_CHARACTER token
const T_BAD_CHARACTER = 57358

// T_ENCAPSED_AND_WHITESPACE token
const T_ENCAPSED_AND_WHITESPACE = 57359

// T_CONSTANT_ENCAPSED_STRING token
const T_CONSTANT_ENCAPSED_STRING = 57360

// T_ECHO token
const T_ECHO = 57361

// T_DO token
const T_DO = 57362

// T_WHILE token
const T_WHILE = 57363

// T_ENDWHILE token
const T_ENDWHILE = 57364

// T_FOR token
const T_FOR = 57365

// T_ENDFOR token
const T_ENDFOR = 57366

// T_FOREACH token
const T_FOREACH = 57367

// T_ENDFOREACH token
const T_ENDFOREACH = 57368

// T_DECLARE token
const T_DECLARE = 57369

// T_ENDDECLARE token
const T_ENDDECLARE = 57370

// T_AS token
const T_AS = 57371

// T_SWITCH token
const T_SWITCH = 57372

// T_ENDSWITCH token
const T_ENDSWITCH = 57373

// T_CASE token
const T_CASE = 57374

// T_DEFAULT token
const T_DEFAULT = 57375

// T_BREAK token
const T_BREAK = 57376

// T_CONTINUE token
const T_CONTINUE = 57377

// T_GOTO token
const T_GOTO = 57378

// T_FUNCTION token
const T_FUNCTION = 57379

// T_CONST token
const T_CONST = 57380

// T_RETURN token
const T_RETURN = 57381

// T_TRY token
const T_TRY = 57382

// T_CATCH token
const T_CATCH = 57383

// T_FINALLY token
const T_FINALLY = 57384

// T_THROW token
const T_THROW = 57385

// T_USE token
const T_USE = 57386

// T_INSTEADOF token
const T_INSTEADOF = 57387

// T_GLOBAL token
const T_GLOBAL = 57388

// T_VAR token
const T_VAR = 57389

// T_UNSET token
const T_UNSET = 57390

// T_ISSET token
const T_ISSET = 57391

// T_EMPTY token
const T_EMPTY = 57392

// T_HALT_COMPILER token
const T_HALT_COMPILER = 57393

// T_CLASS token
const T_CLASS = 57394

// T_TRAIT token
const T_TRAIT = 57395

// T_INTERFACE token
const T_INTERFACE = 57396

// T_EXTENDS token
const T_EXTENDS = 57397

// T_IMPLEMENTS token
const T_IMPLEMENTS = 57398

// T_OBJECT_OPERATOR token
const T_OBJECT_OPERATOR = 57399

// T_DOUBLE_ARROW token
const T_DOUBLE_ARROW = 57400

// T_LIST token
const T_LIST = 57401

// T_ARRAY token
const T_ARRAY = 57402

// T_CALLABLE token
const T_CALLABLE = 57403

// T_CLASS_C token
const T_CLASS_C = 57404

// T_TRAIT_C token
const T_TRAIT_C = 57405

// T_METHOD_C token
const T_METHOD_C = 57406

// T_FUNC_C token
const T_FUNC_C = 57407

// T_LINE token
const T_LINE = 57408

// T_FILE token
const T_FILE = 57409

// T_COMMENT token
const T_COMMENT = 57410

// T_DOC_COMMENT token
const T_DOC_COMMENT = 57411

// T_OPEN_TAG token
const T_OPEN_TAG = 57412

// T_OPEN_TAG_WITH_ECHO token
const T_OPEN_TAG_WITH_ECHO = 57413

// T_CLOSE_TAG token
const T_CLOSE_TAG = 57414

// T_WHITESPACE token
const T_WHITESPACE = 57415

// T_START_HEREDOC token
const T_START_HEREDOC = 57416

// T_END_HEREDOC token
const T_END_HEREDOC = 57417

// T_DOLLAR_OPEN_CURLY_BRACES token
const T_DOLLAR_OPEN_CURLY_BRACES = 57418

// T_CURLY_OPEN token
const T_CURLY_OPEN = 57419

// T_PAAMAYIM_NEKUDOTAYIM token
const T_PAAMAYIM_NEKUDOTAYIM = 57420

// T_NAMESPACE token
const T_NAMESPACE = 57421

// T_NS_C token
const T_NS_C = 57422

// T_DIR token
const T_DIR = 57423

// T_NS_SEPARATOR token
const T_NS_SEPARATOR = 57424

// T_ELLIPSIS token
const T_ELLIPSIS = 57425

// T_EVAL token
const T_EVAL = 57426

// T_REQUIRE token
const T_REQUIRE = 57427

// T_REQUIRE_ONCE token
const T_REQUIRE_ONCE = 57428

// T_LOGICAL_OR token
const T_LOGICAL_OR = 57429

// T_LOGICAL_XOR token
const T_LOGICAL_XOR = 57430

// T_LOGICAL_AND token
const T_LOGICAL_AND = 57431

// T_INSTANCEOF token
const T_INSTANCEOF = 57432

// T_NEW token
const T_NEW = 57433

// T_CLONE token
const T_CLONE = 57434

// T_ELSEIF token
const T_ELSEIF = 57435

// T_ELSE token
const T_ELSE = 57436

// T_ENDIF token
const T_ENDIF = 57437

// T_PRINT token
const T_PRINT = 57438

// T_YIELD token
const T_YIELD = 57439

// T_STATIC token
const T_STATIC = 57440

// T_ABSTRACT token
const T_ABSTRACT = 57441

// T_FINAL token
const T_FINAL = 57442

// T_PRIVATE token
const T_PRIVATE = 57443

// T_PROTECTED token
const T_PROTECTED = 57444

// T_PUBLIC token
const T_PUBLIC = 57445

// T_INC token
const T_INC = 57446

// T_DEC token
const T_DEC = 57447

// T_YIELD_FROM token
const T_YIELD_FROM = 57448

// T_INT_CAST token
const T_INT_CAST = 57449

// T_DOUBLE_CAST token
const T_DOUBLE_CAST = 57450

// T_STRING_CAST token
const T_STRING_CAST = 57451

// T_ARRAY_CAST token
const T_ARRAY_CAST = 57452

// T_OBJECT_CAST token
const T_OBJECT_CAST = 57453

// T_BOOL_CAST token
const T_BOOL_CAST = 57454

// T_UNSET_CAST token
const T_UNSET_CAST = 57455

// T_COALESCE token
const T_COALESCE = 57456

// T_SPACESHIP token
const T_SPACESHIP = 57457

// T_NOELSE token
const T_NOELSE = 57458

// T_PLUS_EQUAL token
const T_PLUS_EQUAL = 57459

// T_MINUS_EQUAL token
const T_MINUS_EQUAL = 57460

// T_MUL_EQUAL token
const T_MUL_EQUAL = 57461

// T_POW_EQUAL token
const T_POW_EQUAL = 57462

// T_DIV_EQUAL token
const T_DIV_EQUAL = 57463

// T_CONCAT_EQUAL token
const T_CONCAT_EQUAL = 57464

// T_MOD_EQUAL token
const T_MOD_EQUAL = 57465

// T_AND_EQUAL token
const T_AND_EQUAL = 57466

// T_OR_EQUAL token
const T_OR_EQUAL = 57467

// T_XOR_EQUAL token
const T_XOR_EQUAL = 57468

// T_SL_EQUAL token
const T_SL_EQUAL = 57469

// T_SR_EQUAL token
const T_SR_EQUAL = 57470

// T_BOOLEAN_OR token
const T_BOOLEAN_OR = 57471

// T_BOOLEAN_AND token
const T_BOOLEAN_AND = 57472

// T_POW token
const T_POW = 57473

// T_SL token
const T_SL = 57474

// T_SR token
const T_SR = 57475

// T_IS_IDENTICAL token
const T_IS_IDENTICAL = 57476

// T_IS_NOT_IDENTICAL token
const T_IS_NOT_IDENTICAL = 57477

// T_IS_EQUAL token
const T_IS_EQUAL = 57478

// T_IS_NOT_EQUAL token
const T_IS_NOT_EQUAL = 57479

// T_IS_SMALLER_OR_EQUAL token
const T_IS_SMALLER_OR_EQUAL = 57480

// T_IS_GREATER_OR_EQUAL token
const T_IS_GREATER_OR_EQUAL = 57481

// Lval parsers yySymType must implement this interface
type Lval interface {
	Token(tkn *Token)
}

// Lexer php lexer
type Lexer struct {
	*lex.Lexer
	StateStack    []int
	PhpDocComment string
	Comments      []*comment.Comment
	heredocLabel  string
	tokenBytesBuf *bytes.Buffer
	TokenPool     sync.Pool
	PositionPool  sync.Pool
}

// Rune2Class returns the rune integer id
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

// NewLexer the Lexer constructor
func NewLexer(src io.Reader, fName string) *Lexer {
	file := token.NewFileSet().AddFile(fName, -1, 1<<31-3)
	lx, err := lex.New(file, bufio.NewReader(src), lex.RuneClass(Rune2Class))
	if err != nil {
		panic(err)
	}

	return &Lexer{
		Lexer:         lx,
		StateStack:    []int{0},
		PhpDocComment: "",
		Comments:      nil,
		heredocLabel:  "",
		tokenBytesBuf: &bytes.Buffer{},
		TokenPool: sync.Pool{
			New: func() interface{} { return &Token{} },
		},
		PositionPool: sync.Pool{
			New: func() interface{} { return &position.Position{} },
		},
	}
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

func (l *Lexer) createToken(chars []lex.Char) *Token {
	firstChar := chars[0]
	lastChar := chars[len(chars)-1]

	pos := l.PositionPool.Get().(*position.Position)

	pos.StartLine = l.File.Line(firstChar.Pos())
	pos.EndLine = l.File.Line(lastChar.Pos())
	pos.StartPos = int(firstChar.Pos())
	pos.EndPos = int(lastChar.Pos())

	token := l.TokenPool.Get().(*Token)
	token.Position = pos
	token.Comments = l.Comments
	token.Value = l.tokenString(chars)

	return token
}

func (l *Lexer) addComment(chars []lex.Char) {
	firstChar := chars[0]
	lastChar := chars[len(chars)-1]

	pos := position.NewPosition(
		l.File.Line(firstChar.Pos()),
		l.File.Line(lastChar.Pos()),
		int(firstChar.Pos()),
		int(lastChar.Pos()),
	)

	c := comment.NewComment(l.tokenString(chars), pos)
	l.Comments = append(l.Comments, c)
}

func (l *Lexer) tokenString(chars []lex.Char) string {
	l.tokenBytesBuf.Reset()

	for _, c := range chars {
		l.tokenBytesBuf.WriteRune(c.Rune)
	}

	return string(l.tokenBytesBuf.Bytes())
}
