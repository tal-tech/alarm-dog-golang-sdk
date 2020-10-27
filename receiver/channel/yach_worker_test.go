package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYachWorker_AddUID(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	yw := NewYachWorker(attributes)
	yw.AddUID(3)
	assert.Len(t, yw.GetAttributes(), 3)
}

func TestYachWorker_AddUIDs(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	yw := NewYachWorker(attributes)
	uids := UidsType{
		3,
		4,
	}
	yw.AddUIDs(uids, false)
	assert.Len(t, yw.GetAttributes(), 4)
}

func TestYachWorker_AddUIDsWithReplace(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	yw := NewYachWorker(attributes)
	uids := UidsType{
		3,
		4,
	}
	yw.AddUIDs(uids, true)
	assert.Len(t, yw.GetAttributes(), 2)
}

func TestYachWorker_GetChannel(t *testing.T) {
	attributes := UidsType{}
	yw := NewYachWorker(attributes)
	assert.Equal(t, defaultYachWorkerChannelName, yw.GetChannel())
}

func TestYachWorker_GetAttributes(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	yw := NewYachWorker(attributes)
	assert.Len(t, yw.GetAttributes(), 2)
}
