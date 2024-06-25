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

// QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate
type QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate int64

// List of qianchuan_ad_update_v1.0_multi_product_creative_list_creative_setting_creative_auto_generate
const (
	Enum_0_QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate = 0
	Enum_1_QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate = 1
)

// All allowed values of QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate enum
var AllowedQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerateEnumValues = []QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate{
	0,
	1,
}

// NewQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerateFromValue returns a pointer to a valid QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerateFromValue(v int64) (*QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate, error) {
	ev := QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate: valid values are %v", v, AllowedQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_multi_product_creative_list_creative_setting_creative_auto_generate value
func (v QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate) Ptr() *QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate {
	return &v
}
