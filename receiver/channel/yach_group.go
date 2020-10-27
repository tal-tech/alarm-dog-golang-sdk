package channel

const defaultYachGroupChannelName = "yachgroup"

// YachGroup struct
type YachGroup struct {
	Channel    string
	Attributes RobotsType
}

// NewYachGroup is used to init robots
func NewYachGroup(robots RobotsType) *YachGroup {
	return &YachGroup{
		Channel:    defaultYachGroupChannelName,
		Attributes: robots,
	}
}

// AddRobot add a robot
func (dg *YachGroup) AddRobot(webhook string, secret string) *YachGroup {
	robot := Robot{
		Webhook: webhook,
		Secret:  secret,
	}
	dg.Attributes = append(dg.Attributes, robot)
	return dg
}

// AddRobots add robots
func (dg *YachGroup) AddRobots(robots RobotsType, replace bool) *YachGroup {
	if replace {
		dg.Attributes = robots
	} else {
		dg.Attributes = append(dg.Attributes, robots...)
	}
	return dg
}

// GetChannel get channel
func (dg *YachGroup) GetChannel() string {
	return dg.Channel
}

// GetAttributes get attributes
func (dg *YachGroup) GetAttributes() []interface{} {
	ret := make([]interface{}, len(dg.Attributes))
	for i := 0; i < len(dg.Attributes); i++ {
		ret[i] = dg.Attributes[i]

	}
	return ret
}
