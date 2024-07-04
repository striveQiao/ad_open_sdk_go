/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
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

type ApiOpenApiV30EventManagerDeepBidTypeGetGetRequestExample struct {
	AdvertiserId       int64                                           `json:"advertiser_id"`
	ExternalAction     EventManagerDeepBidTypeGetV30ExternalAction     `json:"external_action"`
	AssetId            int64                                           `json:"asset_id,omitempty"`
	DeepExternalAction EventManagerDeepBidTypeGetV30DeepExternalAction `json:"deep_external_action,omitempty"`
	ConvertId          int64                                           `json:"convert_id,omitempty"`
	DeliveryMode       EventManagerDeepBidTypeGetV30DeliveryMode       `json:"delivery_mode,omitempty"`
	LandingType        EventManagerDeepBidTypeGetV30LandingType        `json:"landing_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/event_manager/deep_bid_type/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30EventManagerDeepBidTypeGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.EventManagerDeepBidTypeGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).ExternalAction(request.ExternalAction).AssetId(request.AssetId).DeepExternalAction(request.DeepExternalAction).ConvertId(request.ConvertId).DeliveryMode(request.DeliveryMode).LandingType(request.LandingType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
