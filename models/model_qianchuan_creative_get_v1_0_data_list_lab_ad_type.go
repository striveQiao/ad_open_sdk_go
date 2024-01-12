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

// QianchuanCreativeGetV10DataListLabAdType
type QianchuanCreativeGetV10DataListLabAdType string

// List of qianchuan_creative_get_v1.0_data_list_lab_ad_type
const (
	LAB_AD_QianchuanCreativeGetV10DataListLabAdType     QianchuanCreativeGetV10DataListLabAdType = "LAB_AD"
	NOT_LAB_AD_QianchuanCreativeGetV10DataListLabAdType QianchuanCreativeGetV10DataListLabAdType = "NOT_LAB_AD"
)

// All allowed values of QianchuanCreativeGetV10DataListLabAdType enum
var AllowedQianchuanCreativeGetV10DataListLabAdTypeEnumValues = []QianchuanCreativeGetV10DataListLabAdType{
	"LAB_AD",
	"NOT_LAB_AD",
}

// NewQianchuanCreativeGetV10DataListLabAdTypeFromValue returns a pointer to a valid QianchuanCreativeGetV10DataListLabAdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10DataListLabAdTypeFromValue(v string) (*QianchuanCreativeGetV10DataListLabAdType, error) {
	ev := QianchuanCreativeGetV10DataListLabAdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10DataListLabAdType: valid values are %v", v, AllowedQianchuanCreativeGetV10DataListLabAdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10DataListLabAdType) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10DataListLabAdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_data_list_lab_ad_type value
func (v QianchuanCreativeGetV10DataListLabAdType) Ptr() *QianchuanCreativeGetV10DataListLabAdType {
	return &v
}
