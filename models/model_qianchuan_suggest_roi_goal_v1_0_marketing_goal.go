/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanSuggestRoiGoalV10MarketingGoal
type QianchuanSuggestRoiGoalV10MarketingGoal string

// List of qianchuan_suggest_roi_goal_v1.0_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanSuggestRoiGoalV10MarketingGoal  QianchuanSuggestRoiGoalV10MarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanSuggestRoiGoalV10MarketingGoal QianchuanSuggestRoiGoalV10MarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanSuggestRoiGoalV10MarketingGoal enum
var AllowedQianchuanSuggestRoiGoalV10MarketingGoalEnumValues = []QianchuanSuggestRoiGoalV10MarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanSuggestRoiGoalV10MarketingGoalFromValue returns a pointer to a valid QianchuanSuggestRoiGoalV10MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanSuggestRoiGoalV10MarketingGoalFromValue(v string) (*QianchuanSuggestRoiGoalV10MarketingGoal, error) {
	ev := QianchuanSuggestRoiGoalV10MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanSuggestRoiGoalV10MarketingGoal: valid values are %v", v, AllowedQianchuanSuggestRoiGoalV10MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanSuggestRoiGoalV10MarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanSuggestRoiGoalV10MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_suggest_roi_goal_v1.0_marketing_goal value
func (v QianchuanSuggestRoiGoalV10MarketingGoal) Ptr() *QianchuanSuggestRoiGoalV10MarketingGoal {
	return &v
}
