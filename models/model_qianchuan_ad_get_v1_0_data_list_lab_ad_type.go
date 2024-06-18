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

// QianchuanAdGetV10DataListLabAdType
type QianchuanAdGetV10DataListLabAdType string

// List of qianchuan_ad_get_v1.0_data_list_lab_ad_type
const (
	LAB_AD_QianchuanAdGetV10DataListLabAdType     QianchuanAdGetV10DataListLabAdType = "LAB_AD"
	NOT_LAB_AD_QianchuanAdGetV10DataListLabAdType QianchuanAdGetV10DataListLabAdType = "NOT_LAB_AD"
)

// All allowed values of QianchuanAdGetV10DataListLabAdType enum
var AllowedQianchuanAdGetV10DataListLabAdTypeEnumValues = []QianchuanAdGetV10DataListLabAdType{
	"LAB_AD",
	"NOT_LAB_AD",
}

// NewQianchuanAdGetV10DataListLabAdTypeFromValue returns a pointer to a valid QianchuanAdGetV10DataListLabAdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListLabAdTypeFromValue(v string) (*QianchuanAdGetV10DataListLabAdType, error) {
	ev := QianchuanAdGetV10DataListLabAdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListLabAdType: valid values are %v", v, AllowedQianchuanAdGetV10DataListLabAdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListLabAdType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListLabAdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_lab_ad_type value
func (v QianchuanAdGetV10DataListLabAdType) Ptr() *QianchuanAdGetV10DataListLabAdType {
	return &v
}
