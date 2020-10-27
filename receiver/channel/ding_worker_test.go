package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDingWorker_AddUID(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	dw := NewDingWorker(attributes)
	dw.AddUID(3)
	assert.Len(t, dw.GetAttributes(), 3)
}

func TestDingWorker_AddUIDs(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	dw := NewDingWorker(attributes)
	uids := UidsType{
		3,
		4,
	}
	dw.AddUIDs(uids, false)
	assert.Len(t, dw.GetAttributes(), 4)
}

func TestDingWorker_AddUIDsWithReplace(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	dw := NewDingWorker(attributes)
	uids := UidsType{
		3,
		4,
	}
	dw.AddUIDs(uids, true)
	assert.Len(t, dw.GetAttributes(), 2)
}

func TestDingWorker_GetChannel(t *testing.T) {
	attributes := UidsType{}
	dw := NewDingWorker(attributes)
	assert.Equal(t, defaultDingWorkerChannelName, dw.GetChannel())
}

func TestDingWorker_GetAttributes(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	dw := NewDingWorker(attributes)
	assert.Len(t, dw.GetAttributes(), 2)
}
