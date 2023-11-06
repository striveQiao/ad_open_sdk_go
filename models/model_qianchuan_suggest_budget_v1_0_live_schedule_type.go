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

// QianchuanSuggestBudgetV10LiveScheduleType
type QianchuanSuggestBudgetV10LiveScheduleType string

// List of qianchuan_suggest_budget_v1.0_live_schedule_type
const (
	SCHEDULE_FROM_NOW_QianchuanSuggestBudgetV10LiveScheduleType        QianchuanSuggestBudgetV10LiveScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_QianchuanSuggestBudgetV10LiveScheduleType       QianchuanSuggestBudgetV10LiveScheduleType = "SCHEDULE_START_END"
	SCHEDULE_TIME_FIXEDRANGE_QianchuanSuggestBudgetV10LiveScheduleType QianchuanSuggestBudgetV10LiveScheduleType = "SCHEDULE_TIME_FIXEDRANGE"
)

// All allowed values of QianchuanSuggestBudgetV10LiveScheduleType enum
var AllowedQianchuanSuggestBudgetV10LiveScheduleTypeEnumValues = []QianchuanSuggestBudgetV10LiveScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
	"SCHEDULE_TIME_FIXEDRANGE",
}

// NewQianchuanSuggestBudgetV10LiveScheduleTypeFromValue returns a pointer to a valid QianchuanSuggestBudgetV10LiveScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanSuggestBudgetV10LiveScheduleTypeFromValue(v string) (*QianchuanSuggestBudgetV10LiveScheduleType, error) {
	ev := QianchuanSuggestBudgetV10LiveScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanSuggestBudgetV10LiveScheduleType: valid values are %v", v, AllowedQianchuanSuggestBudgetV10LiveScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanSuggestBudgetV10LiveScheduleType) IsValid() bool {
	for _, existing := range AllowedQianchuanSuggestBudgetV10LiveScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_suggest_budget_v1.0_live_schedule_type value
func (v QianchuanSuggestBudgetV10LiveScheduleType) Ptr() *QianchuanSuggestBudgetV10LiveScheduleType {
	return &v
}
