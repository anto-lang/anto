package crowdsec_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/anto-lang/anto"
	"github.com/anto-lang/anto/test/crowdsec"
	"github.com/stretchr/testify/require"
)

func TestCrowdsec(t *testing.T) {
	b, err := os.ReadFile("../../testdata/crowdsec.json")
	require.NoError(t, err)

	var examples []string
	err = json.Unmarshal(b, &examples)
	require.NoError(t, err)

	env := map[string]any{
		"evt": &crowdsec.Event{},
	}

	var opt = []anto.Option{
		anto.Env(env),
	}
	for _, fn := range crowdsec.CustomFunctions {
		opt = append(
			opt,
			anto.Function(
				fn.Name,
				func(params ...any) (any, error) {
					return nil, nil
				},
				fn.Func...,
			),
		)
	}

	for _, line := range examples {
		t.Run(line, func(t *testing.T) {
			_, err = anto.Compile(line, opt...)
			require.NoError(t, err)
		})
	}
}
