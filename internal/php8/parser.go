package php8

import (
	"bytes"

	"github.com/z7zmey/php-parser/internal/position"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/conf"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/token"
)

// Parser structure
type Parser struct {
	Lexer          *Lexer
	currentToken   *token.Token
	rootNode       ast.Vertex
	errHandlerFunc func(*errors.Error)
	builder        *position.Builder
}

// NewParser creates and returns new Parser
func NewParser(lexer *Lexer, config conf.Config) *Parser {
	return &Parser{
		Lexer:          lexer,
		errHandlerFunc: config.ErrorHandlerFunc,
		builder:        position.NewBuilder(),
	}
}

func (p *Parser) Lex(lval *yySymType) int {
	t := p.Lexer.Lex()

	p.currentToken = t
	lval.token = t

	return int(t.ID)
}

func (p *Parser) Error(msg string) {
	if p.errHandlerFunc == nil {
		return
	}

	p.errHandlerFunc(errors.NewError(msg, p.currentToken.Position))
}

// Parse the php7 Parser entrypoint
func (p *Parser) Parse() int {
	p.rootNode = nil

	return yyParse(p)
}

// GetRootNode returns root node
func (p *Parser) GetRootNode() ast.Vertex {
	return p.rootNode
}

func (p *Parser) parseNameRelativeToken(t *token.Token) *ast.NameRelative {
	n := &ast.NameRelative{
		Position: p.builder.NewTokenPosition(t),
	}
	s := t.Position.StartPos
	v := t.Value

	// namespace token

	p1 := p.Lexer.positionPool.Get()
	p1.StartLine = t.Position.StartLine
	p1.EndLine = t.Position.EndLine
	p1.StartPos = s
	p1.EndPos = s + 9

	n.NsTkn = &token.Token{
		ID:       token.T_NAMESPACE,
		Value:    v[:9],
		Position: p1,
	}

	s = s + 9
	v = v[9:]

	// ns separator token

	p1 = p.Lexer.positionPool.Get()
	p1.StartLine = t.Position.StartLine
	p1.EndLine = t.Position.EndLine
	p1.StartPos = s
	p1.EndPos = s + 1

	n.NsSeparatorTkn = &token.Token{
		ID:       token.T_NS_SEPARATOR,
		Value:    v[:1],
		Position: p1,
	}

	s = s + 1
	v = v[1:]

	// parts

	for {
		i := bytes.Index(v, []byte("\\"))
		if i < 0 {
			break
		}

		p1 = p.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + i

		p2 := p.Lexer.positionPool.Get()
		*p2 = *p1

		n.Parts = append(n.Parts, &ast.NamePart{
			Position: p1,
			StringTkn: &token.Token{
				ID:           token.T_STRING,
				Value:        v[:i],
				Position:     p2,
				FreeFloating: t.FreeFloating,
			},
			Value: v[:i],
		})
		t.FreeFloating = nil
		s = s + i
		v = v[i:]

		p1 = p.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + 1

		n.SeparatorTkns = append(n.SeparatorTkns, &token.Token{
			ID:       token.T_NS_SEPARATOR,
			Value:    v[:1],
			Position: p1,
		})
		s = s + 1
		v = v[1:]
	}

	// last part

	p1 = p.Lexer.positionPool.Get()
	p1.StartLine = t.Position.StartLine
	p1.EndLine = t.Position.EndLine
	p1.StartPos = s
	p1.EndPos = s + len(v)
	p2 := p.Lexer.positionPool.Get()
	*p2 = *p1

	n.Parts = append(n.Parts, &ast.NamePart{
		Position: p1,
		StringTkn: &token.Token{
			ID:           token.T_STRING,
			Value:        v,
			Position:     p2,
			FreeFloating: t.FreeFloating,
		},
		Value: v,
	})

	return n
}

