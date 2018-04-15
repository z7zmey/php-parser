package comment

// Comment aggrigates information about comment /**
type Comment struct {
	value string
}

// NewComment - Comment constructor
func NewComment(value string) *Comment {
	return &Comment{
		value,
	}
}

func (c *Comment) String() string {
	return c.value
}
