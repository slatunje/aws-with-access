package cue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlags(t *testing.T) {

	a := assert.New(t)
	c := []struct {
		name string
		args []string
		expd []string
	}{
		{
			"no args"	,
			[]string{},
			[]string{},
		},
		{
			"single arg"	,
			[]string{
				"\\-v",
			},
			[]string{
				"-v",
			},
		},
		{
			"single arg"	,
			[]string{
				"\\-v",
			},
			[]string{
				"-v",
			},
		},
	}

	for _, tc := range c {
		t.Run(tc.name, func(t *testing.T) {
			a.Equal(tc.expd, flags(tc.args))
		})
	}

	a.True(true)

}
