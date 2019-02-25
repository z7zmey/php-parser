package scanner

// TokenPool light version of sync.Pool for Token objects
type TokenPool struct {
	pool []*Token
}

// Get returns *Token from pool or creates new object
func (tp *TokenPool) Get() *Token {
	if len(tp.pool) < 1 {
		return new(Token)
	}

	t := tp.pool[len(tp.pool)-1]
	tp.pool = tp.pool[:len(tp.pool)-1]
	return t
}

// Put returns *Token to pool
func (tp *TokenPool) Put(t *Token) {
	tp.pool = append(tp.pool, t)
}
