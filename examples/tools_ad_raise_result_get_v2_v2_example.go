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

type ApiOpenApi2ToolsAdRaiseResultGetV2GetRequestExample struct {
	AdId           int64                                  `json:"ad_id,omitempty"`
	AdRaiseVersion int64                                  `json:"ad_raise_version,omitempty"`
	AdvertiserId   int64                                  `json:"advertiser_id,omitempty"`
	EndTime        *string                                `json:"end_time,omitempty"`
	OrderField     string                                 `json:"order_field,omitempty"`
	OrderType      ToolsAdRaiseResultGetV2V2OrderType     `json:"order_type,omitempty"`
	Page           int64                                  `json:"page,omitempty"`
	PageSize       int64                                  `json:"page_size,omitempty"`
	StartTime      *string                                `json:"start_time,omitempty"`
	TimeDimension  ToolsAdRaiseResultGetV2V2TimeDimension `json:"time_dimension,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/ad_raise_result/get_v2/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAdRaiseResultGetV2GetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAdRaiseResultGetV2V2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdId(request.AdId).AdRaiseVersion(request.AdRaiseVersion).AdvertiserId(request.AdvertiserId).EndTime(request.EndTime).OrderField(request.OrderField).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).StartTime(request.StartTime).TimeDimension(request.TimeDimension).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
