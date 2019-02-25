package scanner_test

import (
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/scanner"
)

func TestTokenPoolGetNew(t *testing.T) {
	tp := new(scanner.TokenPool)

	newToken := tp.Get()

	if newToken == nil {
		t.Errorf("*TokenPool.Get() must return new *Token object\n")
	}
}

func TestTokenPoolGetFromPool(t *testing.T) {
	tp := new(scanner.TokenPool)

	expectedToken := &scanner.Token{
		Value: "test",
	}

	tp.Put(expectedToken)

	actualToken := tp.Get()

	if !reflect.DeepEqual(expectedToken, actualToken) {
		t.Errorf("*TokenPool.Put() must return *Token object from pool\n")
	}
}
