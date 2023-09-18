/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10CreativeListTitleMaterialTitleType
type QianchuanAdUpdateV10CreativeListTitleMaterialTitleType string

// List of qianchuan_ad_update_v1.0_creative_list_title_material_title_type
const (
	COMMODITY_CARD_QianchuanAdUpdateV10CreativeListTitleMaterialTitleType QianchuanAdUpdateV10CreativeListTitleMaterialTitleType = "COMMODITY_CARD"
	CUSTOM_QianchuanAdUpdateV10CreativeListTitleMaterialTitleType         QianchuanAdUpdateV10CreativeListTitleMaterialTitleType = "CUSTOM"
)

// All allowed values of QianchuanAdUpdateV10CreativeListTitleMaterialTitleType enum
var AllowedQianchuanAdUpdateV10CreativeListTitleMaterialTitleTypeEnumValues = []QianchuanAdUpdateV10CreativeListTitleMaterialTitleType{
	"COMMODITY_CARD",
	"CUSTOM",
}

// NewQianchuanAdUpdateV10CreativeListTitleMaterialTitleTypeFromValue returns a pointer to a valid QianchuanAdUpdateV10CreativeListTitleMaterialTitleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10CreativeListTitleMaterialTitleTypeFromValue(v string) (*QianchuanAdUpdateV10CreativeListTitleMaterialTitleType, error) {
	ev := QianchuanAdUpdateV10CreativeListTitleMaterialTitleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10CreativeListTitleMaterialTitleType: valid values are %v", v, AllowedQianchuanAdUpdateV10CreativeListTitleMaterialTitleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10CreativeListTitleMaterialTitleType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10CreativeListTitleMaterialTitleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_creative_list_title_material_title_type value
func (v QianchuanAdUpdateV10CreativeListTitleMaterialTitleType) Ptr() *QianchuanAdUpdateV10CreativeListTitleMaterialTitleType {
	return &v
}
