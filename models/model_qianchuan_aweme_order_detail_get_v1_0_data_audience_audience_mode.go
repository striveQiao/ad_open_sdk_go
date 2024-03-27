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

// QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode
type QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode string

// List of qianchuan_aweme_order_detail_get_v1.0_data_audience_audience_mode
const (
	ATUO_QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode     QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode = "ATUO"
	CUSTOM_QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode   QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode = "CUSTOM"
	FANS_QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode     QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode = "FANS"
	LIVEFANS_QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode = "LIVEFANS"
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode enum
var AllowedQianchuanAwemeOrderDetailGetV10DataAudienceAudienceModeEnumValues = []QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode{
	"ATUO",
	"CUSTOM",
	"FANS",
	"LIVEFANS",
}

// NewQianchuanAwemeOrderDetailGetV10DataAudienceAudienceModeFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataAudienceAudienceModeFromValue(v string) (*QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataAudienceAudienceModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataAudienceAudienceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_audience_audience_mode value
func (v QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode) Ptr() *QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode {
	return &v
}
