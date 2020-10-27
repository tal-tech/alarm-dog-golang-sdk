package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhone_AddUID(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	p := NewPhone(attributes)
	p.AddUID(3)
	assert.Len(t, p.GetAttributes(), 3)
}

func TestPhone_AddUIDs(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	p := NewPhone(attributes)
	uids := UidsType{
		3,
		4,
	}
	p.AddUIDs(uids, false)
	assert.Len(t, p.GetAttributes(), 4)
}

func TestPhone_AddUIDsWithReplace(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	p := NewPhone(attributes)
	uids := UidsType{
		3,
		4,
	}
	p.AddUIDs(uids, true)
	assert.Len(t, p.GetAttributes(), 2)
}

func TestPhone_GetChannel(t *testing.T) {
	attributes := UidsType{}
	p := NewPhone(attributes)
	assert.Equal(t, defaultPhoneChannelName, p.GetChannel())
}

func TestPhone_GetAttributes(t *testing.T) {
	attributes := UidsType{
		1,
		2,
	}
	p := NewPhone(attributes)
	assert.Len(t, p.GetAttributes(), 2)
}
