package node

import "fmt"
import "github.com/z7zmey/php-parser/comment"

type Node interface {
	Positioner
	Commenter
	Attributes() map[string]interface{}
	Walk(v Visitor)
}

type Commenter interface {
	Comments() *[]comment.Comment
	SetComments(*[]comment.Comment) Node
}

type Positioner interface {
	Position() *Position
	SetPosition(p *Position) Node
}

type Position struct {
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
}

func (p Position) String() string {
	return fmt.Sprintf("Pos{Line: %d-%d Pos: %d-%d}", p.StartLine, p.EndLine, p.StartPos, p.EndPos)
}
