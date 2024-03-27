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

type ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequestExample struct {
	AdId         int64                                            `json:"ad_id,omitempty"`
	AdvertiserId int64                                            `json:"advertiser_id,omitempty"`
	ContentType  ToolsPromotionCardRecommendTitleGetV2ContentType `json:"content_type,omitempty"`
	ExternalUrl  string                                           `json:"external_url,omitempty"`
	IndustryId   int64                                            `json:"industry_id,omitempty"`
	TextType     ToolsPromotionCardRecommendTitleGetV2TextType    `json:"text_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/promotion_card/recommend_title/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsPromotionCardRecommendTitleGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdId(request.AdId).AdvertiserId(request.AdvertiserId).ContentType(request.ContentType).ExternalUrl(request.ExternalUrl).IndustryId(request.IndustryId).TextType(request.TextType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
