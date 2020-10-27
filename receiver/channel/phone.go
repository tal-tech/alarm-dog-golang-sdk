package channel

const defaultPhoneChannelName = "phone"

// Phone struct
type Phone struct {
	Channel    string
	Attributes UidsType
}

// NewPhone is used to init robots
func NewPhone(uids UidsType) *Phone {
	return &Phone{
		Channel:    defaultPhoneChannelName,
		Attributes: uids,
	}
}

// AddUID add a robot
func (p *Phone) AddUID(uid int) *Phone {
	p.Attributes = append(p.Attributes, uid)
	return p
}

// AddUIDs add robots
func (p *Phone) AddUIDs(uids UidsType, replace bool) *Phone {
	if replace {
		p.Attributes = uids
	} else {
		p.Attributes = append(p.Attributes, uids...)
	}
	return p
}

// GetChannel get channel
func (p *Phone) GetChannel() string {
	return p.Channel
}

// GetAttributes get attributes
func (p *Phone) GetAttributes() []interface{} {
	ret := make([]interface{}, len(p.Attributes))
	for i := 0; i < len(p.Attributes); i++ {
		ret[i] = p.Attributes[i]

	}
	return ret
}
