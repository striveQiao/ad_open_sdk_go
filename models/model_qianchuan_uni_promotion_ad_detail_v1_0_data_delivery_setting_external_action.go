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

// QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction
type QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction string

// List of qianchuan_uni_promotion_ad_detail_v1.0_data_delivery_setting_external_action
const (
	AD_CONVERT_TYPE_CARD_ACTIVE_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction                  QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_CARD_ACTIVE"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction    QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction          QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction            QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_PAY_ROI_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction                 QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_PAY_ROI"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction     QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction        QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7_DAYS_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction     QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction            QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction             QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction                  QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_SHOPPING_QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction                     QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_SHOPPING"
)

// All allowed values of QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction enum
var AllowedQianchuanUniPromotionAdDetailV10DataDeliverySettingExternalActionEnumValues = []QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction{
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

// NewQianchuanUniPromotionAdDetailV10DataDeliverySettingExternalActionFromValue returns a pointer to a valid QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionAdDetailV10DataDeliverySettingExternalActionFromValue(v string) (*QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction, error) {
	ev := QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction: valid values are %v", v, AllowedQianchuanUniPromotionAdDetailV10DataDeliverySettingExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionAdDetailV10DataDeliverySettingExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_ad_detail_v1.0_data_delivery_setting_external_action value
func (v QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction) Ptr() *QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction {
	return &v
}
