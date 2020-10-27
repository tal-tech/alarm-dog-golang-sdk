package channel

const defaultDingWorkerChannelName = "dingworker"

// DingWorker struct
type DingWorker struct {
	Channel    string
	Attributes UidsType
}

// NewDingWorker is used to init robots
func NewDingWorker(uids UidsType) *DingWorker {
	return &DingWorker{
		Channel:    defaultDingWorkerChannelName,
		Attributes: uids,
	}
}

// AddUID add a robot
func (dw *DingWorker) AddUID(uid int) *DingWorker {
	dw.Attributes = append(dw.Attributes, uid)
	return dw
}

// AddUIDs add robots
func (dw *DingWorker) AddUIDs(uids UidsType, replace bool) *DingWorker {
	if replace {
		dw.Attributes = uids
	} else {
		dw.Attributes = append(dw.Attributes, uids...)
	}
	return dw
}

// GetChannel get channel
func (dw *DingWorker) GetChannel() string {
	return dw.Channel
}

// GetAttributes get attributes
func (dw *DingWorker) GetAttributes() []interface{} {
	ret := make([]interface{}, len(dw.Attributes))
	for i := 0; i < len(dw.Attributes); i++ {
		ret[i] = dw.Attributes[i]

	}
	return ret
}
