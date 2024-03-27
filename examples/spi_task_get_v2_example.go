/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
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

type ApiOpenApi2SpiTaskGetGetRequestExample struct {
	AppId        int64              `json:"app_id,omitempty"`
	Count        int64              `json:"count,omitempty"`
	Cursor       int64              `json:"cursor,omitempty"`
	EndDate      string             `json:"end_date,omitempty"`
	Page         int64              `json:"page,omitempty"`
	PageSize     int64              `json:"page_size,omitempty"`
	ServiceLabel string             `json:"service_label,omitempty"`
	StartDate    string             `json:"start_date,omitempty"`
	Status       SpiTaskGetV2Status `json:"status,omitempty"`
	SubscribeId  int64              `json:"subscribe_id,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/spi_task/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2SpiTaskGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.SpiTaskGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AppId(request.AppId).Count(request.Count).Cursor(request.Cursor).EndDate(request.EndDate).Page(request.Page).PageSize(request.PageSize).ServiceLabel(request.ServiceLabel).StartDate(request.StartDate).Status(request.Status).SubscribeId(request.SubscribeId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
