package anto_test

import (
	"testing"

	"github.com/anto-lang/anto"
	"github.com/anto-lang/anto/vm"
	"github.com/stretchr/testify/require"
)

func Benchmark_expr(b *testing.B) {
	params := make(map[string]any)
	params["Origin"] = "MOW"
	params["Country"] = "RU"
	params["Adults"] = 1
	params["Value"] = 100

	program, err := anto.Compile(`(Origin == "MOW" || Country == "RU") && (Value >= 100 || Adults == 1)`, anto.Env(params))
	require.NoError(b, err)

	var out any

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, params)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.True(b, out.(bool))
}

func Benchmark_expr_reuseVm(b *testing.B) {
	params := make(map[string]any)
	params["Origin"] = "MOW"
	params["Country"] = "RU"
	params["Adults"] = 1
	params["Value"] = 100

	program, err := anto.Compile(`(Origin == "MOW" || Country == "RU") && (Value >= 100 || Adults == 1)`, anto.Env(params))
	require.NoError(b, err)

	var out any
	v := vm.VM{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = v.Run(program, params)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.True(b, out.(bool))
}

func Benchmark_len(b *testing.B) {
	env := map[string]any{
		"arr": make([]int, 100),
	}

	program, err := anto.Compile(`len(arr)`, anto.Env(env))
	require.NoError(b, err)

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.Equal(b, 100, out)
}

func Benchmark_filter(b *testing.B) {
	type Env struct {
		Ints []int
	}
	env := Env{
		Ints: make([]int, 1000),
	}
	for i := 1; i <= len(env.Ints); i++ {
		env.Ints[i-1] = i
	}

	program, err := anto.Compile(`filter(Ints, # % 7 == 0)`, anto.Env(Env{}))
	require.NoError(b, err)

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.Len(b, out.([]any), 142)
}

func Benchmark_filterLen(b *testing.B) {
	type Env struct {
		Ints []int
	}
	env := Env{
		Ints: make([]int, 1000),
	}
	for i := 1; i <= len(env.Ints); i++ {
		env.Ints[i-1] = i
	}

	program, err := anto.Compile(`len(filter(Ints, # % 7 == 0))`, anto.Env(Env{}))
	require.NoError(b, err)

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.Equal(b, 142, out)
}

func Benchmark_filterFirst(b *testing.B) {
	type Env struct {
		Ints []int
	}
	env := Env{
		Ints: make([]int, 1000),
	}
	for i := 1; i <= len(env.Ints); i++ {
		env.Ints[i-1] = i
	}

	program, err := anto.Compile(`filter(Ints, # % 7 == 0)[0]`, anto.Env(Env{}))
	require.NoError(b, err)

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.Equal(b, 7, out)
}

func Benchmark_filterLast(b *testing.B) {
	type Env struct {
		Ints []int
	}
	env := Env{
		Ints: make([]int, 1000),
	}
	for i := 1; i <= len(env.Ints); i++ {
		env.Ints[i-1] = i
	}

	program, err := anto.Compile(`filter(Ints, # % 7 == 0)[-1]`, anto.Env(Env{}))
	require.NoError(b, err)

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}

	b.StopTimer()

	require.NoError(b, err)
	require.Equal(b, 994, out)
}

func Benchmark_filterMap(b *testing.B) {
	type Env struct {
		Ints []int
	}
	env := Env{
		Ints: make([]int, 100),
	}
	for i := 1; i <= len(env.Ints); i++ {
		env.Ints[i-1] = i
	}

	program, err := anto.Compile(`map(filter(Ints, # % 7 == 0), # * 2)`, anto.Env(Env{}))
	require.NoError(b, err)

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.Len(b, out.([]any), 14)
	require.Equal(b, 14, out.([]any)[0])
}

func Benchmark_arrayIndex(b *testing.B) {
	env := map[string]any{
		"arr": make([]int, 100),
	}
	for i := 0; i < 100; i++ {
		env["arr"].([]int)[i] = i
	}

	program, err := anto.Compile(`arr[50]`, anto.Env(env))
	require.NoError(b, err)

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.Equal(b, 50, out)
}

func Benchmark_envStruct(b *testing.B) {
	type Price struct {
		Value int
	}
	type Env struct {
		Price Price
	}

	program, err := anto.Compile(`Price.Value > 0`, anto.Env(Env{}))
	require.NoError(b, err)

	env := Env{Price: Price{Value: 1}}

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.True(b, out.(bool))
}

func Benchmark_envMap(b *testing.B) {
	type Price struct {
		Value int
	}
	env := map[string]any{
		"price": Price{Value: 1},
	}

	program, err := anto.Compile(`price.Value > 0`, anto.Env(env))
	require.NoError(b, err)

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.True(b, out.(bool))
}

type CallEnv struct {
	A      int
	B      int
	C      int
	Fn     func() bool
	FnFast func(...any) any
	Foo    CallFoo
}

func (CallEnv) Func() string {
	return "func"
}

type CallFoo struct {
	D int
	E int
	F int
}

func (CallFoo) Method() string {
	return "method"
}

func Benchmark_callFunc(b *testing.B) {
	program, err := anto.Compile(`Func()`, anto.Env(CallEnv{}))
	require.NoError(b, err)

	env := CallEnv{}

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.Equal(b, "func", out)
}

func Benchmark_callMethod(b *testing.B) {
	program, err := anto.Compile(`Foo.Method()`, anto.Env(CallEnv{}))
	require.NoError(b, err)

	env := CallEnv{}

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.Equal(b, "method", out)
}

