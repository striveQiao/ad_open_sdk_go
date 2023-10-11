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

// QianchuanSuggestRoiGoalV10DeepExternalAction
type QianchuanSuggestRoiGoalV10DeepExternalAction string

// List of qianchuan_suggest_roi_goal_v1.0_deep_external_action
const (
	AD_CONVERT_TYPE_LIVE_PAY_ROI_QianchuanSuggestRoiGoalV10DeepExternalAction QianchuanSuggestRoiGoalV10DeepExternalAction = "AD_CONVERT_TYPE_LIVE_PAY_ROI"
)

// All allowed values of QianchuanSuggestRoiGoalV10DeepExternalAction enum
var AllowedQianchuanSuggestRoiGoalV10DeepExternalActionEnumValues = []QianchuanSuggestRoiGoalV10DeepExternalAction{
	"AD_CONVERT_TYPE_LIVE_PAY_ROI",
}

// NewQianchuanSuggestRoiGoalV10DeepExternalActionFromValue returns a pointer to a valid QianchuanSuggestRoiGoalV10DeepExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanSuggestRoiGoalV10DeepExternalActionFromValue(v string) (*QianchuanSuggestRoiGoalV10DeepExternalAction, error) {
	ev := QianchuanSuggestRoiGoalV10DeepExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanSuggestRoiGoalV10DeepExternalAction: valid values are %v", v, AllowedQianchuanSuggestRoiGoalV10DeepExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanSuggestRoiGoalV10DeepExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanSuggestRoiGoalV10DeepExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_suggest_roi_goal_v1.0_deep_external_action value
func (v QianchuanSuggestRoiGoalV10DeepExternalAction) Ptr() *QianchuanSuggestRoiGoalV10DeepExternalAction {
	return &v
}
