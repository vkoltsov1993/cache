Go in-memory cache
================================

In cache memory library stores any value in RAM by key
It has three simple methods: Set, Get, Delete
## Example #1

```go
package main

import (
	"fmt"
	"github.com/vkoltsov1993/cache"
)

func main() {
	cache := cache.New()
	cache.Set("key1", 123)
	fmt.Println("Before delete", cache.Get("key1"))
	cache.Delete("key1")
	fmt.Println("After delete", cache.Get("key1"))

}

```
## Example 2 Force update

```go
package main

import (
	"fmt"
	"github.com/vkoltsov1993/cache"
)

func main() {
	cache := cache.New()
	cache.Set("key1", 123)
	fmt.Println("Before delete", cache.Get("key1"))
	cache.Delete("key1")
	fmt.Println("After delete", cache.Get("key1"))

}

```
