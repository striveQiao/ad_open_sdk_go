/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanSuggestRoiGoalV10ExternalAction
type QianchuanSuggestRoiGoalV10ExternalAction string

// List of qianchuan_suggest_roi_goal_v1.0_external_action
const (
	AD_CONVERT_TYPE_CARD_ACTIVE_QianchuanSuggestRoiGoalV10ExternalAction                  QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_CARD_ACTIVE"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_QianchuanSuggestRoiGoalV10ExternalAction    QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_QianchuanSuggestRoiGoalV10ExternalAction          QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_QianchuanSuggestRoiGoalV10ExternalAction            QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_PAY_ROI_QianchuanSuggestRoiGoalV10ExternalAction                 QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_LIVE_PAY_ROI"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_QianchuanSuggestRoiGoalV10ExternalAction     QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanSuggestRoiGoalV10ExternalAction        QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7_DAYS_QianchuanSuggestRoiGoalV10ExternalAction QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE_QianchuanSuggestRoiGoalV10ExternalAction     QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_QianchuanSuggestRoiGoalV10ExternalAction            QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_QianchuanSuggestRoiGoalV10ExternalAction             QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_QianchuanSuggestRoiGoalV10ExternalAction                  QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_SHOPPING_QianchuanSuggestRoiGoalV10ExternalAction                     QianchuanSuggestRoiGoalV10ExternalAction = "AD_CONVERT_TYPE_SHOPPING"
)

// All allowed values of QianchuanSuggestRoiGoalV10ExternalAction enum
var AllowedQianchuanSuggestRoiGoalV10ExternalActionEnumValues = []QianchuanSuggestRoiGoalV10ExternalAction{
	"AD_CONVERT_TYPE_CARD_ACTIVE",
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_LIVE_ENTER_ACTION",
	"AD_CONVERT_TYPE_LIVE_PAY_ROI",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_SHOPPING",
}

// NewQianchuanSuggestRoiGoalV10ExternalActionFromValue returns a pointer to a valid QianchuanSuggestRoiGoalV10ExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanSuggestRoiGoalV10ExternalActionFromValue(v string) (*QianchuanSuggestRoiGoalV10ExternalAction, error) {
	ev := QianchuanSuggestRoiGoalV10ExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanSuggestRoiGoalV10ExternalAction: valid values are %v", v, AllowedQianchuanSuggestRoiGoalV10ExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanSuggestRoiGoalV10ExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanSuggestRoiGoalV10ExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_suggest_roi_goal_v1.0_external_action value
func (v QianchuanSuggestRoiGoalV10ExternalAction) Ptr() *QianchuanSuggestRoiGoalV10ExternalAction {
	return &v
}
