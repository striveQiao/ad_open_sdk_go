/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanSuggestRoiGoalV10CampaignScene
type QianchuanSuggestRoiGoalV10CampaignScene string

// List of qianchuan_suggest_roi_goal_v1.0_campaign_scene
const (
	DAILY_SALE_QianchuanSuggestRoiGoalV10CampaignScene                  QianchuanSuggestRoiGoalV10CampaignScene = "DAILY_SALE"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanSuggestRoiGoalV10CampaignScene QianchuanSuggestRoiGoalV10CampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
)

// All allowed values of QianchuanSuggestRoiGoalV10CampaignScene enum
var AllowedQianchuanSuggestRoiGoalV10CampaignSceneEnumValues = []QianchuanSuggestRoiGoalV10CampaignScene{
	"DAILY_SALE",
	"NEW_CUSTOMER_TRANSFORMATION",
}

// NewQianchuanSuggestRoiGoalV10CampaignSceneFromValue returns a pointer to a valid QianchuanSuggestRoiGoalV10CampaignScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanSuggestRoiGoalV10CampaignSceneFromValue(v string) (*QianchuanSuggestRoiGoalV10CampaignScene, error) {
	ev := QianchuanSuggestRoiGoalV10CampaignScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanSuggestRoiGoalV10CampaignScene: valid values are %v", v, AllowedQianchuanSuggestRoiGoalV10CampaignSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanSuggestRoiGoalV10CampaignScene) IsValid() bool {
	for _, existing := range AllowedQianchuanSuggestRoiGoalV10CampaignSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_suggest_roi_goal_v1.0_campaign_scene value
func (v QianchuanSuggestRoiGoalV10CampaignScene) Ptr() *QianchuanSuggestRoiGoalV10CampaignScene {
	return &v
}
