# Visitor and Patch

The [ast](https://pkg.go.dev/github.com/anto-lang/anto/ast?tab=doc) package 
provides the `ast.Visitor` interface and the `ast.Walk` function. It can be
used to customize the AST before compiling.

For example, to collect all variable names:

```go
package main

import (
	"fmt"

	"github.com/anto-lang/anto/ast"
	"github.com/anto-lang/anto/parser"
)

type visitor struct {
	identifiers []string
}

func (v *visitor) Visit(node *ast.Node) {
	if n, ok := (*node).(*ast.IdentifierNode); ok {
		v.identifiers = append(v.identifiers, n.Value)
	}
}

func main() {
	tree, err := parser.Parse("foo + bar")
	if err != nil {
		panic(err)
	}

	visitor := &visitor{}
	ast.Walk(&tree.Node, visitor)

	fmt.Printf("%v", visitor.identifiers) // outputs [foo bar]
}
```

## Patch

Specify a visitor to modify the AST with `anto.Patch` function.  

```go
program, err := anto.Compile(code, anto.Patch(&visitor{}))
```
 
For example, let's pass a context to every function call:

```go
package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/anto-lang/anto"
	"github.com/anto-lang/anto/ast"
)

func main() {
	env := map[string]interface{}{
		"foo": func(ctx context.Context, arg int) any {
			// ...
		},
		"ctx": context.Background(),
	}

	code := `foo(42)` // will be converted to foo(ctx, 42)

	program, err := anto.Compile(code, anto.Env(env), anto.Patch(patcher{}))
	if err != nil {
		panic(err)
	}

	output, err := anto.Run(program, env)
	if err != nil {
		panic(err)
	}
	fmt.Print(output)
}

type patcher struct{}

var contextType = reflect.TypeOf((*context.Context)(nil)).Elem()

func (patcher) Visit(node *ast.Node) {
	callNode, ok := (*node).(*ast.CallNode)
	if !ok {
		return
	}
	callNode.Arguments = append([]ast.Node{&ast.IdentifierNode{Value: "ctx"}}, callNode.Arguments...)
}

```
