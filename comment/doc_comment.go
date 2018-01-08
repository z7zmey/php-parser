package comment

type DocComment struct {
	value string
}

func NewDocComment(value string) *DocComment {
	return &DocComment{
		value,
	}
}

func (c *DocComment) String() string {
	return c.value
}
