package comment

type PlainComment struct {
	value string
}

func NewPlainComment(value string) *PlainComment {
	return &PlainComment{
		value,
	}
}

func (c *PlainComment) String() string {
	return c.value
}
