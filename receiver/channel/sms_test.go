package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSMS_AddUID(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	s := NewSMS(attributes)
	s.AddUID(3)
	assert.Len(t, s.GetAttributes(), 3)
}

func TestSMS_AddUIDs(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	s := NewSMS(attributes)
	uids := UidsType{
		3,
		4,
	}
	s.AddUIDs(uids, false)
	assert.Len(t, s.GetAttributes(), 4)
}

func TestSMS_AddUIDsWithReplace(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	s := NewSMS(attributes)
	uids := UidsType{
		3,
		4,
	}
	s.AddUIDs(uids, true)
	assert.Len(t, s.GetAttributes(), 2)
}

func TestSMS_GetChannel(t *testing.T) {
	attributes := UidsType{}
	s := NewSMS(attributes)
	assert.Equal(t, defaultSMSChannelName, s.GetChannel())
}

func TestSMS_GetAttributes(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	s := NewSMS(attributes)
	assert.Len(t, s.GetAttributes(), 2)
}
