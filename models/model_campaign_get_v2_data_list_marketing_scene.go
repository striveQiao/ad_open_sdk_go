/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CampaignGetV2DataListMarketingScene
type CampaignGetV2DataListMarketingScene string

// List of campaign_get_v2_data_list_marketing_scene
const (
	EDUCATION_CampaignGetV2DataListMarketingScene         CampaignGetV2DataListMarketingScene = "EDUCATION"
	NOVEL_CampaignGetV2DataListMarketingScene             CampaignGetV2DataListMarketingScene = "NOVEL"
	GAME_PROMOTION_CampaignGetV2DataListMarketingScene    CampaignGetV2DataListMarketingScene = "GAME_PROMOTION"
	PROMOTION_PURPOSE_CampaignGetV2DataListMarketingScene CampaignGetV2DataListMarketingScene = "PROMOTION_PURPOSE"
	CAR_CampaignGetV2DataListMarketingScene               CampaignGetV2DataListMarketingScene = "CAR"
	VIDEO_INFO_CampaignGetV2DataListMarketingScene        CampaignGetV2DataListMarketingScene = "VIDEO_INFO"
	MERCHANTS_CampaignGetV2DataListMarketingScene         CampaignGetV2DataListMarketingScene = "MERCHANTS"
	SOCIAL_CampaignGetV2DataListMarketingScene            CampaignGetV2DataListMarketingScene = "SOCIAL"
	GAME_SUBSCRIBE_CampaignGetV2DataListMarketingScene    CampaignGetV2DataListMarketingScene = "GAME_SUBSCRIBE"
	ECOMMERCE_CampaignGetV2DataListMarketingScene         CampaignGetV2DataListMarketingScene = "ECOMMERCE"
)

// All allowed values of CampaignGetV2DataListMarketingScene enum
var AllowedCampaignGetV2DataListMarketingSceneEnumValues = []CampaignGetV2DataListMarketingScene{
	"EDUCATION",
	"NOVEL",
	"GAME_PROMOTION",
	"PROMOTION_PURPOSE",
	"CAR",
	"VIDEO_INFO",
	"MERCHANTS",
	"SOCIAL",
	"GAME_SUBSCRIBE",
	"ECOMMERCE",
}

// NewCampaignGetV2DataListMarketingSceneFromValue returns a pointer to a valid CampaignGetV2DataListMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2DataListMarketingSceneFromValue(v string) (*CampaignGetV2DataListMarketingScene, error) {
	ev := CampaignGetV2DataListMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2DataListMarketingScene: valid values are %v", v, AllowedCampaignGetV2DataListMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2DataListMarketingScene) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2DataListMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_data_list_marketing_scene value
func (v CampaignGetV2DataListMarketingScene) Ptr() *CampaignGetV2DataListMarketingScene {
	return &v
}
