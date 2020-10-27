package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhook_SetWebhook(t *testing.T) {
	webhooks := WebhooksType{}
	wh := NewWebhook(webhooks)
	wh.SetWebhook("http://webhook.com")
	assert.Len(t, wh.Attributes, 1)
}

func TestWebhook_SetWebhookWithErrorURI(t *testing.T) {
	webhooks := WebhooksType{}
	wh := NewWebhook(webhooks)
	wh.SetWebhook("htt://webhook.com")
	assert.Len(t, wh.Attributes, 0)
}
