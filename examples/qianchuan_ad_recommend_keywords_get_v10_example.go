/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
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

type ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequestExample struct {
	AdvertiserId int64                                        `json:"advertiser_id"`
	Filtering    QianchuanAdRecommendKeywordsGetV10Filtering  `json:"filtering"`
	OrderField   QianchuanAdRecommendKeywordsGetV10OrderField `json:"order_field,omitempty"`
	OrderType    QianchuanAdRecommendKeywordsGetV10OrderType  `json:"order_type,omitempty"`
	CacheId      string                                       `json:"cache_id,omitempty"`
	Page         int32                                        `json:"page,omitempty"`
	PageSize     int32                                        `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/ad/recommend_keywords/get Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanAdRecommendKeywordsGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Filtering(request.Filtering).OrderField(request.OrderField).OrderType(request.OrderType).CacheId(request.CacheId).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
