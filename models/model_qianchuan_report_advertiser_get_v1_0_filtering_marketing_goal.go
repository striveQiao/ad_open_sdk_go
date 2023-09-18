/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportAdvertiserGetV10FilteringMarketingGoal
type QianchuanReportAdvertiserGetV10FilteringMarketingGoal string

// List of qianchuan_report_advertiser_get_v1.0_filtering_marketing_goal
const (
	ALL_QianchuanReportAdvertiserGetV10FilteringMarketingGoal              QianchuanReportAdvertiserGetV10FilteringMarketingGoal = "ALL"
	LIVE_PROM_GOODS_QianchuanReportAdvertiserGetV10FilteringMarketingGoal  QianchuanReportAdvertiserGetV10FilteringMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanReportAdvertiserGetV10FilteringMarketingGoal QianchuanReportAdvertiserGetV10FilteringMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanReportAdvertiserGetV10FilteringMarketingGoal enum
var AllowedQianchuanReportAdvertiserGetV10FilteringMarketingGoalEnumValues = []QianchuanReportAdvertiserGetV10FilteringMarketingGoal{
	"ALL",
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanReportAdvertiserGetV10FilteringMarketingGoalFromValue returns a pointer to a valid QianchuanReportAdvertiserGetV10FilteringMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdvertiserGetV10FilteringMarketingGoalFromValue(v string) (*QianchuanReportAdvertiserGetV10FilteringMarketingGoal, error) {
	ev := QianchuanReportAdvertiserGetV10FilteringMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdvertiserGetV10FilteringMarketingGoal: valid values are %v", v, AllowedQianchuanReportAdvertiserGetV10FilteringMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdvertiserGetV10FilteringMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdvertiserGetV10FilteringMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_advertiser_get_v1.0_filtering_marketing_goal value
func (v QianchuanReportAdvertiserGetV10FilteringMarketingGoal) Ptr() *QianchuanReportAdvertiserGetV10FilteringMarketingGoal {
	return &v
}
