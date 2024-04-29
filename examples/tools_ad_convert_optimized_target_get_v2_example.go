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

type ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequestExample struct {
	AdvertiserId       int64                                                `json:"advertiser_id,omitempty"`
	AppSchema          string                                               `json:"app_schema,omitempty"`
	AppType            ToolsAdConvertOptimizedTargetGetV2AppType            `json:"app_type,omitempty"`
	CampaignType       ToolsAdConvertOptimizedTargetGetV2CampaignType       `json:"campaign_type,omitempty"`
	ConvertId          int64                                                `json:"convert_id,omitempty"`
	ConvertName        string                                               `json:"convert_name,omitempty"`
	ConvertType        ToolsAdConvertOptimizedTargetGetV2ConvertType        `json:"convert_type,omitempty"`
	DedicateType       ToolsAdConvertOptimizedTargetGetV2DedicateType       `json:"dedicate_type,omitempty"`
	DeepExternalAction ToolsAdConvertOptimizedTargetGetV2DeepExternalAction `json:"deep_external_action,omitempty"`
	ExternalAction     ToolsAdConvertOptimizedTargetGetV2ExternalAction     `json:"external_action,omitempty"`
	ExternalUrl        string                                               `json:"external_url,omitempty"`
	ItunesUrl          string                                               `json:"itunes_url,omitempty"`
	LandingType        ToolsAdConvertOptimizedTargetGetV2LandingType        `json:"landing_type,omitempty"`
	LaunchTargetType   ToolsAdConvertOptimizedTargetGetV2LaunchTargetType   `json:"launch_target_type,omitempty"`
	MarketingPurpose   ToolsAdConvertOptimizedTargetGetV2MarketingPurpose   `json:"marketing_purpose,omitempty"`
	PackageName        string                                               `json:"package_name,omitempty"`
	Page               int64                                                `json:"page,omitempty"`
	PageSize           int64                                                `json:"page_size,omitempty"`
	PromotionContent   ToolsAdConvertOptimizedTargetGetV2PromotionContent   `json:"promotion_content,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/ad_convert/optimized_target/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAdConvertOptimizedTargetGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AppSchema(request.AppSchema).AppType(request.AppType).CampaignType(request.CampaignType).ConvertId(request.ConvertId).ConvertName(request.ConvertName).ConvertType(request.ConvertType).DedicateType(request.DedicateType).DeepExternalAction(request.DeepExternalAction).ExternalAction(request.ExternalAction).ExternalUrl(request.ExternalUrl).ItunesUrl(request.ItunesUrl).LandingType(request.LandingType).LaunchTargetType(request.LaunchTargetType).MarketingPurpose(request.MarketingPurpose).PackageName(request.PackageName).Page(request.Page).PageSize(request.PageSize).PromotionContent(request.PromotionContent).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