func Benchmark_callField(b *testing.B) {
	program, err := anto.Compile(`Fn()`, anto.Env(CallEnv{}))
	require.NoError(b, err)

	env := CallEnv{
		Fn: func() bool {
			return true
		},
	}

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.True(b, out.(bool))
}

func Benchmark_callFast(b *testing.B) {
	program, err := anto.Compile(`FnFast()`, anto.Env(CallEnv{}))
	if err != nil {
		b.Fatal(err)
	}

	env := CallEnv{
		FnFast: func(s ...any) any {
			return "fn_fast"
		},
	}

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.Equal(b, "fn_fast", out)
}

func Benchmark_callConstExpr(b *testing.B) {
	program, err := anto.Compile(`Func()`, anto.Env(CallEnv{}), anto.ConstExpr("Func"))
	require.NoError(b, err)

	env := CallEnv{}

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.Equal(b, "func", out)
}

func Benchmark_largeStructAccess(b *testing.B) {
	type Env struct {
		Data  [1024 * 1024 * 10]byte
		Field int
	}

	program, err := anto.Compile(`Field > 0 && Field > 1 && Field < 99`, anto.Env(Env{}))
	require.NoError(b, err)

	env := Env{Field: 21}

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, &env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.True(b, out.(bool))
}

func Benchmark_largeNestedStructAccess(b *testing.B) {
	type Env struct {
		Inner struct {
			Data  [1024 * 1024 * 10]byte
			Field int
		}
	}

	program, err := anto.Compile(`Inner.Field > 0 && Inner.Field > 1 && Inner.Field < 99`, anto.Env(Env{}))
	require.NoError(b, err)

	env := Env{}
	env.Inner.Field = 21

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, &env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.True(b, out.(bool))
}

func Benchmark_largeNestedArrayAccess(b *testing.B) {
	type Env struct {
		Data [1][1024 * 1024 * 10]byte
	}

	program, err := anto.Compile(`Data[0][0] > 0`, anto.Env(Env{}))
	require.NoError(b, err)

	env := Env{}
	env.Data[0][0] = 1

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, &env)
	}
	b.StopTimer()

	require.NoError(b, err)
	require.True(b, out.(bool))
}

func Benchmark_sort(b *testing.B) {
	env := map[string]any{
		"arr": []any{55, 58, 42, 61, 75, 52, 64, 62, 16, 79, 40, 14, 50, 76, 23, 2, 5, 80, 89, 51, 21, 96, 91, 13, 71, 82, 65, 63, 11, 17, 94, 81, 74, 4, 97, 1, 39, 3, 28, 8, 84, 90, 47, 85, 7, 56, 49, 93, 33, 12, 19, 60, 86, 100, 44, 45, 36, 72, 95, 77, 34, 92, 24, 73, 18, 38, 43, 26, 41, 69, 67, 57, 9, 27, 66, 87, 46, 35, 59, 70, 10, 20, 53, 15, 32, 98, 68, 31, 54, 25, 83, 88, 22, 48, 29, 37, 6, 78, 99, 30},
	}

	program, err := anto.Compile(`sort(arr)`, anto.Env(env))
	require.NoError(b, err)

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ = vm.Run(program, env)
	}
	b.StopTimer()

	require.Equal(b, 1, out.([]any)[0])
	require.Equal(b, 100, out.([]any)[99])
}

func Benchmark_sortBy(b *testing.B) {
	type Foo struct {
		Value int
	}
	arr := []any{55, 58, 42, 61, 75, 52, 64, 62, 16, 79, 40, 14, 50, 76, 23, 2, 5, 80, 89, 51, 21, 96, 91, 13, 71, 82, 65, 63, 11, 17, 94, 81, 74, 4, 97, 1, 39, 3, 28, 8, 84, 90, 47, 85, 7, 56, 49, 93, 33, 12, 19, 60, 86, 100, 44, 45, 36, 72, 95, 77, 34, 92, 24, 73, 18, 38, 43, 26, 41, 69, 67, 57, 9, 27, 66, 87, 46, 35, 59, 70, 10, 20, 53, 15, 32, 98, 68, 31, 54, 25, 83, 88, 22, 48, 29, 37, 6, 78, 99, 30}
	env := map[string]any{
		"arr": make([]Foo, len(arr)),
	}
	for i, v := range arr {
		env["arr"].([]Foo)[i] = Foo{Value: v.(int)}
	}

	program, err := anto.Compile(`sortBy(arr, "Value")`, anto.Env(env))
	require.NoError(b, err)

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ = vm.Run(program, env)
	}
	b.StopTimer()

	require.Equal(b, 1, out.([]any)[0].(Foo).Value)
	require.Equal(b, 100, out.([]any)[99].(Foo).Value)
}

func Benchmark_groupBy(b *testing.B) {
	program, err := anto.Compile(`groupBy(1..100, # % 7)[6]`)
	require.NoError(b, err)

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ = vm.Run(program, nil)
	}
	b.StopTimer()

	require.Equal(b, 6, out.([]any)[0])
}

func Benchmark_reduce(b *testing.B) {
	program, err := anto.Compile(`reduce(1..100, # + #acc)`)
	require.NoError(b, err)

	var out any
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ = vm.Run(program, nil)
	}
	b.StopTimer()

	require.Equal(b, 5050, out.(int))
}
