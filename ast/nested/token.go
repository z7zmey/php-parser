package nested

import (
	"encoding/json"

	"github.com/z7zmey/php-parser/scanner"
)

type Token struct {
	Type  scanner.TokenType
	Value string
}

type token struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (t Token) MarshalJSON() ([]byte, error) {
	out := token{
		Type:  t.Type.String(),
		Value: t.Value,
	}

	return json.Marshal(out)
}
