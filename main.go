package main

import (
	"bytes"
	"os"
)

const src = `
<?php
namespace Test;

/**
 * Class foo
 */
class foo
{
    
}
`

func main() {
	yyDebug = 0
	yyErrorVerbose = true
	l := newLexer(bytes.NewBufferString(src), os.Stdout, "file.name")
	yyParse(l)
}
