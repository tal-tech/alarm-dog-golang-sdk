package channel

// Channel interface
type Channel interface {
	GetChannel() string
	GetAttributes() []interface{}
}

// UidsType type 员工工号
type UidsType []int

// RobotsType type
type RobotsType []Robot

// Robot struct
type Robot struct {
	Webhook string `json:"webhook,omitempty"`
	Secret  string `json:"secret,omitempty"`
}
