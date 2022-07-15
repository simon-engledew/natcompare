A basic (zero-dependency) natural comparison function for Go.

```go
package main

import (
    "golang.org/x/exp/slices"
    "github.com/simon-engledew/natcompare"
)

func main() {
	values := []string{
		"2000-3-23",
		"2000-1-2",
		"1999-3-3",
		"2000-1-10",
		"1999-12-25",
	}
	
	slices.SortFunc(values, natcompare.Less)
}
```
