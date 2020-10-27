package receiver

import "github.com/tal-tech/alarm-dog-golang-sdk/receiver/channel"

// Receiver struct
type Receiver struct {
	AlarmGroups []int
	Channels    []channel.Channel
}

// ToArrayStruct struct
type ToArrayStruct struct {
	AlarmGroup []int                    `json:"alarmgroup,omitempty"`
	Channels   map[string][]interface{} `json:"channels,omitempty"`
}

// NewReceiver is used to init robots
func NewReceiver(alarmGroups []int, channels []channel.Channel) *Receiver {
	return &Receiver{
		AlarmGroups: alarmGroups,
		Channels:    channels,
	}
}

// AddAlarmGroup used to add groupID
func (r *Receiver) AddAlarmGroup(groupID int) *Receiver {
	r.AlarmGroups = append(r.AlarmGroups, groupID)
	return r
}

// AddAlarmGroups used to add groupID
func (r *Receiver) AddAlarmGroups(groupIDs []int, replace bool) *Receiver {
	if replace {
		r.AlarmGroups = groupIDs
	} else {
		r.AlarmGroups = append(r.AlarmGroups, groupIDs...)
	}
	return r
}

// AddChannel used to add groupID
func (r *Receiver) AddChannel(channel channel.Channel) *Receiver {
	r.Channels = append(r.Channels, channel)
	return r
}

// AddChannels used to add groupID
func (r *Receiver) AddChannels(channels []channel.Channel, replace bool) *Receiver {
	if replace {
		r.Channels = channels
	} else {
		r.Channels = append(r.Channels, channels...)
	}
	return r
}

// ToArray get all channels
func (r *Receiver) ToArray() *ToArrayStruct {
	ret := ToArrayStruct{
		AlarmGroup: r.AlarmGroups,
		Channels:   make(map[string][]interface{}, 0),
	}

	for _, channel := range r.Channels {
		ret.Channels[channel.GetChannel()] = channel.GetAttributes()
	}

	return &ret
}
