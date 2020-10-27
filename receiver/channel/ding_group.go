package channel

const defaultDingGroupChannelName = "dinggroup"

// DingGroup struct
type DingGroup struct {
	Channel    string
	Attributes RobotsType
}

// NewDingGroup is used to init robots
func NewDingGroup(robots RobotsType) *DingGroup {
	return &DingGroup{
		Channel:    defaultDingGroupChannelName,
		Attributes: robots,
	}
}

// AddRobot add a robot
func (dg *DingGroup) AddRobot(webhook string, secret string) *DingGroup {
	robot := Robot{
		Webhook: webhook,
		Secret:  secret,
	}
	dg.Attributes = append(dg.Attributes, robot)
	return dg
}

// AddRobots add robots
func (dg *DingGroup) AddRobots(robots RobotsType, replace bool) *DingGroup {
	if replace {
		dg.Attributes = robots
	} else {
		dg.Attributes = append(dg.Attributes, robots...)
	}
	return dg
}

// GetChannel get channel
func (dg *DingGroup) GetChannel() string {
	return dg.Channel
}

// GetAttributes get attributes
func (dg *DingGroup) GetAttributes() []interface{} {
	ret := make([]interface{}, len(dg.Attributes))
	for i := 0; i < len(dg.Attributes); i++ {
		ret[i] = dg.Attributes[i]

	}
	return ret
}
