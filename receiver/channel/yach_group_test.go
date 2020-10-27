package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYachGroup_AddRobot(t *testing.T) {
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
	dg := NewYachGroup(attributes)
	dg.AddRobot("webhook3", "secret2")
	assert.Len(t, dg.GetAttributes(), 3)
}

func TestYachGroup_AddRobots(t *testing.T) {
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
	dg := NewYachGroup(attributes)
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

func TestYachGroup_AddRobotsWithReplace(t *testing.T) {
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
	dg := NewYachGroup(attributes)
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

func TestYachGroup_GetChannel(t *testing.T) {
	attributes := RobotsType{}
	dg := NewYachGroup(attributes)
	assert.Equal(t, defaultYachGroupChannelName, dg.GetChannel())
}

func TestYachGroup_GetAttributes(t *testing.T) {
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
	dg := NewYachGroup(attributes)
	assert.Len(t, dg.GetAttributes(), 2)
}
