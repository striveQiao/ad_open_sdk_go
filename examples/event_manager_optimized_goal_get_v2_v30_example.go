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

type ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequestExample struct {
	AdvertiserId       int64                                               `json:"advertiser_id"`
	LandingType        EventManagerOptimizedGoalGetV2V30LandingType        `json:"landing_type"`
	AdType             EventManagerOptimizedGoalGetV2V30AdType             `json:"ad_type"`
	AssetType          EventManagerOptimizedGoalGetV2V30AssetType          `json:"asset_type"`
	AssetId            int64                                               `json:"asset_id,omitempty"`
	PackageName        string                                              `json:"package_name,omitempty"`
	AppType            EventManagerOptimizedGoalGetV2V30AppType            `json:"app_type,omitempty"`
	AppPromotionType   EventManagerOptimizedGoalGetV2V30AppPromotionType   `json:"app_promotion_type,omitempty"`
	MarketingGoal      EventManagerOptimizedGoalGetV2V30MarketingGoal      `json:"marketing_goal,omitempty"`
	QuickAppId         int64                                               `json:"quick_app_id,omitempty"`
	DeliveryMode       EventManagerOptimizedGoalGetV2V30DeliveryMode       `json:"delivery_mode,omitempty"`
	MiniProgramId      string                                              `json:"mini_program_id,omitempty"`
	DpaAdtype          EventManagerOptimizedGoalGetV2V30DpaAdtype          `json:"dpa_adtype,omitempty"`
	MicroPromotionType EventManagerOptimizedGoalGetV2V30MicroPromotionType `json:"micro_promotion_type,omitempty"`
	MicroAppInstanceId int64                                               `json:"micro_app_instance_id,omitempty"`
	DeliveryType       EventManagerOptimizedGoalGetV2V30DeliveryType       `json:"delivery_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/event_manager/optimized_goal/get_v2/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.EventManagerOptimizedGoalGetV2V30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).LandingType(request.LandingType).AdType(request.AdType).AssetType(request.AssetType).AssetId(request.AssetId).PackageName(request.PackageName).AppType(request.AppType).AppPromotionType(request.AppPromotionType).MarketingGoal(request.MarketingGoal).QuickAppId(request.QuickAppId).DeliveryMode(request.DeliveryMode).MiniProgramId(request.MiniProgramId).DpaAdtype(request.DpaAdtype).MicroPromotionType(request.MicroPromotionType).MicroAppInstanceId(request.MicroAppInstanceId).DeliveryType(request.DeliveryType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
