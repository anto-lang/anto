package patch_test

import (
	"testing"

	"github.com/anto-lang/anto"
	"github.com/anto-lang/anto/ast"
	"github.com/anto-lang/anto/test/mock"
	"github.com/stretchr/testify/require"
)

type lengthPatcher struct{}

func (p *lengthPatcher) Visit(node *ast.Node) {
	switch n := (*node).(type) {
	case *ast.MemberNode:
		if prop, ok := n.Property.(*ast.StringNode); ok && prop.Value == "length" {
			ast.Patch(node, &ast.BuiltinNode{
				Name:      "len",
				Arguments: []ast.Node{n.Node},
			})
		}
	}
}

func TestPatch_length(t *testing.T) {
	program, err := anto.Compile(
		`String.length == 5`,
		anto.Env(mock.Env{}),
		anto.Patch(&lengthPatcher{}),
	)
	require.NoError(t, err)

	env := mock.Env{String: "hello"}
	output, err := anto.Run(program, env)
	require.NoError(t, err)
	require.Equal(t, true, output)
}
