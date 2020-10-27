package main

import "github.com/tal-tech/alarm-dog-golang-sdk/alarm"
import "fmt"

/**
 * This is a test file, please read the README.md file for usage.
 */
func main() {
    a := alarm.New()
    a.SetTaskID(637).SetToken("349f44f37cdde77b2c393c7761e667f382387b0d").SetBaseURI("https://alarm-dog-service.domain.com")

    var content = make(map[string]string)
    content["errno"] = "10086"
    content["error"] = "您的余额已不足"

    reportBody := alarm.ReportBody{
        Ctn: content,
    }

	response := a.Report(reportBody)
	
    fmt.Println(response)
}
