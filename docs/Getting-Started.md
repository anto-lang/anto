# Getting Started

**Expr** is a simple, fast and extensible expression language for Go. It is
designed to be easy to use and integrate into your Go application. Let's delve
deeper into its core features:

- **Memory safe** - Designed to prevent vulnerabilities like buffer overflows and memory leaks.
- **Type safe** - Enforces strict type rules, aligning with Go's type system.
- **Terminating** - Ensures every expression evaluation cannot loop indefinitely.
- **Side effects free** - Evaluations won't modify global states or variables.

Let's start with a simple example:

```go
program, err := anto.Compile(`2 + 2`)
if err != nil {
    panic(err)
}

output, err := anto.Run(program, nil)
if err != nil {
    panic(err)
}

fmt.Print(output) // 4
```

Expr compiles the expression `2 + 2` into a bytecode program. Then we run
the program and get the output. The output is `4` as expected.

The `anto.Compile` function returns a `*vm.Program` and an error. The program
can be reused between runs. The `anto.Run` function takes a program and an
environment. The environment is a map of variables that can be used in the
expression. In this example, we use `nil` as an environment because we don't
need any variables.

Now let's pass some variables to the expression:

```go
env := map[string]any{
    "foo": 100,
	"bar": 200,
}

program, err := anto.Compile(`foo + bar`, anto.Env(env))
if err != nil {
    panic(err)
}

output, err := anto.Run(program, env)
if err != nil {
    panic(err)
}

fmt.Print(output) // 300
```

Why do we need to pass the environment to the `anto.Compile` function? Expr can be used as a type-safe language. 
Expr can infer the type of the expression and check it against the environment. Here is an example:

```go
env := map[string]any{
    "name": "Anton",
    "age": 35,
}

program, err := anto.Compile(`name + age`, anto.Env(env))
if err != nil {
    panic(err) // Will panic with "invalid operation: string + int"
}
```

Expr can work with any Go types. Here is an example:

```go
env := map[string]interface{}{
	"greet":   "Hello, %v!",
	"names":   []string{"world", "you"},
	"sprintf": fmt.Sprintf,
}

code := `sprintf(greet, names[0])`

program, err := anto.Compile(code, anto.Env(env))
if err != nil {
	panic(err)
}

output, err := anto.Run(program, env)
if err != nil {
	panic(err)
}

fmt.Print(output) // Hello, world!
```

Also, Expr can use a struct as an environment. Methods defined on the struct become functions.
The struct fields can be renamed with the `expr` tag. Here is an example:

```go
type Env struct {
	Posts []Post `expr:"posts"`
}

func (Env) Format(t time.Time) string { 
	return t.Format(time.RFC822) 
}

type Post struct {
	Body string
	Date time.Time
}

func main() {
	code := `map(posts, Format(.Date) + ": " + .Body)`
	
	program, err := anto.Compile(code, anto.Env(Env{}))
	if err != nil {
		panic(err)
	}

	env := Env{
		Posts: []Post{
			{"Oh My God!", time.Now()}, 
			{"How you doin?", time.Now()}, 
			{"Could I be wearing any more clothes?", time.Now()},
		},
	}

	output, err := anto.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Print(output)
}
```

The compiled program can be reused between runs. Here is an example:

```go
type Env struct {
    X int
    Y int
}

program, err := anto.Compile(`X + Y`, anto.Env(Env{}))
if err != nil {
    panic(err)
}

output, err := anto.Run(program, Env{1, 2})
if err != nil {
    panic(err)
}

fmt.Print(output) // 3

output, err = anto.Run(program, Env{3, 4})
if err != nil {
    panic(err)
}

fmt.Print(output) // 7
```

:::tip
For one-off expressions, you can use the `anto.Eval` function. It compiles and runs the expression in one step.
```go
output, err := anto.Eval(`2 + 2`, env)
```
:::
