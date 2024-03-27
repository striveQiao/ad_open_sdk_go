/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdGetV10DataListMarketingGoal
type QianchuanAdGetV10DataListMarketingGoal string

// List of qianchuan_ad_get_v1.0_data_list_marketing_goal
const (
	ALL_QianchuanAdGetV10DataListMarketingGoal              QianchuanAdGetV10DataListMarketingGoal = "ALL"
	LIVE_PROM_GOODS_QianchuanAdGetV10DataListMarketingGoal  QianchuanAdGetV10DataListMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanAdGetV10DataListMarketingGoal QianchuanAdGetV10DataListMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanAdGetV10DataListMarketingGoal enum
var AllowedQianchuanAdGetV10DataListMarketingGoalEnumValues = []QianchuanAdGetV10DataListMarketingGoal{
	"ALL",
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanAdGetV10DataListMarketingGoalFromValue returns a pointer to a valid QianchuanAdGetV10DataListMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListMarketingGoalFromValue(v string) (*QianchuanAdGetV10DataListMarketingGoal, error) {
	ev := QianchuanAdGetV10DataListMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListMarketingGoal: valid values are %v", v, AllowedQianchuanAdGetV10DataListMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_marketing_goal value
func (v QianchuanAdGetV10DataListMarketingGoal) Ptr() *QianchuanAdGetV10DataListMarketingGoal {
	return &v
}
