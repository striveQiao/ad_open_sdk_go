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

// QianchuanAwemeOrderCreateV10DeliverySettingBidMode
type QianchuanAwemeOrderCreateV10DeliverySettingBidMode string

// List of qianchuan_aweme_order_create_v1.0_delivery_setting_bid_mode
const (
	PRICING_ACTION_QianchuanAwemeOrderCreateV10DeliverySettingBidMode QianchuanAwemeOrderCreateV10DeliverySettingBidMode = "PRICING_ACTION"
)

// All allowed values of QianchuanAwemeOrderCreateV10DeliverySettingBidMode enum
var AllowedQianchuanAwemeOrderCreateV10DeliverySettingBidModeEnumValues = []QianchuanAwemeOrderCreateV10DeliverySettingBidMode{
	"PRICING_ACTION",
}

// NewQianchuanAwemeOrderCreateV10DeliverySettingBidModeFromValue returns a pointer to a valid QianchuanAwemeOrderCreateV10DeliverySettingBidMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderCreateV10DeliverySettingBidModeFromValue(v string) (*QianchuanAwemeOrderCreateV10DeliverySettingBidMode, error) {
	ev := QianchuanAwemeOrderCreateV10DeliverySettingBidMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderCreateV10DeliverySettingBidMode: valid values are %v", v, AllowedQianchuanAwemeOrderCreateV10DeliverySettingBidModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderCreateV10DeliverySettingBidMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderCreateV10DeliverySettingBidModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_create_v1.0_delivery_setting_bid_mode value
func (v QianchuanAwemeOrderCreateV10DeliverySettingBidMode) Ptr() *QianchuanAwemeOrderCreateV10DeliverySettingBidMode {
	return &v
}
