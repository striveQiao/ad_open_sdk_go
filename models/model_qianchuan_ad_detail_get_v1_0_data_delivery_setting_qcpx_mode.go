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

// QianchuanAdDetailGetV10DataDeliverySettingQcpxMode
type QianchuanAdDetailGetV10DataDeliverySettingQcpxMode string

// List of qianchuan_ad_detail_get_v1.0_data_delivery_setting_qcpx_mode
const (
	QCPX_MODE_DEFAULT_QianchuanAdDetailGetV10DataDeliverySettingQcpxMode QianchuanAdDetailGetV10DataDeliverySettingQcpxMode = "QCPX_MODE_DEFAULT"
	QCPX_MODE_OFF_QianchuanAdDetailGetV10DataDeliverySettingQcpxMode     QianchuanAdDetailGetV10DataDeliverySettingQcpxMode = "QCPX_MODE_OFF"
	QCPX_MODE_ON_QianchuanAdDetailGetV10DataDeliverySettingQcpxMode      QianchuanAdDetailGetV10DataDeliverySettingQcpxMode = "QCPX_MODE_ON"
)

// All allowed values of QianchuanAdDetailGetV10DataDeliverySettingQcpxMode enum
var AllowedQianchuanAdDetailGetV10DataDeliverySettingQcpxModeEnumValues = []QianchuanAdDetailGetV10DataDeliverySettingQcpxMode{
	"QCPX_MODE_DEFAULT",
	"QCPX_MODE_OFF",
	"QCPX_MODE_ON",
}

// NewQianchuanAdDetailGetV10DataDeliverySettingQcpxModeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataDeliverySettingQcpxMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataDeliverySettingQcpxModeFromValue(v string) (*QianchuanAdDetailGetV10DataDeliverySettingQcpxMode, error) {
	ev := QianchuanAdDetailGetV10DataDeliverySettingQcpxMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataDeliverySettingQcpxMode: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataDeliverySettingQcpxModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataDeliverySettingQcpxMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataDeliverySettingQcpxModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_delivery_setting_qcpx_mode value
func (v QianchuanAdDetailGetV10DataDeliverySettingQcpxMode) Ptr() *QianchuanAdDetailGetV10DataDeliverySettingQcpxMode {
	return &v
}
