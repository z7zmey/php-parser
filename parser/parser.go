package parser

import (
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/version"
)

// Parser interface
type Parser interface {
	Parse() int
	GetRootNode() node.Node
	GetErrors() []*errors.Error
	WithFreeFloating()
}

func NewParser(src []byte, v string) (Parser, error) {
	var parser Parser

	r, err := version.Compare(v, "7.0")
	if err != nil {
		return nil, err
	}

	if r == -1 {
		parser = php5.NewParser(src, v)
	} else {
		parser = php7.NewParser(src, v)
	}

	return parser, nil
}
