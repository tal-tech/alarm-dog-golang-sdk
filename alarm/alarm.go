package alarm

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/tal-tech/alarm-dog-golang-sdk/receiver"
)

// BASEURI is default URI
const BASEURI = "https://alarm-dog-service.domain.com"

// Leveltype is alarm level

const (
	// NOTICE 通知
	NOTICE int = iota // 0
	// WARNING 警告
	WARNING // 1
	// ERROR 错误
	ERROR // 2
	// EMERGENCY 紧急
	EMERGENCY // 3
)

// Alarm struct
type Alarm struct {
	token   string
	taskID  int
	baseURI string
}

// ReportBody struct
type ReportBody struct {
	TaskID     int                `json:"taskid,omitempty"`
	Timestamp  int64              `json:"timestamp,omitempty"`
	Sign       string             `json:"sign,omitempty"`
	Ctn        map[string]string  `json:"ctn,omitempty"`
	NoticeTime int64              `json:"notice_time,omitempty"`
	Level      int                `json:"level,omitempty"`
	Receiver   *receiver.Receiver `json:"receiver,omitempty"`
}

// reportBody_ struct
type realReportBody struct {
	TaskID     int                    `json:"taskid,omitempty"`
	Timestamp  int64                  `json:"timestamp,omitempty"`
	Sign       string                 `json:"sign,omitempty"`
	Ctn        map[string]string      `json:"ctn,omitempty"`
	NoticeTime int64                  `json:"notice_time,omitempty"`
	Level      int                    `json:"level,omitempty"`
	Receiver   receiver.ToArrayStruct `json:"receiver,omitempty"`
}

// ResponseBody struct
type ResponseBody struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data,omitempty"`
}

// New used to init Alarm
func New() *Alarm {
	return &Alarm{
		token:   "token",
		taskID:  1,
		baseURI: BASEURI,
	}
}

// Report used to report content to dog service
func (a *Alarm) Report(reportBody ReportBody) (response *http.Response) {
	timestamp, sign := a.genSign()

	reportBody.TaskID = a.taskID
	if reportBody.Timestamp == 0 {
		reportBody.Timestamp = timestamp
	}
	reportBody.Sign = sign

	var realReportBody interface{}

	realReportBody = reportBody

	// 必须在赋值完所有的reportBody之后进行转换
	if reportBody.Receiver != nil {
		realReportBody = convert(&reportBody)
	}

	reportBodyJSON, err := json.Marshal(realReportBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(string(reportBodyJSON))
	response, err = http.Post(a.baseURI+"/alarm/report", "application/json", strings.NewReader(string(reportBodyJSON)))
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// ResolveResponse used to parse http.Response and get ResponseBody struct
func (a *Alarm) ResolveResponse(response *http.Response) (responseBody *ResponseBody, err error) {
	responseBody = &ResponseBody{}
	statusCode := response.StatusCode
	if statusCode != 200 {
		responseBody.Code = statusCode
		responseBody.Msg = "signature invalid"
		err = fmt.Errorf(fmt.Sprintf("send alarm failed: response status code not 200 but got %d", statusCode))
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	if err = json.Unmarshal([]byte(string(body)), responseBody); err != nil {
		return
	}
	if responseBody.Code != 0 {
		err = fmt.Errorf(fmt.Sprintf("send alarm failed: response error: %s", responseBody.Msg))
		return
	}
	return
}

// SetTaskID used to set taskID
func (a *Alarm) SetTaskID(taskID int) *Alarm {
	a.taskID = taskID
	return a
}

// GetTaskID used to get taskID
func (a *Alarm) GetTaskID() int {
	return a.taskID
}

// SetToken used to set token
func (a *Alarm) SetToken(token string) *Alarm {
	a.token = token
	return a
}

// GetToken used to get token
func (a *Alarm) GetToken() string {
	return a.token
}

// SetBaseURI used to set base URI. If pass a error URI, it will use default URI
func (a *Alarm) SetBaseURI(baseURI string) *Alarm {
	a.baseURI = baseURI
	return a
}

// GetBaseURI used to get base URI
func (a *Alarm) GetBaseURI() string {
	return a.baseURI
}

func (a *Alarm) genSign() (timestamp int64, sign string) {
	timestamp = time.Now().Unix()
	input := fmt.Sprintf("%d&%d%s", a.taskID, timestamp, a.token)

	digest := md5.New()
	_, err := digest.Write([]byte(input))
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	// 这里必须得使用md5的字符串格式
	sign = hex.EncodeToString(digest.Sum(nil))
	return
}

func convert(body *ReportBody) interface{} {
	return &realReportBody{
		TaskID:     body.TaskID,
		Timestamp:  body.Timestamp,
		Sign:       body.Sign,
		Ctn:        body.Ctn,
		NoticeTime: body.NoticeTime,
		Level:      body.Level,
		Receiver:   *body.Receiver.ToArray(),
	}
}
