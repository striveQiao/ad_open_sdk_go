/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
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

type ApiOpenApi2ToolsUnionFlowPackageReportGetRequestExample struct {
	AdvertiserId int64                                  `json:"advertiser_id,omitempty"`
	Filter       ToolsUnionFlowPackageReportV2Filter    `json:"filter,omitempty"`
	OrderField   string                                 `json:"order_field,omitempty"`
	OrderType    ToolsUnionFlowPackageReportV2OrderType `json:"order_type,omitempty"`
	Page         int64                                  `json:"page,omitempty"`
	PageSize     int64                                  `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/union/flow_package/report/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsUnionFlowPackageReportGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsUnionFlowPackageReportV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Filter(request.Filter).OrderField(request.OrderField).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
