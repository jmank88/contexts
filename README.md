# Contexts [![GoDoc](https://godoc.org/github.com/jmank88/contexts?status.svg)](https://godoc.org/github.com/jmank88/contexts) [![Go Report Card](https://goreportcard.com/badge/github.com/jmank88/contexts)](https://goreportcard.com/report/github.com/jmank88/contexts)
A Go package for creating multi-value contexts.

```
ctx := contexts.WithValues(context.Background(), map[interface{}]interface{} {
    key1: "value 1",
    key2: 100,
    key3: "a third value",
})
```

## Rationale

The standard context package provides a WithValue function for generating a new
context with the parent's values and a single new key-value pair. When adding
multiple values, this approach has a few potential problems:
- Multiple calls must be made in order to add more than one value, creating a
series of linked contexts
- Looking up a value involves traversing the linked series of parent contexts
until that value is found - an O(n) operation

This package provides a WithValues() function for generating contexts with
multiple values. This approach circumvents the aforementioned problems:
- A single call creates a single context holding a map of values
- Looking up a value reads from the map - a constant time operation

## Benchmarks

The following benchmarks measure looking up the first context value from:
- a standard context constructed via WithValue
- a multi-value context constructed via WithValues
```
BenchmarkStdValue1     	50000000	        38.3 ns/op
BenchmarkStdValue2     	30000000	        47.1 ns/op
BenchmarkStdValue4     	20000000	        64.1 ns/op
BenchmarkStdValue8     	20000000	       108 ns/op
BenchmarkStdValue16    	10000000	       187 ns/op
BenchmarkStdValue32    	 5000000	       330 ns/op

BenchmarkValue1        	20000000	        58.2 ns/op
BenchmarkValue2        	30000000	        56.8 ns/op
BenchmarkValue4        	30000000	        56.8 ns/op
BenchmarkValue8        	30000000	        58.3 ns/op
BenchmarkValue16       	30000000	        56.7 ns/op
BenchmarkValue32       	30000000	        56.8 ns/op
```

As expected, the lookup times for standard contexts continue to grow, while the
multi-valued contexts remain constant.


The following benchmarks measure creation of a context with multiple values:
```
BenchmarkStdWithValue1 	20000000	       108 ns/op
BenchmarkStdWithValue2 	10000000	       216 ns/op
BenchmarkStdWithValue4 	 3000000	       438 ns/op
BenchmarkStdWithValue8 	 2000000	       899 ns/op
BenchmarkStdWithValue16	 1000000	      1732 ns/op
BenchmarkStdWithValue32	  500000	      3405 ns/op

BenchmarkWithValue1    	10000000	       172 ns/op
BenchmarkWithValue2    	 5000000	       287 ns/op
BenchmarkWithValue4    	 3000000	       558 ns/op
BenchmarkWithValue8    	 1000000	      1064 ns/op
BenchmarkWithValue16   	  500000	      3659 ns/op
BenchmarkWithValue32   	  200000	      7596 ns/op
```
Creating a standard context is faster than creating a multi-valued context when
a new value map must be allocated and populated. However, in some cases the map
of values may persist and be used as-is in the multi-valued context:
```
BenchmarkWithValuePersist1 	200000000	         6.94 ns/op
BenchmarkWithValuePersist2 	200000000	         6.94 ns/op
BenchmarkWithValuePersist4 	200000000	         6.94 ns/op
BenchmarkWithValuePersist8 	200000000	         6.94 ns/op
BenchmarkWithValuePersist16	200000000	         6.95 ns/op
BenchmarkWithValuePersist32	200000000	         6.95 ns/op
```
Using a pre-allocated and populated map of values dramatically decreases
context creation time.


These results indicate that a sweet-spot use case might be something like an
unmodifiable (or rarely modified) set of ~4 or more system wide values.
For example, a server which initializes some identifying values on start-up
(such as Name, DataCenter, A/B Test-Group, etc), and includes them on every
request context. Based on our benchmarks, with 4 context values we expect value
lookups on par with standard contexts (~50-60ns on my machine), and much faster
context creation (7ns vs. 438ns on my machine). The benefits improve as the
number of context values increases, since the multi-value context's numbers
remain constant while the standard context's increase.
