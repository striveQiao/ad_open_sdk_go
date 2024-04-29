/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend
type PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend string

// List of promotion_list_v3.0_data_list_promotion_materials_blue_flow_material_recommend
const (
	OFF_PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend = "OFF"
	ON_PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend  PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend = "ON"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend enum
var AllowedPromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommendEnumValues = []PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend{
	"OFF",
	"ON",
}

// NewPromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommendFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommendFromValue(v string) (*PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend, error) {
	ev := PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommendEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_blue_flow_material_recommend value
func (v PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend) Ptr() *PromotionListV30DataListPromotionMaterialsBlueFlowMaterialRecommend {
	return &v
}
