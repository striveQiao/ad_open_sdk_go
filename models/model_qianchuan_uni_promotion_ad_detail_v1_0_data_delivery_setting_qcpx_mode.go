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

// QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode
type QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode string

// List of qianchuan_uni_promotion_ad_detail_v1.0_data_delivery_setting_qcpx_mode
const (
	QCPX_MODE_DEFAULT_QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode = "QCPX_MODE_DEFAULT"
	QCPX_MODE_OFF_QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode     QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode = "QCPX_MODE_OFF"
	QCPX_MODE_ON_QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode      QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode = "QCPX_MODE_ON"
)

// All allowed values of QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode enum
var AllowedQianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxModeEnumValues = []QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode{
	"QCPX_MODE_DEFAULT",
	"QCPX_MODE_OFF",
	"QCPX_MODE_ON",
}

// NewQianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxModeFromValue returns a pointer to a valid QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxModeFromValue(v string) (*QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode, error) {
	ev := QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode: valid values are %v", v, AllowedQianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_ad_detail_v1.0_data_delivery_setting_qcpx_mode value
func (v QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode) Ptr() *QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode {
	return &v
}
