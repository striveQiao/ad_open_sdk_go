/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeReportOrderGetV10DataListMarketingGoal
type QianchuanAwemeReportOrderGetV10DataListMarketingGoal string

// List of qianchuan_aweme_report_order_get_v1.0_data_list_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanAwemeReportOrderGetV10DataListMarketingGoal  QianchuanAwemeReportOrderGetV10DataListMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanAwemeReportOrderGetV10DataListMarketingGoal QianchuanAwemeReportOrderGetV10DataListMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanAwemeReportOrderGetV10DataListMarketingGoal enum
var AllowedQianchuanAwemeReportOrderGetV10DataListMarketingGoalEnumValues = []QianchuanAwemeReportOrderGetV10DataListMarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanAwemeReportOrderGetV10DataListMarketingGoalFromValue returns a pointer to a valid QianchuanAwemeReportOrderGetV10DataListMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeReportOrderGetV10DataListMarketingGoalFromValue(v string) (*QianchuanAwemeReportOrderGetV10DataListMarketingGoal, error) {
	ev := QianchuanAwemeReportOrderGetV10DataListMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeReportOrderGetV10DataListMarketingGoal: valid values are %v", v, AllowedQianchuanAwemeReportOrderGetV10DataListMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeReportOrderGetV10DataListMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeReportOrderGetV10DataListMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_report_order_get_v1.0_data_list_marketing_goal value
func (v QianchuanAwemeReportOrderGetV10DataListMarketingGoal) Ptr() *QianchuanAwemeReportOrderGetV10DataListMarketingGoal {
	return &v
}
