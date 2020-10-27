package channel

const defaultYachWorkerChannelName = "yachworker"

// YachWorker struct
type YachWorker struct {
	Channel    string
	Attributes UidsType
}

// NewYachWorker is used to init robots
func NewYachWorker(uids UidsType) *YachWorker {
	return &YachWorker{
		Channel:    defaultYachWorkerChannelName,
		Attributes: uids,
	}
}

// AddUID add a robot
func (yw *YachWorker) AddUID(uid int) *YachWorker {
	yw.Attributes = append(yw.Attributes, uid)
	return yw
}

// AddUIDs add robots
func (yw *YachWorker) AddUIDs(uids UidsType, replace bool) *YachWorker {
	if replace {
		yw.Attributes = uids
	} else {
		yw.Attributes = append(yw.Attributes, uids...)
	}
	return yw
}

// GetChannel get channel
func (yw *YachWorker) GetChannel() string {
	return yw.Channel
}

// GetAttributes get attributes
func (yw *YachWorker) GetAttributes() []interface{} {
	ret := make([]interface{}, len(yw.Attributes))
	for i := 0; i < len(yw.Attributes); i++ {
		ret[i] = yw.Attributes[i]

	}
	return ret
}
