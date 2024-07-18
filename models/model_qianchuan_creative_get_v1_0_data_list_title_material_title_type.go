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

// QianchuanCreativeGetV10DataListTitleMaterialTitleType
type QianchuanCreativeGetV10DataListTitleMaterialTitleType string

// List of qianchuan_creative_get_v1.0_data_list_title_material_title_type
const (
	AWEME_CAROUSEL_QianchuanCreativeGetV10DataListTitleMaterialTitleType QianchuanCreativeGetV10DataListTitleMaterialTitleType = "AWEME_CAROUSEL"
	COMMODITY_CARD_QianchuanCreativeGetV10DataListTitleMaterialTitleType QianchuanCreativeGetV10DataListTitleMaterialTitleType = "COMMODITY_CARD"
	CUSTOM_QianchuanCreativeGetV10DataListTitleMaterialTitleType         QianchuanCreativeGetV10DataListTitleMaterialTitleType = "CUSTOM"
)

// All allowed values of QianchuanCreativeGetV10DataListTitleMaterialTitleType enum
var AllowedQianchuanCreativeGetV10DataListTitleMaterialTitleTypeEnumValues = []QianchuanCreativeGetV10DataListTitleMaterialTitleType{
	"AWEME_CAROUSEL",
	"COMMODITY_CARD",
	"CUSTOM",
}

// NewQianchuanCreativeGetV10DataListTitleMaterialTitleTypeFromValue returns a pointer to a valid QianchuanCreativeGetV10DataListTitleMaterialTitleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10DataListTitleMaterialTitleTypeFromValue(v string) (*QianchuanCreativeGetV10DataListTitleMaterialTitleType, error) {
	ev := QianchuanCreativeGetV10DataListTitleMaterialTitleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10DataListTitleMaterialTitleType: valid values are %v", v, AllowedQianchuanCreativeGetV10DataListTitleMaterialTitleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10DataListTitleMaterialTitleType) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10DataListTitleMaterialTitleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_data_list_title_material_title_type value
func (v QianchuanCreativeGetV10DataListTitleMaterialTitleType) Ptr() *QianchuanCreativeGetV10DataListTitleMaterialTitleType {
	return &v
}
