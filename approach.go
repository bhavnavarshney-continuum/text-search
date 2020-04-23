package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/buger/jsonparser"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

type DeviceData struct {
	Items []struct {
		Dated                       int         `json:"dated"`
		AvStatus                    bool        `json:"avStatus"`
		ClientID                    int         `json:"clientId"`
		ClosestNextTaskName         string      `json:"closestNextTaskName"`
		ClosestNextTaskRunDate      int         `json:"closestNextTaskRunDate"`
		ClosestPreviousTaskName     string      `json:"closestPreviousTaskName"`
		ClosestPreviousTaskRunDate  int         `json:"closestPreviousTaskRunDate"`
		ClosestPreviousTaskStatus   bool        `json:"closestPreviousTaskStatus"`
		CriticalTicketsCnt          int         `json:"criticalTicketsCnt"`
		CurrentEndpointAgentVersion float64     `json:"currentEndpointAgentVersion"`
		EndPointID                  int         `json:"endPointId"`
		FriendlyName                string      `json:"friendlyName"`
		HbAgentDateTimeUTC          int         `json:"hbAgentDateTimeUTC"`
		HbAvailability              bool        `json:"hbAvailability"`
		LastRestartDate             int         `json:"lastRestartDate"`
		LatestEnabledAgentVersion   interface{} `json:"latestEnabledAgentVersion"`
		LmiStatus                   int         `json:"lmiStatus"`
		MachineName                 string      `json:"machineName"`
		OsProduct                   string      `json:"osProduct"`
		PartnerID                   int         `json:"partnerId"`
		PatchingStatus              bool        `json:"patchingStatus"`
		PchStatus                   bool        `json:"pchStatus"`
		RegID                       int         `json:"regId"`
		ResourceType                string      `json:"resourceType"`
		SiteID                      int         `json:"siteId"`
		SiteName                    string      `json:"siteName"`
		SsLogonTime                 int         `json:"ssLogonTime"`
		SsStatus                    bool        `json:"ssStatus"`
		SsUserName                  string      `json:"ssUserName"`
		TaskCount                   int         `json:"taskCount"`
		TicketsCnt                  int         `json:"ticketsCnt"`
		TimeZone                    int         `json:"timeZone"`
		TimeZoneDescription         int         `json:"timeZoneDescription"`
		WrStatus                    bool        `json:"wrStatus"`
	} `json:"items"`
}

func unmarshalJSON() (DeviceData, error) {
	model := DeviceData{}
	err := json.Unmarshal([]byte(jsonData10K), &model)
	if err != nil {
		return DeviceData{}, err
	}
	return model, nil
}

func goJSONQ(hint string) {
	start := time.Now()
	q1 := gojsonq.New().FromString(jsonData60K).From("items").Where("siteName", "=", hint)
	q2 := gojsonq.New().FromString(jsonData60K).From("items").Where("osProduct", "=", hint)
	q3 := gojsonq.New().FromString(jsonData60K).From("items").Where("friendlyName", "=", hint)
	q4 := gojsonq.New().FromString(jsonData60K).From("items").Where("machineName", "=", hint)
	q5 := gojsonq.New().FromString(jsonData60K).From("items").Where("resourceType", "=", hint)

	elapsed := time.Since(start)
	fmt.Printf("\nGoJSONQ took %v time", elapsed)

	fmt.Printf("\nSiteName:%v, OsProduct:%v, FriendlyName:%v, MachineName:%v,ResourceType:%v ", q1.Count(), q2.Count(), q3.Count(), q4.Count(), q5.Count())
}

// func jsonParserLibrary() {
// 	start := time.Now()
// 	_, err := jsonparser.ArrayEach([]byte(jsonData30K), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
// 		//fmt.Println(jsonparser.Get(value, "closestNextTaskName"))
// 		jsonparser.Get(value, "closestNextTaskName")
// 	}, "items")

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	//fmt.Println(string(val))
// 	elapsed := time.Since(start)
// 	fmt.Println("\nJsonParser took %v time", elapsed)
// }

func jsonParserLibrary(hint string) {
	start := time.Now()
	c := 0
	_, err := jsonparser.ArrayEach([]byte(jsonData1K), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		//fmt.Println(jsonparser.Get(value, "closestNextTaskName"))
		osProduct, _, _, _ := jsonparser.Get(value, "osProduct")
		siteName, _, _, _ := jsonparser.Get(value, "siteName")
		friendlyName, _, _, _ := jsonparser.Get(value, "friendlyName")
		machineName, _, _, _ := jsonparser.Get(value, "machineName")
		resourceType, _, _, _ := jsonparser.Get(value, "resourceType")

		if string(osProduct) == hint || string(siteName) == hint || string(friendlyName) == hint || string(machineName) == hint || string(resourceType) == hint {
			c++
		}
	}, "items")

	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(val))
	elapsed := time.Since(start)
	fmt.Printf("\n Match found in %v records", c)
	fmt.Printf("\nJsonParser took %v time", elapsed)
}

func naiveSearch(hint string) {
	start := time.Now()
	goData, err := unmarshalJSON()
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	result := make(map[int]interface{})
	j := 0

	for i := 0; i < len(goData.Items); i++ {
		if goData.Items[i].SiteName == hint || goData.Items[i].FriendlyName == hint || goData.Items[i].MachineName == hint || goData.Items[i].OsProduct == hint || goData.Items[i].ResourceType == hint || goData.Items[i].ClosestNextTaskName == hint {
			result[j] = goData.Items[i]
			j++
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Match Found in %v", len(result))
	fmt.Println("\nNaive Search took %v time", elapsed)
}
