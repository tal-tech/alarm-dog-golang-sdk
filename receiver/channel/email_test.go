package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail_AddUID(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	e := NewEmail(attributes)
	e.AddUID(3)
	assert.Len(t, e.GetAttributes(), 3)
}

func TestEmail_AddUIDs(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	e := NewEmail(attributes)
	uids := UidsType{
		3,
		4,
	}
	e.AddUIDs(uids, false)
	assert.Len(t, e.GetAttributes(), 4)
}

func TestEmail_AddUIDsWithReplace(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	e := NewEmail(attributes)
	uids := UidsType{
		3,
		4,
	}
	e.AddUIDs(uids, true)
	assert.Len(t, e.GetAttributes(), 2)
}

func TestEmail_GetChannel(t *testing.T) {
	attributes := UidsType{}
	e := NewEmail(attributes)
	assert.Equal(t, defaultEmailChannelName, e.GetChannel())
}

func TestEmail_GetAttributes(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	e := NewEmail(attributes)
	assert.Len(t, e.GetAttributes(), 2)
}
