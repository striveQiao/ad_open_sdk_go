/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

type ApiOpenApi2ClueSmartphoneRecordGetRequestExample struct {
	AdvertiserId int64   `json:"advertiser_id"`
	InstanceIds  []int64 `json:"instance_ids,omitempty"`
	ClueIds      []int64 `json:"clue_ids,omitempty"`
	SiteIds      []int64 `json:"site_ids,omitempty"`
	AdIds        []int64 `json:"ad_ids,omitempty"`
	StartTime    string  `json:"start_time,omitempty"`
	EndTime      string  `json:"end_time,omitempty"`
	Page         int32   `json:"page,omitempty"`
	PageSize     int32   `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/clue/smartphone/record/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ClueSmartphoneRecordGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ClueSmartphoneRecordV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).InstanceIds(request.InstanceIds).ClueIds(request.ClueIds).SiteIds(request.SiteIds).AdIds(request.AdIds).StartTime(request.StartTime).EndTime(request.EndTime).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
