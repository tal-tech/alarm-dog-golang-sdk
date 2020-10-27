package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDingGroup_AddRobot(t *testing.T) {
	attributes := RobotsType{
		Robot{
			Webhook: "webhook1",
			Secret:  "secret1",
		},
		Robot{
			Webhook: "webhook2",
			Secret:  "secret2",
		},
	}
	dg := NewDingGroup(attributes)
	dg.AddRobot("webhook3", "secret2")
	assert.Len(t, dg.GetAttributes(), 3)
}

func TestDingGroup_AddRobots(t *testing.T) {
	attributes := RobotsType{
		Robot{
			Webhook: "webhook1",
			Secret:  "secret1",
		},
		Robot{
			Webhook: "webhook2",
			Secret:  "secret2",
		},
	}
	dg := NewDingGroup(attributes)
	robots := RobotsType{
		Robot{
			Webhook: "webhook3",
			Secret:  "secret3",
		},
		Robot{
			Webhook: "webhook4",
			Secret:  "secret4",
		},
	}
	dg.AddRobots(robots, false)
	assert.Len(t, dg.GetAttributes(), 4)
}

func TestDingGroup_AddRobotsWithReplace(t *testing.T) {
	attributes := RobotsType{
		Robot{
			Webhook: "webhook1",
			Secret:  "secret1",
		},
		Robot{
			Webhook: "webhook2",
			Secret:  "secret2",
		},
	}
	dg := NewDingGroup(attributes)
	robots := RobotsType{
		Robot{
			Webhook: "webhook3",
			Secret:  "secret3",
		},
		Robot{
			Webhook: "webhook4",
			Secret:  "secret4",
		},
	}
	dg.AddRobots(robots, true)
	assert.Len(t, dg.GetAttributes(), 2)
}

func TestDingGroup_GetChannel(t *testing.T) {
	attributes := RobotsType{}
	dg := NewDingGroup(attributes)
	assert.Equal(t, defaultDingGroupChannelName, dg.GetChannel())
}

func TestDingGroup_GetAttributes(t *testing.T) {
	attributes := RobotsType{
		Robot{
			Webhook: "webhook1",
			Secret:  "secret1",
		},
		Robot{
			Webhook: "webhook2",
			Secret:  "secret2",
		},
	}
	dg := NewDingGroup(attributes)
	assert.Len(t, dg.GetAttributes(), 2)
}
