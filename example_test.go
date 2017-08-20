package contexts_test

import (
	"fmt"

	"github.com/jmank88/contexts"
	"golang.org/x/net/context"
)

type ctxKey int

const (
	key1 ctxKey = iota
	key2
	key3
)

func Example() {
	ctx := contexts.WithValues(context.Background(), map[interface{}]interface{}{
		key1: "value 1",
		key2: 100,
		key3: "a third value",
	})

	fmt.Println(ctx.Value(key1))
	fmt.Println(ctx.Value(key2))
	fmt.Println(ctx.Value(key3))

	// Output:
	// value 1
	// 100
	// a third value
}
