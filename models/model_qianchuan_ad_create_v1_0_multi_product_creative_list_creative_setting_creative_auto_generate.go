/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate
type QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate int64

// List of qianchuan_ad_create_v1.0_multi_product_creative_list_creative_setting_creative_auto_generate
const (
	Enum_0_QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate = 0
	Enum_1_QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate = 1
)

// All allowed values of QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate enum
var AllowedQianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerateEnumValues = []QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate{
	0,
	1,
}

// NewQianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerateFromValue returns a pointer to a valid QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerateFromValue(v int64) (*QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate, error) {
	ev := QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate: valid values are %v", v, AllowedQianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_multi_product_creative_list_creative_setting_creative_auto_generate value
func (v QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate) Ptr() *QianchuanAdCreateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate {
	return &v
}
