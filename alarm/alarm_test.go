package alarm

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/tal-tech/alarm-dog-golang-sdk/receiver"
	"github.com/tal-tech/alarm-dog-golang-sdk/receiver/channel"

	"github.com/stretchr/testify/assert"
)

const (
	validwebhhok = "5fc8fb9191473d10ab16fbf40702d9b7c53dcb843d854a73b7045b3d491b3dc4"
	validsecret  = "SEC02421cebaba9ee05ac0165447403a2d333a846fb2e147ce2a91ec5ae8b1b8e9c"
)

const (
	validTaskID = 637
	validToken  = "349f44f37cdde77b2c393c7761e667f382387b0d"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestAlarm_Report(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	var content = make(map[string]string)
	content["errno"] = "10086"
	content["error"] = "您的余额已不足"

	reportBody := ReportBody{
		Ctn: content,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(0), responseMap["code"])
	assert.Equal(t, "success", responseMap["msg"])
}

func TestAlarm_ReportWithLevel(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	var content = make(map[string]string)
	content["errno"] = "10086"
	content["error"] = "您的余额已不足"

	reportBody := ReportBody{
		Ctn:   content,
		Level: NOTICE,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(0), responseMap["code"])
	assert.Equal(t, "success", responseMap["msg"])
}

func TestAlarm_ReportWithErrorAlarmGroupReceiver(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	channels := make([]channel.Channel, 2)
	dg := channel.NewDingGroup(channel.RobotsType{
		channel.Robot{
			Webhook: "webhook1",
			Secret:  "secret1",
		},
		channel.Robot{
			Webhook: "webhook2",
			Secret:  "secret2",
		},
	})
	wg := channel.NewDingWorker(channel.UidsType{
		1,
		2,
	})
	channels[0] = dg
	channels[1] = wg
	r := receiver.NewReceiver([]int{1, 2, 3, 4}, channels)

	var content = make(map[string]string)
	content["errno"] = "10086"
	content["error"] = "您的余额已不足"

	reportBody := ReportBody{
		Ctn:      content,
		Level:    NOTICE,
		Receiver: r,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(4000), responseMap["code"])
	assert.Equal(t, "receiver.alarmgroup is invalid(4000)", responseMap["msg"])
}

func TestAlarm_ReportWithDingWorkerReceiver(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	channels := make([]channel.Channel, 1)
	wg := channel.NewDingWorker(channel.UidsType{
		167453,
	})
	channels[0] = wg
	r := receiver.NewReceiver([]int{}, channels)

	var content = make(map[string]string)
	content["errno"] = "10086"
	content["error"] = "您的余额已不足"

	reportBody := ReportBody{
		Ctn:      content,
		Level:    NOTICE,
		Receiver: r,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(0), responseMap["code"])
	assert.Equal(t, "success", responseMap["msg"])
}

func TestAlarm_ReportWithDingGroupReceiver(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	channels := make([]channel.Channel, 1)
	dg := channel.NewDingGroup(channel.RobotsType{
		channel.Robot{
			Webhook: validwebhhok,
			Secret:  validsecret,
		},
	})
	channels[0] = dg
	r := receiver.NewReceiver([]int{}, channels)

	var content = make(map[string]string)
	content["errno"] = "dinggroup"
	content["error"] = "golang dinggroup test"

	reportBody := ReportBody{
		Ctn:      content,
		Level:    NOTICE,
		Receiver: r,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(0), responseMap["code"])
	assert.Equal(t, "success", responseMap["msg"])
}

func TestAlarm_ReportWithErrorDingGroupReceiver(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	channels := make([]channel.Channel, 1)
	wg := channel.NewDingGroup(channel.RobotsType{
		channel.Robot{
			Webhook: "http://11111",
			Secret:  "xxxxxx",
		},
	})
	channels[0] = wg
	r := receiver.NewReceiver([]int{}, channels)

	var content = make(map[string]string)
	content["errno"] = "dinggroup"
	content["error"] = "dinggroup"

	reportBody := ReportBody{
		Ctn:      content,
		Level:    NOTICE,
		Receiver: r,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(0), responseMap["code"])
	assert.Equal(t, "success", responseMap["msg"])
}

func TestAlarm_ReportWithYachGroupReceiver(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	channels := make([]channel.Channel, 1)
	wg := channel.NewYachGroup(channel.RobotsType{
		channel.Robot{
			Webhook: "d2NpMVFJSHlYM3dzalhnOHVPQjdWOFpJSXBkMkovanhOU05ZamVmeWl5UmdUbEkvc2k3Vnp5YWp6V2EwSU9ITg",
			Secret:  "SECf74f72fd30a923ffb534de04fd1a6dff",
		},
	})
	channels[0] = wg
	r := receiver.NewReceiver([]int{}, channels)

	var content = make(map[string]string)
	content["errno"] = "dinggroup"
	content["error"] = "golang yachgroup test"

	reportBody := ReportBody{
		Ctn:      content,
		Level:    NOTICE,
		Receiver: r,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(0), responseMap["code"])
	assert.Equal(t, "success", responseMap["msg"])
}

func TestAlarm_ReportWithErrorYachGroupReceiver(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	channels := make([]channel.Channel, 1)
	wg := channel.NewYachGroup(channel.RobotsType{
		channel.Robot{
			Webhook: "http://11111",
			Secret:  "xxxxxx",
		},
	})
	channels[0] = wg
	r := receiver.NewReceiver([]int{}, channels)

	var content = make(map[string]string)
	content["errno"] = "yachgroup"
	content["error"] = "yachgroup"

	reportBody := ReportBody{
		Ctn:      content,
		Level:    NOTICE,
		Receiver: r,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(0), responseMap["code"])
	assert.Equal(t, "success", responseMap["msg"])
}

func TestAlarm_ReportWithEmailReceiver(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	channels := make([]channel.Channel, 1)
	wg := channel.NewEmail(channel.UidsType{
		167453,
	})
	channels[0] = wg
	r := receiver.NewReceiver([]int{}, channels)

	var content = make(map[string]string)
	content["errno"] = "10086"
	content["error"] = "您的余额已不足"

	reportBody := ReportBody{
		Ctn:      content,
		Level:    NOTICE,
		Receiver: r,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(0), responseMap["code"])
	assert.Equal(t, "success", responseMap["msg"])
}

func TestAlarm_ReportWithPhoneReceiver(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	channels := make([]channel.Channel, 1)
	wg := channel.NewPhone(channel.UidsType{
		167453,
	})
	channels[0] = wg
	r := receiver.NewReceiver([]int{}, channels)

	var content = make(map[string]string)
	content["errno"] = "167453"
	content["error"] = "手机测试"

	reportBody := ReportBody{
		Ctn:      content,
		Level:    NOTICE,
		Receiver: r,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(0), responseMap["code"])
	assert.Equal(t, "success", responseMap["msg"])
}

func TestAlarm_ReportWithSMSReceiver(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	channels := make([]channel.Channel, 1)
	wg := channel.NewSMS(channel.UidsType{
		167453,
	})
	channels[0] = wg
	r := receiver.NewReceiver([]int{}, channels)

	var content = make(map[string]string)
	content["errno"] = "sms test"
	content["error"] = "sms test"

	reportBody := ReportBody{
		Ctn:      content,
		Level:    NOTICE,
		Receiver: r,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(0), responseMap["code"])
	assert.Equal(t, "success", responseMap["msg"])
}

func TestAlarm_ReportWithErrorWebhookReceiver(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	channels := make([]channel.Channel, 1)
	wg := channel.NewWebhook(channel.WebhooksType{
		"http://127.0.0.1",
	})
	channels[0] = wg
	r := receiver.NewReceiver([]int{}, channels)

	var content = make(map[string]string)
	content["errno"] = "error webhook"
	content["error"] = "error webhook"

	reportBody := ReportBody{
		Ctn:      content,
		Level:    NOTICE,
		Receiver: r,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(4000), responseMap["code"])
	assert.Equal(t, "receiver.channels.webhook.url must be a valid url(4000)", responseMap["msg"])
}

func TestAlarm_ReportWithYachWorkerReceiver(t *testing.T) {
	a := New()
	a.SetTaskID(validTaskID).SetToken(validToken)

	channels := make([]channel.Channel, 1)
	wg := channel.NewYachWorker(channel.UidsType{
		167453,
	})
	channels[0] = wg
	r := receiver.NewReceiver([]int{}, channels)

	var content = make(map[string]string)
	content["errno"] = "10086"
	content["error"] = "您的余额已不足"

	reportBody := ReportBody{
		Ctn:      content,
		Level:    NOTICE,
		Receiver: r,
	}

	response := a.Report(reportBody)
	assert.NotNil(t, response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode)

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &responseMap); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(0), responseMap["code"])
	assert.Equal(t, "success", responseMap["msg"])
}

