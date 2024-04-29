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

// QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode
type QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode string

// List of qianchuan_aweme_order_detail_get_v1.0_data_delivery_setting_bid_mode
const (
	PRICING_ACTION_QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode = "PRICING_ACTION"
	PRICING_PLAY_QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode   QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode = "PRICING_PLAY"
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode enum
var AllowedQianchuanAwemeOrderDetailGetV10DataDeliverySettingBidModeEnumValues = []QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode{
	"PRICING_ACTION",
	"PRICING_PLAY",
}

// NewQianchuanAwemeOrderDetailGetV10DataDeliverySettingBidModeFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataDeliverySettingBidModeFromValue(v string) (*QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataDeliverySettingBidModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataDeliverySettingBidModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_delivery_setting_bid_mode value
func (v QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode) Ptr() *QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode {
	return &v
}
