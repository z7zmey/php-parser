package meta

// Type TODO
type Type int

//go:generate stringer -type=Type -output ./type_string.go
const (
	CommentType Type = iota
	WhiteSpaceType
	TokenType
)
