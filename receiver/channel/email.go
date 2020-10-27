package channel

const defaultEmailChannelName = "email"

// Email struct
type Email struct {
	Channel    string
	Attributes UidsType
}

// NewEmail is used to init robots
func NewEmail(uids UidsType) *Email {
	return &Email{
		Channel:    defaultEmailChannelName,
		Attributes: uids,
	}
}

// AddUID add a robot
func (e *Email) AddUID(uid int) *Email {
	e.Attributes = append(e.Attributes, uid)
	return e
}

// AddUIDs add robots
func (e *Email) AddUIDs(uids UidsType, replace bool) *Email {
	if replace {
		e.Attributes = uids
	} else {
		e.Attributes = append(e.Attributes, uids...)
	}
	return e
}

// GetChannel get channel
func (e *Email) GetChannel() string {
	return e.Channel
}

// GetAttributes get attributes
func (e *Email) GetAttributes() []interface{} {
	ret := make([]interface{}, len(e.Attributes))
	for i := 0; i < len(e.Attributes); i++ {
		ret[i] = e.Attributes[i]

	}
	return ret
}
