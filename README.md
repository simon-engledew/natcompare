A basic (zero-dependency) natural comparison function for Go.

Sort strings that contain numbers in a way that is more intuitive, e.g:

Dates:

| Natcompare (Low-High) | 1999-3-3 | 1999-12-25 | 2000-1-2 | 2000-1-10 | 2000-3-23 |
| --- | --- | --- | --- | --- | --- |
| Standard (Low-High) | 1999-12-25 | 1999-3-3 | 2000-1-10 | 2000-1-2 | 2000-3-23 |

Filenames:

| Natcompare (Low-High) | 99.doc | 100.doc | 1000.doc |
| --- | --- | --- | --- |
| Standard (Low-High) | 100.doc | 1000.doc | 99.doc |

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
