/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeVideoGetV10MarketingGoal
type QianchuanAwemeVideoGetV10MarketingGoal string

// List of qianchuan_aweme_video_get_v1.0_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanAwemeVideoGetV10MarketingGoal  QianchuanAwemeVideoGetV10MarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanAwemeVideoGetV10MarketingGoal QianchuanAwemeVideoGetV10MarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanAwemeVideoGetV10MarketingGoal enum
var AllowedQianchuanAwemeVideoGetV10MarketingGoalEnumValues = []QianchuanAwemeVideoGetV10MarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanAwemeVideoGetV10MarketingGoalFromValue returns a pointer to a valid QianchuanAwemeVideoGetV10MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeVideoGetV10MarketingGoalFromValue(v string) (*QianchuanAwemeVideoGetV10MarketingGoal, error) {
	ev := QianchuanAwemeVideoGetV10MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeVideoGetV10MarketingGoal: valid values are %v", v, AllowedQianchuanAwemeVideoGetV10MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeVideoGetV10MarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeVideoGetV10MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_video_get_v1.0_marketing_goal value
func (v QianchuanAwemeVideoGetV10MarketingGoal) Ptr() *QianchuanAwemeVideoGetV10MarketingGoal {
	return &v
}
