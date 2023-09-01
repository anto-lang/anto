package operator_test

import (
	"testing"
	"time"

	"github.com/anto-lang/anto"
	"github.com/anto-lang/anto/test/mock"
	"github.com/stretchr/testify/require"
)

func TestOperator_struct(t *testing.T) {
	env := mock.Env{
		Time: time.Date(2017, time.October, 23, 18, 30, 0, 0, time.UTC),
	}

	code := `Time == "2017-10-23"`

	program, err := anto.Compile(code, anto.Env(mock.Env{}), anto.Operator("==", "TimeEqualString"))
	require.NoError(t, err)

	output, err := anto.Run(program, env)
	require.NoError(t, err)
	require.Equal(t, true, output)
}

func TestOperator_options_another_order(t *testing.T) {
	code := `Time == "2017-10-23"`
	_, err := anto.Compile(code, anto.Operator("==", "TimeEqualString"), anto.Env(mock.Env{}))
	require.NoError(t, err)
}

func TestOperator_no_env(t *testing.T) {
	code := `Time == "2017-10-23"`
	require.Panics(t, func() {
		_, _ = anto.Compile(code, anto.Operator("==", "TimeEqualString"))
	})
}

func TestOperator_interface(t *testing.T) {
	env := mock.Env{}

	code := `Foo == "Foo.String" && "Foo.String" == Foo && Time != Foo && Time == Time`

	program, err := anto.Compile(
		code,
		anto.Env(mock.Env{}),
		anto.Operator("==", "StringerStringEqual", "StringStringerEqual", "StringerStringerEqual"),
		anto.Operator("!=", "NotStringerStringEqual", "NotStringStringerEqual", "NotStringerStringerEqual"),
	)
	require.NoError(t, err)

	output, err := anto.Run(program, env)
	require.NoError(t, err)
	require.Equal(t, true, output)
}