func (p *Parser) parseNameFullyQualifiedToken(t *token.Token) *ast.NameFullyQualified {
	n := &ast.NameFullyQualified{
		Position: p.builder.NewTokenPosition(t),
	}
	s := t.Position.StartPos
	v := t.Value

	// ns separator token

	p1 := p.Lexer.positionPool.Get()
	p1.StartLine = t.Position.StartLine
	p1.EndLine = t.Position.EndLine
	p1.StartPos = s
	p1.EndPos = s + 1

	n.NsSeparatorTkn = &token.Token{
		ID:       token.T_NS_SEPARATOR,
		Value:    v[:1],
		Position: p1,
	}

	s = s + 1
	v = v[1:]

	// parts

	for {
		i := bytes.Index(v, []byte("\\"))
		if i < 0 {
			break
		}

		p1 = p.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + i

		p2 := p.Lexer.positionPool.Get()
		*p2 = *p1

		n.Parts = append(n.Parts, &ast.NamePart{
			Position: p1,
			StringTkn: &token.Token{
				ID:           token.T_STRING,
				Value:        v[:i],
				Position:     p2,
				FreeFloating: t.FreeFloating,
			},
			Value: v[:i],
		})
		t.FreeFloating = nil
		s = s + i
		v = v[i:]

		p1 = p.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + 1

		n.SeparatorTkns = append(n.SeparatorTkns, &token.Token{
			ID:       token.T_NS_SEPARATOR,
			Value:    v[:1],
			Position: p1,
		})
		s = s + 1
		v = v[1:]
	}

	// last part

	p1 = p.Lexer.positionPool.Get()
	p1.StartLine = t.Position.StartLine
	p1.EndLine = t.Position.EndLine
	p1.StartPos = s
	p1.EndPos = s + len(v)
	p2 := p.Lexer.positionPool.Get()
	*p2 = *p1

	n.Parts = append(n.Parts, &ast.NamePart{
		Position: p1,
		StringTkn: &token.Token{
			ID:           token.T_STRING,
			Value:        v,
			Position:     p2,
			FreeFloating: t.FreeFloating,
		},
		Value: v,
	})

	return n
}

func (p *Parser) parseNameToken(t *token.Token) *ast.Name {
	n := &ast.Name{
		Position: p.builder.NewTokenPosition(t),
	}
	s := t.Position.StartPos
	v := t.Value

	for {
		i := bytes.Index(v, []byte("\\"))
		if i < 0 {
			break
		}

		p1 := p.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + i

		p2 := p.Lexer.positionPool.Get()
		*p2 = *p1

		n.Parts = append(n.Parts, &ast.NamePart{
			Position: p1,
			StringTkn: &token.Token{
				ID:           token.T_STRING,
				Value:        v[:i],
				Position:     p2,
				FreeFloating: t.FreeFloating,
			},
			Value: v[:i],
		})
		t.FreeFloating = nil
		s = s + i
		v = v[i:]

		p1 = p.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + 1

		n.SeparatorTkns = append(n.SeparatorTkns, &token.Token{
			ID:       token.T_NS_SEPARATOR,
			Value:    v[:1],
			Position: p1,
		})
		s = s + 1
		v = v[1:]
	}

	p1 := p.Lexer.positionPool.Get()
	p1.StartLine = t.Position.StartLine
	p1.EndLine = t.Position.EndLine
	p1.StartPos = s
	p1.EndPos = s + len(v)
	p2 := p.Lexer.positionPool.Get()
	*p2 = *p1

	n.Parts = append(n.Parts, &ast.NamePart{
		Position: p1,
		StringTkn: &token.Token{
			ID:           token.T_STRING,
			Value:        v,
			Position:     p2,
			FreeFloating: t.FreeFloating,
		},
		Value: v,
	})

	return n
}

// helpers

func lastNode(nn []ast.Vertex) ast.Vertex {
	if len(nn) == 0 {
		return nil
	}
	return nn[len(nn)-1]
}
