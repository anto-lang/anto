package main

import (
	"fmt"

	"github.com/anto-lang/anto/docgen"
	"github.com/anto-lang/anto/test/real_world"
)

func main() {
	doc := docgen.CreateDoc(real_world.NewEnv())

	fmt.Println(doc.Markdown())
}
