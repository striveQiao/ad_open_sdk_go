/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderCreateV10MarketingGoal
type QianchuanAwemeOrderCreateV10MarketingGoal string

// List of qianchuan_aweme_order_create_v1.0_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanAwemeOrderCreateV10MarketingGoal  QianchuanAwemeOrderCreateV10MarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanAwemeOrderCreateV10MarketingGoal QianchuanAwemeOrderCreateV10MarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanAwemeOrderCreateV10MarketingGoal enum
var AllowedQianchuanAwemeOrderCreateV10MarketingGoalEnumValues = []QianchuanAwemeOrderCreateV10MarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanAwemeOrderCreateV10MarketingGoalFromValue returns a pointer to a valid QianchuanAwemeOrderCreateV10MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderCreateV10MarketingGoalFromValue(v string) (*QianchuanAwemeOrderCreateV10MarketingGoal, error) {
	ev := QianchuanAwemeOrderCreateV10MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderCreateV10MarketingGoal: valid values are %v", v, AllowedQianchuanAwemeOrderCreateV10MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderCreateV10MarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderCreateV10MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_create_v1.0_marketing_goal value
func (v QianchuanAwemeOrderCreateV10MarketingGoal) Ptr() *QianchuanAwemeOrderCreateV10MarketingGoal {
	return &v
}
