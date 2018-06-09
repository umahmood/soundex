# Soundex

Soundex is a go implementation of the Soundex algorithm. 

*"Soundex is a phonetic algorithm for indexing names by sound, as pronounced in English. The goal is for homophones to be encoded to the same representation so that they can be matched despite minor differences in spelling."* - [Wikipedia](https://en.wikipedia.org/wiki/Soundex)

# Installation

> $ go get github.com/umahmood/soundex

# Usage

```
package main

import (
    "fmt"

    "github.com/umahmood/soundex"
)

func main() {
    fmt.Println(soundex.Code("bingo"))
}
```

# Documentation

> http://godoc.org/github.com/umahmood/soundex

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
