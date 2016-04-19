package contexts

import (
	"fmt"

	"golang.org/x/net/context"
)

func WithValues(parent context.Context, values map[interface{}]interface{}) context.Context {
	return &valuesCtx{parent, values}
}

type valuesCtx struct {
	context.Context
	values map[interface{}]interface{}
}

func (c *valuesCtx) String() string {
	return fmt.Sprintf("%v.WithValues(%v)", c.Context, c.values)
}

func (c *valuesCtx) Value(key interface{}) interface{} {
	if val, ok := c.values[key]; ok {
		return val
	}
	return c.Context.Value(key)
}