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

// QianchuanAdCreateV10LabAdType
type QianchuanAdCreateV10LabAdType string

// List of qianchuan_ad_create_v1.0_lab_ad_type
const (
	NOT_LAB_AD_QianchuanAdCreateV10LabAdType QianchuanAdCreateV10LabAdType = "NOT_LAB_AD"
	LAB_AD_QianchuanAdCreateV10LabAdType     QianchuanAdCreateV10LabAdType = "LAB_AD"
)

// All allowed values of QianchuanAdCreateV10LabAdType enum
var AllowedQianchuanAdCreateV10LabAdTypeEnumValues = []QianchuanAdCreateV10LabAdType{
	"NOT_LAB_AD",
	"LAB_AD",
}

// NewQianchuanAdCreateV10LabAdTypeFromValue returns a pointer to a valid QianchuanAdCreateV10LabAdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10LabAdTypeFromValue(v string) (*QianchuanAdCreateV10LabAdType, error) {
	ev := QianchuanAdCreateV10LabAdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10LabAdType: valid values are %v", v, AllowedQianchuanAdCreateV10LabAdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10LabAdType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10LabAdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_lab_ad_type value
func (v QianchuanAdCreateV10LabAdType) Ptr() *QianchuanAdCreateV10LabAdType {
	return &v
}
