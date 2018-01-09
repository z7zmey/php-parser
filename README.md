<!--
  Title: PHP Parser
  Description: A Parser for PHP written in Go.
  Author: Slizov Vadim
  -->

# PHP-Parser

[![Exago](https://api.exago.io:443/badge/tests/github.com/z7zmey/php-parser)](https://exago.io/project/github.com/z7zmey/php-parser)

A Parser for PHP written in Go inspired by [Nikic PHP Parser](https://github.com/nikic/PHP-Parser)

Library uses [cznic/golex](https://github.com/cznic/golex) and [goyacc](https://godoc.org/golang.org/x/tools/cmd/goyacc)

## Features:
- Fully support PHP7 syntax (PHP5 in future)
- Abstract syntax tree representation
- Traversing AST

## Roadmap
- [X] Lexer
- [x] PHP 7 syntax analyzer (completely)
- [x] AST nodes
- [x] AST visitor
- [x] AST dumper
- [x] node position
- [x] handling comments
- [ ] Tests
- [ ] Error handling
- [ ] Stabilize api
- [ ] Documentation
- [ ] PHP 5 syntax analyzer
- [ ] Code flow graph
- [ ] Pretty printer
