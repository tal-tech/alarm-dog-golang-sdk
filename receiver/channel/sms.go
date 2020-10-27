package channel

const defaultSMSChannelName = "sms"

// SMS struct 手机短信
type SMS struct {
	Channel    string
	Attributes UidsType
}

// NewSMS is used to init robots
func NewSMS(uids UidsType) *SMS {
	return &SMS{
		Channel:    defaultSMSChannelName,
		Attributes: uids,
	}
}

// AddUID add a robot
func (s *SMS) AddUID(uid int) *SMS {
	s.Attributes = append(s.Attributes, uid)
	return s
}

// AddUIDs add robots
func (s *SMS) AddUIDs(uids UidsType, replace bool) *SMS {
	if replace {
		s.Attributes = uids
	} else {
		s.Attributes = append(s.Attributes, uids...)
	}
	return s
}

// GetChannel get channel
func (s *SMS) GetChannel() string {
	return s.Channel
}

// GetAttributes get attributes
func (s *SMS) GetAttributes() []interface{} {
	ret := make([]interface{}, len(s.Attributes))
	for i := 0; i < len(s.Attributes); i++ {
		ret[i] = s.Attributes[i]

	}
	return ret
}
