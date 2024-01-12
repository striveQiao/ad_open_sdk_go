/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdGetV10DataListDeliverySettingExternalAction
type QianchuanAdGetV10DataListDeliverySettingExternalAction string

// List of qianchuan_ad_get_v1.0_data_list_delivery_setting_external_action
const (
	AD_CONVERT_TYPE_CARD_ACTIVE_QianchuanAdGetV10DataListDeliverySettingExternalAction                  QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_CARD_ACTIVE"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_QianchuanAdGetV10DataListDeliverySettingExternalAction    QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_QianchuanAdGetV10DataListDeliverySettingExternalAction          QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_QianchuanAdGetV10DataListDeliverySettingExternalAction            QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_PAY_ROI_QianchuanAdGetV10DataListDeliverySettingExternalAction                 QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_PAY_ROI"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_QianchuanAdGetV10DataListDeliverySettingExternalAction     QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanAdGetV10DataListDeliverySettingExternalAction        QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7_DAYS_QianchuanAdGetV10DataListDeliverySettingExternalAction QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE_QianchuanAdGetV10DataListDeliverySettingExternalAction     QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_QianchuanAdGetV10DataListDeliverySettingExternalAction            QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_QianchuanAdGetV10DataListDeliverySettingExternalAction             QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_QianchuanAdGetV10DataListDeliverySettingExternalAction                  QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_SHOPPING_QianchuanAdGetV10DataListDeliverySettingExternalAction                     QianchuanAdGetV10DataListDeliverySettingExternalAction = "AD_CONVERT_TYPE_SHOPPING"
)

// All allowed values of QianchuanAdGetV10DataListDeliverySettingExternalAction enum
var AllowedQianchuanAdGetV10DataListDeliverySettingExternalActionEnumValues = []QianchuanAdGetV10DataListDeliverySettingExternalAction{
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

// NewQianchuanAdGetV10DataListDeliverySettingExternalActionFromValue returns a pointer to a valid QianchuanAdGetV10DataListDeliverySettingExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListDeliverySettingExternalActionFromValue(v string) (*QianchuanAdGetV10DataListDeliverySettingExternalAction, error) {
	ev := QianchuanAdGetV10DataListDeliverySettingExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListDeliverySettingExternalAction: valid values are %v", v, AllowedQianchuanAdGetV10DataListDeliverySettingExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListDeliverySettingExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListDeliverySettingExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_delivery_setting_external_action value
func (v QianchuanAdGetV10DataListDeliverySettingExternalAction) Ptr() *QianchuanAdGetV10DataListDeliverySettingExternalAction {
	return &v
}