func TestAlarm_ResolveResponse(t *testing.T) {
	a := New()
	{ // error
		a.SetTaskID(1).SetToken(validToken)

		var content = make(map[string]string)
		content["errno"] = "10086"
		content["error"] = "您的余额已不足"

		reportBody := ReportBody{
			Ctn: content,
		}

		response := a.Report(reportBody)
		assert.NotNil(t, response)

		defer response.Body.Close()
		responseBody, err := a.ResolveResponse(response)
		assert.NotNil(t, err)
		assert.NotNil(t, responseBody)
		assert.Equal(t, 401, responseBody.Code)
		assert.Equal(t, "signature invalid", responseBody.Msg)
	}
	{ // success
		a.SetTaskID(validTaskID).SetToken(validToken)

		var content = make(map[string]string)
		content["errno"] = "10086"
		content["error"] = "您的余额已不足"

		reportBody := ReportBody{
			Ctn: content,
		}

		response := a.Report(reportBody)
		assert.NotNil(t, response)

		defer response.Body.Close()
		responseBody, err := a.ResolveResponse(response)
		assert.Nil(t, err)
		assert.NotNil(t, responseBody)
		assert.Equal(t, 0, responseBody.Code)
		assert.Equal(t, "success", responseBody.Msg)
	}
}

func TestAlarm_SetBaseURI(t *testing.T) {
	a := New()
	{
		a.SetBaseURI("http://test")
		assert.Equal(t, BASEURI, a.GetBaseURI())
	}
	{
		a.SetBaseURI("http://test.com")
		assert.Equal(t, "http://test.com", a.GetBaseURI())
	}
	{
		a.SetBaseURI("https://test.com")
		assert.Equal(t, "https://test.com", a.GetBaseURI())
	}
}
