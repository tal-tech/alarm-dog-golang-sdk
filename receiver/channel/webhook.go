package channel

import "regexp"

// WebhooksType type
type WebhooksType []string

const defaultWebhookChannelName = "webhook"

// Webhook struct
type Webhook struct {
	Channel    string
	Attributes WebhooksType
}

// NewWebhook is used to init robots
func NewWebhook(webhooks WebhooksType) *Webhook {
	return &Webhook{
		Channel:    defaultWebhookChannelName,
		Attributes: webhooks,
	}
}

// SetWebhook used to set webhook
func (s *Webhook) SetWebhook(webhook string) *Webhook {
	re := regexp.MustCompile(`^((http|https)://)?[a-zA-Z0-9-]+(\.[a-zA-z0-9-]+)+/?$`)
	isMatch := re.MatchString(webhook)
	if isMatch {
		s.Attributes = append(s.Attributes, webhook)
	}
	return s
}

// GetChannel get channel
func (s *Webhook) GetChannel() string {
	return s.Channel
}

// GetAttributes get attributes
func (s *Webhook) GetAttributes() []interface{} {
	ret := make([]interface{}, len(s.Attributes))
	for i := 0; i < len(s.Attributes); i++ {
		ret[i] = s.Attributes[i]

	}
	return ret
}
