package contexts

import (
	"testing"

	"golang.org/x/net/context"
)

func BenchmarkValue1(b *testing.B) {
	benchmarkValue(1, b)
}

func BenchmarkValue2(b *testing.B) {
	benchmarkValue(2, b)
}

func BenchmarkValue4(b *testing.B) {
	benchmarkValue(4, b)
}

func BenchmarkValue8(b *testing.B) {
	benchmarkValue(8, b)
}

func BenchmarkValue16(b *testing.B) {
	benchmarkValue(16, b)
}

func BenchmarkValue32(b *testing.B) {
	benchmarkValue(32, b)
}

func benchmarkValue(cnt int, b *testing.B) {
	values := make(map[interface{}]interface{})
	for i := 0; i<cnt; i++ {
		values[i] = i
	}
	ctx := WithValues(context.Background(), values)
	for n := 0; n<b.N; n++ {
		ctx.Value(0)
	}
}

func BenchmarkStdValue1(b *testing.B) {
	benchmarkStdValue(1, b)
}

func BenchmarkStdValue2(b *testing.B) {
	benchmarkStdValue(2, b)
}

func BenchmarkStdValue4(b *testing.B) {
	benchmarkStdValue(4, b)
}

func BenchmarkStdValue8(b *testing.B) {
	benchmarkStdValue(8, b)
}

func BenchmarkStdValue16(b *testing.B) {
	benchmarkStdValue(16, b)
}

func BenchmarkStdValue32(b *testing.B) {
	benchmarkStdValue(32, b)
}

func benchmarkStdValue(cnt int, b *testing.B) {
	ctx := context.Background()
	for i := 0; i<cnt; i++ {
		ctx = context.WithValue(ctx, i, i)
	}
	for n := 0; n<b.N; n++ {
		ctx.Value(0)
	}
}

func BenchmarkWithValue1(b *testing.B) {
	benchmarkWithValue(1, b)
}

func BenchmarkWithValue2(b *testing.B) {
	benchmarkWithValue(2, b)
}

func BenchmarkWithValue4(b *testing.B) {
	benchmarkWithValue(4, b)
}

func BenchmarkWithValue8(b *testing.B) {
	benchmarkWithValue(8, b)
}

func BenchmarkWithValue16(b *testing.B) {
	benchmarkWithValue(16, b)
}

func BenchmarkWithValue32(b *testing.B) {
	benchmarkWithValue(32, b)
}

func benchmarkWithValue(cnt int, b *testing.B) {
	for n := 0; n<b.N; n++ {
		values := make(map[interface{}]interface{})
		for i := 0; i<cnt; i++ {
			values[i] = i
		}
		_ = WithValues(context.Background(), values)
	}
}

func BenchmarkStdWithValue1(b *testing.B) {
	benchmarkStdWithValue(1, b)
}

func BenchmarkStdWithValue2(b *testing.B) {
	benchmarkStdWithValue(2, b)
}

func BenchmarkStdWithValue4(b *testing.B) {
	benchmarkStdWithValue(4, b)
}

func BenchmarkStdWithValue8(b *testing.B) {
	benchmarkStdWithValue(8, b)
}

func BenchmarkStdWithValue16(b *testing.B) {
	benchmarkStdWithValue(16, b)
}

func BenchmarkStdWithValue32(b *testing.B) {
	benchmarkStdWithValue(32, b)
}

func benchmarkStdWithValue(cnt int, b *testing.B) {
	for n := 0; n<b.N; n++ {
		ctx := context.Background()
		for i := 0; i<cnt; i++ {
			ctx = context.WithValue(ctx, i, i)
		}
	}
}


func BenchmarkWithValuePersist1(b *testing.B) {
	benchmarkWithValuePersist(1, b)
}

func BenchmarkWithValuePersist2(b *testing.B) {
	benchmarkWithValuePersist(2, b)
}

func BenchmarkWithValuePersist4(b *testing.B) {
	benchmarkWithValuePersist(4, b)
}

func BenchmarkWithValuePersist8(b *testing.B) {
	benchmarkWithValuePersist(8, b)
}

func BenchmarkWithValuePersist16(b *testing.B) {
	benchmarkWithValuePersist(16, b)
}

func BenchmarkWithValuePersist32(b *testing.B) {
	benchmarkWithValuePersist(32, b)
}

func benchmarkWithValuePersist(cnt int, b *testing.B) {
	values := make(map[interface{}]interface{})
	for i := 0; i<cnt; i++ {
		values[i] = i
	}
	b.ResetTimer()
	for n := 0; n<b.N; n++ {
		_ = WithValues(context.Background(), values)
	}
}