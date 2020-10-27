package receiver

import (
	"testing"

	"github.com/tal-tech/alarm-dog-golang-sdk/receiver/channel"

	"github.com/stretchr/testify/assert"
)

func TestReceiver_AddAlarmGroup(t *testing.T) {
	r := NewReceiver([]int{}, []channel.Channel{})
	r.AddAlarmGroup(1)
	assert.Len(t, r.AlarmGroups, 1)
}

func TestReceiver_AddAlarmGroups(t *testing.T) {
	alarmGroups := []int{
		1,
		2,
	}
	r := NewReceiver(alarmGroups, []channel.Channel{})
	r.AddAlarmGroups([]int{
		3,
		4,
	}, false)

	assert.Len(t, r.AlarmGroups, 4)
}

func TestReceiver_AddAlarmGroupsWithReplace(t *testing.T) {
	alarmGroups := []int{
		1,
		2,
	}
	r := NewReceiver(alarmGroups, []channel.Channel{})
	r.AddAlarmGroups([]int{
		3,
		4,
	}, true)

	assert.Len(t, r.AlarmGroups, 2)
}

func TestReceiver_AddChannel(t *testing.T) {
	r := NewReceiver([]int{}, []channel.Channel{})
	dg := channel.NewDingGroup(channel.RobotsType{})
	wg := channel.NewDingWorker(channel.UidsType{})

	assert.NotNil(t, r.AddChannel(dg))
	assert.NotNil(t, r.AddChannel(wg))
}

func TestReceiver_AddChannels(t *testing.T) {
	channels := make([]channel.Channel, 2)
	dg := channel.NewDingGroup(channel.RobotsType{})
	wg := channel.NewDingWorker(channel.UidsType{})
	channels[0] = dg
	channels[1] = wg
	r := NewReceiver([]int{}, channels)

	assert.NotNil(t, r.AddChannels(channels, false))
	assert.Len(t, r.Channels, 4)
}

func TestReceiver_AddChannelsWithReplace(t *testing.T) {
	channels := make([]channel.Channel, 2)
	dg := channel.NewDingGroup(channel.RobotsType{})
	wg := channel.NewDingWorker(channel.UidsType{})
	channels[0] = dg
	channels[1] = wg
	r := NewReceiver([]int{}, channels)

	assert.NotNil(t, r.AddChannels(channels, true))
	assert.Len(t, r.Channels, 2)
}

func TestReceiver_ToArray(t *testing.T) {
	channels := make([]channel.Channel, 2)
	dg := channel.NewDingGroup(channel.RobotsType{
		channel.Robot{
			Webhook: "webhook1",
			Secret:  "secret1",
		},
		channel.Robot{
			Webhook: "webhook2",
			Secret:  "secret2",
		},
	})
	wg := channel.NewDingWorker(channel.UidsType{
		1,
		2,
	})
	channels[0] = dg
	channels[1] = wg
	r := NewReceiver([]int{1, 2, 3, 4}, channels)
	toArrayStruct := r.ToArray()
	assert.NotNil(t, toArrayStruct)
	assert.Equal(t, []int{1, 2, 3, 4}, toArrayStruct.AlarmGroup)
	assert.NotNil(t, toArrayStruct.Channels[dg.GetChannel()])
	assert.NotNil(t, toArrayStruct.Channels[wg.GetChannel()])
}
