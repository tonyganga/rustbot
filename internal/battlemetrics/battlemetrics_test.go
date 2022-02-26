package battlemetrics

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRustServerFieldsAreValid(t *testing.T) {
	req := GetRustServer("6324892")
	var i int
	var s string

	assert.IsType(t, i, req.Data.Attributes.Rank)
	assert.IsType(t, s, req.Data.Attributes.Name)
	assert.IsType(t, s, req.Data.Attributes.ID)
}

func TestRustMessageFieldsAreValid(t *testing.T) {
	req := GetRustServer("6324892")
	var i int
	var s string
	var time time.Time
	var f float64

	assert.IsType(t, s, req.Data.Attributes.Name)
	assert.IsType(t, s, req.Data.Attributes.Details.RustDescription)
	assert.IsType(t, s, req.Data.Attributes.Details.RustURL)
	assert.IsType(t, s, req.Data.Attributes.Details.RustHeaderimage)
	assert.IsType(t, i, req.Data.Attributes.Rank)
	assert.IsType(t, time, req.Data.Attributes.Details.RustLastWipe)
	assert.IsType(t, i, req.Data.Attributes.Players)
	assert.IsType(t, i, req.Data.Attributes.MaxPlayers)
	assert.IsType(t, i, req.Data.Attributes.Details.RustQueuedPlayers)
	assert.IsType(t, f, req.Data.Attributes.Details.RustFpsAvg)
	assert.IsType(t, i, req.Data.Attributes.Details.RustWorldSize)
	assert.IsType(t, s, req.Data.Attributes.IP)
	assert.IsType(t, i, req.Data.Attributes.Port)
}
