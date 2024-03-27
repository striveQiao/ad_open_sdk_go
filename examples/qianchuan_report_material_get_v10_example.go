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

type ApiOpenApiV10QianchuanReportMaterialGetGetRequestExample struct {
	StartDate    string                                 `json:"start_date"`
	EndDate      string                                 `json:"end_date"`
	AdvertiserId int64                                  `json:"advertiser_id"`
	Fields       []string                               `json:"fields"`
	OrderType    QianchuanReportMaterialGetV10OrderType `json:"order_type,omitempty"`
	OrderField   string                                 `json:"order_field,omitempty"`
	Filtering    QianchuanReportMaterialGetV10Filtering `json:"filtering,omitempty"`
	Page         int32                                  `json:"page,omitempty"`
	PageSize     int32                                  `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/report/material/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanReportMaterialGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanReportMaterialGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		StartDate(request.StartDate).EndDate(request.EndDate).AdvertiserId(request.AdvertiserId).Fields(request.Fields).OrderType(request.OrderType).OrderField(request.OrderField).Filtering(request.Filtering).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
