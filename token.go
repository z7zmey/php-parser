package main

type token struct {
	value     []byte
	startLine int
	endLine   int
}

func newToken(value []byte, startLine int, endLine int) token {
	return token{value, startLine, endLine}
}

func (t token) String() string {
	return string(t.value)
}
