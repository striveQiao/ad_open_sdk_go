/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative
type QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative int64

// List of qianchuan_ad_update_v1.0_multi_product_creative_list_creative_setting_dynamic_creative
const (
	Enum_0_QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative = 0
	Enum_1_QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative = 1
)

// All allowed values of QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative enum
var AllowedQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreativeEnumValues = []QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative{
	0,
	1,
}

// NewQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreativeFromValue returns a pointer to a valid QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreativeFromValue(v int64) (*QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative, error) {
	ev := QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative: valid values are %v", v, AllowedQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreativeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreativeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_multi_product_creative_list_creative_setting_dynamic_creative value
func (v QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative) Ptr() *QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative {
	return &v
}
