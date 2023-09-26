/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
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

type ApiOpenApi2ToolsClueFormGetGetRequestExample struct {
	AdvertiserId int64                      `json:"advertiser_id"`
	Page         int32                      `json:"page,omitempty"`
	PageSize     int32                      `json:"page_size,omitempty"`
	StartTime    string                     `json:"start_time,omitempty"`
	EndTime      string                     `json:"end_time,omitempty"`
	InstanceId   int64                      `json:"instance_id,omitempty"`
	Name         string                     `json:"name,omitempty"`
	IsDel        int64                      `json:"is_del,omitempty"`
	FormType     ToolsClueFormGetV2FormType `json:"form_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/clue/form/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsClueFormGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsClueFormGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Page(request.Page).PageSize(request.PageSize).StartTime(request.StartTime).EndTime(request.EndTime).InstanceId(request.InstanceId).Name(request.Name).IsDel(request.IsDel).FormType(request.FormType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
