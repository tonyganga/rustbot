package battlemetrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRustServer(t *testing.T) {
	type test struct {
		input    string
		want     interface{}
		contains string
	}

	tests := []test{
		{input: "6792417", want: "6792417", contains: "UKN"},
		{input: "6324892", want: "6324892", contains: "Rustoria"},
		{input: "3461363", want: "3461363", contains: "Bloo"},
	}

	for _, tc := range tests {
		got := GetRustServer(tc.input)
		assert.NotEmpty(t, got.Data)
		assert.Equal(t, tc.want, got.Data.ID)
		assert.Contains(t, got.Data.Attributes.Name, tc.contains)
	}
}

func TestGetRustServersWithQuery(t *testing.T) {
	type test struct {
		input string
		want  interface{}
	}

	tests := []test{
		{input: "rustoria", want: 0},
		{input: "ukn", want: 0},
		{input: "moose", want: 0},
		{input: "stevious", want: 0},
		{input: "bloo", want: 0},
		{input: "", want: 0},
	}

	for _, tc := range tests {
		got := GetListOfRustServers(tc.input)
		assert.NotEmpty(t, got.Data)
		assert.Greater(t, len(got.Data), tc.want)
	}
}
