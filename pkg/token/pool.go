package token

const DefaultBlockSize = 1024

type Pool struct {
	block []Token
	off   int
}

func NewPool(blockSize int) *Pool {
	return &Pool{
		block: make([]Token, blockSize),
	}
}

func (p *Pool) Get() *Token {
	if len(p.block) == 0 {
		return nil
	}

	if len(p.block) == p.off {
		p.block = make([]Token, len(p.block))
		p.off = 0
	}

	p.off++

	return &p.block[p.off-1]
}
