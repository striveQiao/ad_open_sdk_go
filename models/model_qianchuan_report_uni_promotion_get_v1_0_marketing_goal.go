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

// QianchuanReportUniPromotionGetV10MarketingGoal
type QianchuanReportUniPromotionGetV10MarketingGoal string

// List of qianchuan_report_uni_promotion_get_v1.0_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanReportUniPromotionGetV10MarketingGoal QianchuanReportUniPromotionGetV10MarketingGoal = "LIVE_PROM_GOODS"
)

// All allowed values of QianchuanReportUniPromotionGetV10MarketingGoal enum
var AllowedQianchuanReportUniPromotionGetV10MarketingGoalEnumValues = []QianchuanReportUniPromotionGetV10MarketingGoal{
	"LIVE_PROM_GOODS",
}

// NewQianchuanReportUniPromotionGetV10MarketingGoalFromValue returns a pointer to a valid QianchuanReportUniPromotionGetV10MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportUniPromotionGetV10MarketingGoalFromValue(v string) (*QianchuanReportUniPromotionGetV10MarketingGoal, error) {
	ev := QianchuanReportUniPromotionGetV10MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportUniPromotionGetV10MarketingGoal: valid values are %v", v, AllowedQianchuanReportUniPromotionGetV10MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportUniPromotionGetV10MarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanReportUniPromotionGetV10MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_uni_promotion_get_v1.0_marketing_goal value
func (v QianchuanReportUniPromotionGetV10MarketingGoal) Ptr() *QianchuanReportUniPromotionGetV10MarketingGoal {
	return &v
}
