package position

// Position represents node position
type Position struct {
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
}

// NewPosition Position constructor
func NewPosition(StartLine int, EndLine int, StartPos int, EndPos int) *Position {
	return &Position{
		StartLine: StartLine,
		EndLine:   EndLine,
		StartPos:  StartPos,
		EndPos:    EndPos,
	}
}
