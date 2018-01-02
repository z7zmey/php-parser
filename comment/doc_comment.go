package comment

type DocComment struct {
	value string
}

func NewDocComment(value string) Comment {
	return &DocComment{
		value,
	}
}

func (c *DocComment) String() string {
	return c.value
}
