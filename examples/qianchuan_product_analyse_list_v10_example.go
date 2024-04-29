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

type ApiOpenApiV10QianchuanProductAnalyseListGetRequestExample struct {
	AdvertiserId int64                                   `json:"advertiser_id"`
	Keyword      string                                  `json:"keyword,omitempty"`
	TimeRange    QianchuanProductAnalyseListV10TimeRange `json:"time_range,omitempty"`
	OrderType    QianchuanProductAnalyseListV10OrderType `json:"order_type,omitempty"`
	OrderField   string                                  `json:"order_field,omitempty"`
	Page         int32                                   `json:"page,omitempty"`
	PageSize     int32                                   `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/product/analyse/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanProductAnalyseListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanProductAnalyseListV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Keyword(request.Keyword).TimeRange(request.TimeRange).OrderType(request.OrderType).OrderField(request.OrderField).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
