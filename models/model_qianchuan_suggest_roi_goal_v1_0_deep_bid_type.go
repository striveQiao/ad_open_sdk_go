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

// QianchuanSuggestRoiGoalV10DeepBidType
type QianchuanSuggestRoiGoalV10DeepBidType string

// List of qianchuan_suggest_roi_goal_v1.0_deep_bid_type
const (
	COMM_ROI_QianchuanSuggestRoiGoalV10DeepBidType QianchuanSuggestRoiGoalV10DeepBidType = "COMM_ROI"
	MIN_QianchuanSuggestRoiGoalV10DeepBidType      QianchuanSuggestRoiGoalV10DeepBidType = "MIN"
	PAY_ROI_QianchuanSuggestRoiGoalV10DeepBidType  QianchuanSuggestRoiGoalV10DeepBidType = "PAY_ROI"
)

// All allowed values of QianchuanSuggestRoiGoalV10DeepBidType enum
var AllowedQianchuanSuggestRoiGoalV10DeepBidTypeEnumValues = []QianchuanSuggestRoiGoalV10DeepBidType{
	"COMM_ROI",
	"MIN",
	"PAY_ROI",
}

// NewQianchuanSuggestRoiGoalV10DeepBidTypeFromValue returns a pointer to a valid QianchuanSuggestRoiGoalV10DeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanSuggestRoiGoalV10DeepBidTypeFromValue(v string) (*QianchuanSuggestRoiGoalV10DeepBidType, error) {
	ev := QianchuanSuggestRoiGoalV10DeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanSuggestRoiGoalV10DeepBidType: valid values are %v", v, AllowedQianchuanSuggestRoiGoalV10DeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanSuggestRoiGoalV10DeepBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanSuggestRoiGoalV10DeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_suggest_roi_goal_v1.0_deep_bid_type value
func (v QianchuanSuggestRoiGoalV10DeepBidType) Ptr() *QianchuanSuggestRoiGoalV10DeepBidType {
	return &v
}
