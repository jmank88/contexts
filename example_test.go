package contexts_test

import (
	"fmt"

	"golang.org/x/net/context"
	"github.com/jmank88/contexts"
)

type ctxKey int

const (
	key1 ctxKey = iota
	key2
	key3
)

func Example() {
	ctx := context.Background()
	ctx = contexts.WithValues(ctx, map[interface{}]interface{} {
		key1: "value 1",
		key2: 100,
		key3: "third value",
	})

	fmt.Println(ctx.Value(key1))
	fmt.Println(ctx.Value(key2))
	fmt.Println(ctx.Value(key3))

	// Output:
	// value 1
	// 100
	// third value
}