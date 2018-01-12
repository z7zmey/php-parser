package comment

// DocComment represents comments that start /**
type DocComment struct {
	value string
}

// NewDocComment - DocComment constructor
func NewDocComment(value string) *DocComment {
	return &DocComment{
		value,
	}
}

func (c *DocComment) String() string {
	return c.value
}
