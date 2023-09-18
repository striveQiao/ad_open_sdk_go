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

// PromotionListV30DataListPromotionMaterialsParamsType
type PromotionListV30DataListPromotionMaterialsParamsType string

// List of promotion_list_v3.0_data_list_promotion_materials_params_type
const (
	CUSTOM_PromotionListV30DataListPromotionMaterialsParamsType PromotionListV30DataListPromotionMaterialsParamsType = "CUSTOM"
	DPA_PromotionListV30DataListPromotionMaterialsParamsType    PromotionListV30DataListPromotionMaterialsParamsType = "DPA"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsParamsType enum
var AllowedPromotionListV30DataListPromotionMaterialsParamsTypeEnumValues = []PromotionListV30DataListPromotionMaterialsParamsType{
	"CUSTOM",
	"DPA",
}

// NewPromotionListV30DataListPromotionMaterialsParamsTypeFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsParamsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsParamsTypeFromValue(v string) (*PromotionListV30DataListPromotionMaterialsParamsType, error) {
	ev := PromotionListV30DataListPromotionMaterialsParamsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsParamsType: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsParamsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsParamsType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsParamsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_params_type value
func (v PromotionListV30DataListPromotionMaterialsParamsType) Ptr() *PromotionListV30DataListPromotionMaterialsParamsType {
	return &v
}
