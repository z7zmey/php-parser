package node

import (
	"io"
)

type Node interface {
	Name() string
	Print(out io.Writer, indent string)
}
