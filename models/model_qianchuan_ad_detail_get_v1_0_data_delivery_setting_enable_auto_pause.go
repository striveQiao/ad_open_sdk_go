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

// QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause
type QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause int64

// List of qianchuan_ad_detail_get_v1.0_data_delivery_setting_enable_auto_pause
const (
	Enum_0_QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause = 0
	Enum_1_QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause = 1
)

// All allowed values of QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause enum
var AllowedQianchuanAdDetailGetV10DataDeliverySettingEnableAutoPauseEnumValues = []QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause{
	0,
	1,
}

// NewQianchuanAdDetailGetV10DataDeliverySettingEnableAutoPauseFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataDeliverySettingEnableAutoPauseFromValue(v int64) (*QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause, error) {
	ev := QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataDeliverySettingEnableAutoPauseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataDeliverySettingEnableAutoPauseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_delivery_setting_enable_auto_pause value
func (v QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause) Ptr() *QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause {
	return &v
}
