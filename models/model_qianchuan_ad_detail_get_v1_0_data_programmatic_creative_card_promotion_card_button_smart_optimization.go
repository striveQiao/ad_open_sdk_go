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

// QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization
type QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization int64

// List of qianchuan_ad_detail_get_v1.0_data_programmatic_creative_card_promotion_card_button_smart_optimization
const (
	Enum_0_QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization = 0
	Enum_1_QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization = 1
)

// All allowed values of QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization enum
var AllowedQianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimizationEnumValues = []QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization{
	0,
	1,
}

// NewQianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimizationFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimizationFromValue(v int64) (*QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization, error) {
	ev := QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimizationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_programmatic_creative_card_promotion_card_button_smart_optimization value
func (v QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization) Ptr() *QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization {
	return &v
}
