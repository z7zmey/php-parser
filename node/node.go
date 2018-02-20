package node

import "github.com/z7zmey/php-parser/walker"

// Node interface
type Node interface {
	walker.Walkable
	Attributes() map[string]interface{} // Attributes returns node attributes as map
}
