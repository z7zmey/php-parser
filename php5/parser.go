package php5

import (
	"io"
	"strings"

	"github.com/z7zmey/php-parser/freefloating"

	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/scanner"
)

func (lval *yySymType) Token(t *scanner.Token) {
	lval.token = t
}

// Parser structure
type Parser struct {
	*scanner.Lexer
	path            string
	currentToken    *scanner.Token
	positionBuilder *parser.PositionBuilder
	rootNode        node.Node
}

// NewParser creates and returns new Parser
func NewParser(src io.Reader, path string) *Parser {
	lexer := scanner.NewLexer(src, path)

	return &Parser{
		lexer,
		path,
		nil,
		nil,
		nil,
	}
}

// Lex proxy to lexer Lex
func (l *Parser) Lex(lval *yySymType) int {
	t := l.Lexer.Lex(lval)
	l.currentToken = lval.token
	return t
}

func (l *Parser) Error(msg string) {
	pos := &position.Position{
		StartLine: l.currentToken.StartLine,
		EndLine:   l.currentToken.EndLine,
		StartPos:  l.currentToken.StartPos,
		EndPos:    l.currentToken.EndPos,
	}

	l.Lexer.Errors = append(l.Lexer.Errors, errors.NewError(msg, pos))
}

func (l *Parser) WithFreeFloating() {
	l.Lexer.WithFreeFloating = true
}

// Parse the php7 Parser entrypoint
func (l *Parser) Parse() int {
	// init
	l.Lexer.Errors = nil
	l.rootNode = nil
	l.positionBuilder = &parser.PositionBuilder{}

	// parse

	return yyParse(l)
}

// GetPath return path to file
func (l *Parser) GetPath() string {
	return l.path
}

// GetRootNode returns root node
func (l *Parser) GetRootNode() node.Node {
	return l.rootNode
}

// GetErrors returns errors list
func (l *Parser) GetErrors() []*errors.Error {
	return l.Lexer.Errors
}

// helpers

func lastNode(nn []node.Node) node.Node {
	if len(nn) == 0 {
		return nil
	}
	return nn[len(nn)-1]
}

func firstNode(nn []node.Node) node.Node {
	return nn[0]
}

func isDollar(r rune) bool {
	return r == '$'
}

func (l *Parser) MoveFreeFloating(src node.Node, dst node.Node) {
	if l.Lexer.WithFreeFloating == false {
		return
	}

	if src.GetFreeFloating() == nil {
		return
	}

	l.setFreeFloating(dst, freefloating.Start, (*src.GetFreeFloating())[freefloating.Start])
	delete((*src.GetFreeFloating()), freefloating.Start)
}

func (l *Parser) setFreeFloating(dst node.Node, p freefloating.Position, strings []freefloating.String) {
	if l.Lexer.WithFreeFloating == false {
		return
	}

	if len(strings) == 0 {
		return
	}

	dstCollection := dst.GetFreeFloating()
	if *dstCollection == nil {
		*dstCollection = make(freefloating.Collection)
	}

	(*dstCollection)[p] = strings
}

func (l *Parser) GetFreeFloatingToken(t *scanner.Token) []freefloating.String {
	if l.Lexer.WithFreeFloating == false {
		return []freefloating.String{}
	}

	return t.GetFreeFloatingToken()
}

func (l *Parser) addDollarToken(v node.Node) {
	if l.Lexer.WithFreeFloating == false {
		return
	}

	l.setFreeFloating(v, freefloating.Dollar, []freefloating.String{
		{
			StringType: freefloating.TokenType,
			Value:      "$",
			Position: &position.Position{
				StartLine: v.GetPosition().StartLine,
				EndLine:   v.GetPosition().StartLine,
				StartPos:  v.GetPosition().StartPos,
				EndPos:    v.GetPosition().StartPos + 1,
			},
		},
	})
}

func (l *Parser) splitSemiColonAndPhpCloseTag(htmlNode node.Node, prevNode node.Node) {
	if l.Lexer.WithFreeFloating == false {
		return
	}

	semiColon := (*prevNode.GetFreeFloating())[freefloating.SemiColon]
	delete((*prevNode.GetFreeFloating()), freefloating.SemiColon)
	if len(semiColon) == 0 {
		return
	}

	p := semiColon[0].Position
	if semiColon[0].Value[0] == ';' {
		l.setFreeFloating(prevNode, freefloating.SemiColon, []freefloating.String{
			{
				StringType: freefloating.TokenType,
				Value:      ";",
				Position: &position.Position{
					StartLine: p.StartLine,
					EndLine:   p.StartLine,
					StartPos:  p.StartPos,
					EndPos:    p.StartPos + 1,
				},
			},
		})
	}

	vlen := len(semiColon[0].Value)
	tlen := 2
	if strings.HasSuffix(semiColon[0].Value, "?>\n") {
		tlen = 3
	}

	phpCloseTag := []freefloating.String{}
	if vlen-tlen > 1 {
		phpCloseTag = append(phpCloseTag, freefloating.String{
			StringType: freefloating.WhiteSpaceType,
			Value:      semiColon[0].Value[1 : vlen-tlen],
			Position: &position.Position{
				StartLine: p.StartLine,
				EndLine:   p.EndLine,
				StartPos:  p.StartPos + 1,
				EndPos:    p.EndPos - tlen,
			},
		})
	}

	phpCloseTag = append(phpCloseTag, freefloating.String{
		StringType: freefloating.WhiteSpaceType,
		Value:      semiColon[0].Value[vlen-tlen:],
		Position: &position.Position{
			StartLine: p.EndLine,
			EndLine:   p.EndLine,
			StartPos:  p.EndPos - tlen,
			EndPos:    p.EndPos,
		},
	})

	l.setFreeFloating(htmlNode, freefloating.Start, append(phpCloseTag, (*htmlNode.GetFreeFloating())[freefloating.Start]...))
}

func (p *Parser) returnTokenToPool(yyDollar []yySymType, yyVAL *yySymType) {
	for i := 1; i < len(yyDollar); i++ {
		if yyDollar[i].token != nil {
			p.TokenPool.Put(yyDollar[i].token)
		}
		yyDollar[i].token = nil
	}
	yyVAL.token = nil
}
