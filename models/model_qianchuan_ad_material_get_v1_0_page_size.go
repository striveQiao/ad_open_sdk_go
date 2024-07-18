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

// QianchuanAdMaterialGetV10PageSize
type QianchuanAdMaterialGetV10PageSize int64

// List of qianchuan_ad_material_get_v1.0_page_size
const (
	Enum_10_QianchuanAdMaterialGetV10PageSize  QianchuanAdMaterialGetV10PageSize = 10
	Enum_20_QianchuanAdMaterialGetV10PageSize  QianchuanAdMaterialGetV10PageSize = 20
	Enum_50_QianchuanAdMaterialGetV10PageSize  QianchuanAdMaterialGetV10PageSize = 50
	Enum_100_QianchuanAdMaterialGetV10PageSize QianchuanAdMaterialGetV10PageSize = 100
)

// All allowed values of QianchuanAdMaterialGetV10PageSize enum
var AllowedQianchuanAdMaterialGetV10PageSizeEnumValues = []QianchuanAdMaterialGetV10PageSize{
	10,
	20,
	50,
	100,
}

// NewQianchuanAdMaterialGetV10PageSizeFromValue returns a pointer to a valid QianchuanAdMaterialGetV10PageSize
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdMaterialGetV10PageSizeFromValue(v int64) (*QianchuanAdMaterialGetV10PageSize, error) {
	ev := QianchuanAdMaterialGetV10PageSize(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdMaterialGetV10PageSize: valid values are %v", v, AllowedQianchuanAdMaterialGetV10PageSizeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdMaterialGetV10PageSize) IsValid() bool {
	for _, existing := range AllowedQianchuanAdMaterialGetV10PageSizeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_material_get_v1.0_page_size value
func (v QianchuanAdMaterialGetV10PageSize) Ptr() *QianchuanAdMaterialGetV10PageSize {
	return &v
}
