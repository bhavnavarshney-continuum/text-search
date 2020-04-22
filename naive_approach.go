package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/buger/jsonparser"
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
		CurrentEndpointAgentVersion string      `json:"currentEndpointAgentVersion"`
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
	err := json.Unmarshal([]byte(jsonData10k), &model)
	if err != nil {
		return DeviceData{}, err
	}
	return model, nil
}

func improvedUnmarshalling() {
	start := time.Now()
	_, err := jsonparser.ArrayEach([]byte(jsonData30K), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		//fmt.Println(jsonparser.Get(value, "closestNextTaskName"))
		jsonparser.Get(value, "closestNextTaskName")
	}, "items")

	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(val))
	elapsed := time.Since(start)
	fmt.Println("\nNaive Search took %v time", elapsed)
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
	fmt.Printf("%+v", result)
	fmt.Println("\nNaive Search took %v time", elapsed)
}
