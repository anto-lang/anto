# Tips

## Reuse VM

It is possible to reuse a virtual machine between re-runs on the program.
In some cases, it can add a small increase in performance (~10%).

```go
package main

import (
	"fmt"
	"github.com/anto-lang/anto"
	"github.com/anto-lang/anto/vm"
)

func main() {
	env := map[string]interface{}{
		"foo": 1,
		"bar": 2,
	}

	program, err := anto.Compile("foo + bar", anto.Env(env))
	if err != nil {
		panic(err)
	}

	// Reuse this vm instance between runs
	v := vm.VM{}

	out, err := v.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Print(out)
}
```
