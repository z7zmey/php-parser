package php7

import (
	"io"
	"strings"

	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/meta"

	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"
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

func (l *Parser) WithMeta() {
	l.Lexer.WithMeta = true
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

func newInheritMetaFilter() meta.Filter {
	return meta.StopOnFailureFilter(
		meta.AndFilter(
			meta.TokenNameFilter(meta.NodeStart),
			meta.OrFilter(
				meta.TypeFilter(meta.CommentType, meta.WhiteSpaceType),
				meta.ValueFilter("<?php", "<?"),
			),
		),
	)
}

func (l *Parser) appendMetaToken(n node.Node, t *scanner.Token, tn meta.TokenName) {
	if !l.Lexer.WithMeta {
		return
	}

	m := &meta.Data{
		Value:     t.Value,
		Type:      meta.TokenType,
		Position:  l.positionBuilder.NewTokenPosition(t),
		TokenName: tn,
	}

	n.GetMeta().Push(m)
}

func (l *Parser) appendMeta(n node.Node, m *meta.Data, tn meta.TokenName) {
	if !l.Lexer.WithMeta {
		return
	}

	n.GetMeta().Push(m)
}

func (l *Parser) prependMetaToken(n node.Node, t *scanner.Token, tn meta.TokenName) {
	if !l.Lexer.WithMeta {
		return
	}

	m := &meta.Data{
		Value:     t.Value,
		Type:      meta.TokenType,
		Position:  l.positionBuilder.NewTokenPosition(t),
		TokenName: tn,
	}

	n.GetMeta().Unshift(m)
}

func (l *Parser) splitSemicolonTokenAndPhpCloseTag(htmlNode node.Node, prevNode node.Node) {
	SemiColonTokenMeta := prevNode.GetMeta().Cut(meta.AndFilter(
		meta.TokenNameFilter(meta.SemiColonToken),
		meta.TypeFilter(meta.TokenType),
	))

	if len(*SemiColonTokenMeta) < 1 {
		return
	}

	metaTokenValue := (*SemiColonTokenMeta)[0].Value

	i := strings.Index(metaTokenValue, "?>")

	if i < 0 {
		SemiColonTokenMeta.AppendTo(prevNode.GetMeta())
	} else {
		if metaTokenValue[0] == ';' {
			prevNode.GetMeta().Push(&meta.Data{
				Value:     metaTokenValue[0:1],
				Type:      meta.TokenType,
				Position:  nil,
				TokenName: meta.SemiColonToken,
			})

			htmlNode.GetMeta().Push(&meta.Data{
				Value:     metaTokenValue[1:i],
				Type:      meta.WhiteSpaceType,
				Position:  nil,
				TokenName: meta.NodeStart,
			})

			htmlNode.GetMeta().Push(&meta.Data{
				Value:     metaTokenValue[i : i+2],
				Type:      meta.TokenType,
				Position:  nil,
				TokenName: meta.NodeStart,
			})

			if len(metaTokenValue) > i+2 {
				htmlNode.GetMeta().Push(&meta.Data{
					Value:     metaTokenValue[i+2:],
					Type:      meta.WhiteSpaceType,
					Position:  nil,
					TokenName: meta.NodeStart,
				})
			}
		} else {
			htmlNode.GetMeta().Push(&meta.Data{
				Value:     metaTokenValue[:2],
				Type:      meta.TokenType,
				Position:  nil,
				TokenName: meta.NodeStart,
			})

			if len(metaTokenValue) > 2 {
				htmlNode.GetMeta().Push(&meta.Data{
					Value:     metaTokenValue[2:],
					Type:      meta.WhiteSpaceType,
					Position:  nil,
					TokenName: meta.NodeStart,
				})
			}
		}
	}
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
