package meta

import (
	"github.com/z7zmey/php-parser/position"
)

type Meta interface {
	String() string
	SetTokenName(tn TokenName)
	GetTokenName() TokenName
	GetPosition() *position.Position
}
