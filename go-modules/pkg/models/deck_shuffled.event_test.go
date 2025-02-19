package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidEventDeckShuffledFn(t *testing.T) {
	CleanRegistry()
	event, _ := NewEvent(EventDeckShuffled, map[string]any{"key": "value"})

	// hardcode an invalid type
	event.Type = "Not a valid event type"
	err := EventDeckShuffledFn(event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid event type")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)

	// hardcode an invalid status
	event.Data["status"] = SOEEnqueued
	event.Type = EventDeckShuffled
	err = EventDeckShuffledFn(event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid event status")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)

	// forget to add deck property on purpose!
	event.Data["status"] = SOEProcessing
	event.Type = EventDeckShuffled
	err = EventDeckShuffledFn(event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "deck missing")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)

}
