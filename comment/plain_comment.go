package comment

// PlainComment represents comments that dont start /**
type PlainComment struct {
	value string
}

// NewPlainComment - PlainComment constructor
func NewPlainComment(value string) *PlainComment {
	return &PlainComment{
		value,
	}
}

func (c *PlainComment) String() string {
	return c.value
}
