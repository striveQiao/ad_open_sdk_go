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

// QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization
type QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization int64

// List of qianchuan_ad_create_v1.0_creative_list_promotion_card_material_button_smart_optimization
const (
	Enum_0_QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization = 0
	Enum_1_QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization = 1
)

// All allowed values of QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization enum
var AllowedQianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimizationEnumValues = []QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization{
	0,
	1,
}

// NewQianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimizationFromValue returns a pointer to a valid QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimizationFromValue(v int64) (*QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization, error) {
	ev := QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization: valid values are %v", v, AllowedQianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimizationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_creative_list_promotion_card_material_button_smart_optimization value
func (v QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization) Ptr() *QianchuanAdCreateV10CreativeListPromotionCardMaterialButtonSmartOptimization {
	return &v
}
