package php7

import (
	"github.com/z7zmey/php-parser/ast/linear"
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/scanner"
)

func (lval *yySymType) Token(t *scanner.Token) {
	lval.token = t
}

// Parser structure
type Parser struct {
	yyParserImpl
	Lexer        scanner.Scanner
	ast          *linear.AST
	list         stackedNodeList
	currentToken *scanner.Token
	rootNode     node.Node
	tokens       []linear.Token
}

// NewParser creates and returns new Parser
func NewParser(src []byte) *Parser {
	lexer := scanner.NewLexer(src)
	return &Parser{
		yyParserImpl{},
		lexer,
		nil,
		stackedNodeList{},
		nil,
		nil,
		[]linear.Token{},
	}
}

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

	l.Lexer.AddError(errors.NewError(msg, pos))
}

// Parse the php7 Parser entrypoint
func (l *Parser) Parse(src []byte, a *linear.AST) int {
	l.Lexer.Reset(src)
	l.list.Reset()

	a.FileData = src

	l.ast = a
	l.Lexer.SetErrors(nil)
	l.rootNode = nil

	// parse

	return l.yyParserImpl.Parse(l)
}

// GetRootNode returns root node
func (l *Parser) GetRootNode() node.Node {
	return l.rootNode
}

// GetErrors returns errors list
func (l *Parser) GetErrors() []*errors.Error {
	return l.Lexer.GetErrors()
}

// helpers

// func (l *Parser) splitSemiColonAndPhpCloseTag(htmlNode node.Node, prevNode node.Node) {
// 	semiColon := (*prevNode.GetFreeFloating())[freefloating.SemiColon]
// 	delete((*prevNode.GetFreeFloating()), freefloating.SemiColon)
// 	if len(semiColon) == 0 {
// 		return
// 	}

// 	p := semiColon[0].Position
// 	if semiColon[0].Value[0] == ';' {
// 		l.setFreeFloating(prevNode, freefloating.SemiColon, []freefloating.String{
// 			{
// 				StringType: freefloating.TokenType,
// 				Value:      ";",
// 				Position: &position.Position{
// 					StartLine: p.StartLine,
// 					EndLine:   p.StartLine,
// 					StartPos:  p.StartPos,
// 					EndPos:    p.StartPos + 1,
// 				},
// 			},
// 		})
// 	}

// 	vlen := len(semiColon[0].Value)
// 	tlen := 2
// 	if strings.HasSuffix(semiColon[0].Value, "?>\n") {
// 		tlen = 3
// 	}

// 	phpCloseTag := []freefloating.String{}
// 	if vlen-tlen > 1 {
// 		phpCloseTag = append(phpCloseTag, freefloating.String{
// 			StringType: freefloating.WhiteSpaceType,
// 			Value:      semiColon[0].Value[1 : vlen-tlen],
// 			Position: &position.Position{
// 				StartLine: p.StartLine,
// 				EndLine:   p.EndLine,
// 				StartPos:  p.StartPos + 1,
// 				EndPos:    p.EndPos - tlen,
// 			},
// 		})
// 	}

// 	phpCloseTag = append(phpCloseTag, freefloating.String{
// 		StringType: freefloating.WhiteSpaceType,
// 		Value:      semiColon[0].Value[vlen-tlen:],
// 		Position: &position.Position{
// 			StartLine: p.EndLine,
// 			EndLine:   p.EndLine,
// 			StartPos:  p.EndPos - tlen,
// 			EndPos:    p.EndPos,
// 		},
// 	})

// 	l.setFreeFloating(htmlNode, freefloating.Start, append(phpCloseTag, (*htmlNode.GetFreeFloating())[freefloating.Start]...))
// }

func (p *Parser) returnTokenToPool(yyDollar []yySymType, yyVAL *yySymType) {
	for i := 1; i < len(yyDollar); i++ {
		if yyDollar[i].token != nil {
			p.Lexer.ReturnTokenToPool(yyDollar[i].token)
		}
		yyDollar[i].token = nil
	}
	yyVAL.token = nil
}
