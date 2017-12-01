<?php

namespace z7zmey\Example;

use \Exception;
use z7zmey\Foo\{Bar, function Baz};

abstract class Foo extends Bar implements Buz, Buzz {
    use \z7zmey\_Trait;

    public const CC = 0;

    public function &test(bool $a, string $b = null): ?void {
        
    }
}