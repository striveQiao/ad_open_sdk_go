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

// PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend
type PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend string

// List of promotion_create_v3.0_promotion_materials_blue_flow_material_recommend
const (
	OFF_PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend = "OFF"
	ON_PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend  PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend = "ON"
)

// All allowed values of PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend enum
var AllowedPromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommendEnumValues = []PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend{
	"OFF",
	"ON",
}

// NewPromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommendFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommendFromValue(v string) (*PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend, error) {
	ev := PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommendEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_blue_flow_material_recommend value
func (v PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend) Ptr() *PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend {
	return &v
}
