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

// PromotionCreateV30PromotionMaterialsParamsType
type PromotionCreateV30PromotionMaterialsParamsType string

// List of promotion_create_v3.0_promotion_materials_params_type
const (
	CUSTOM_PromotionCreateV30PromotionMaterialsParamsType PromotionCreateV30PromotionMaterialsParamsType = "CUSTOM"
	DPA_PromotionCreateV30PromotionMaterialsParamsType    PromotionCreateV30PromotionMaterialsParamsType = "DPA"
)

// All allowed values of PromotionCreateV30PromotionMaterialsParamsType enum
var AllowedPromotionCreateV30PromotionMaterialsParamsTypeEnumValues = []PromotionCreateV30PromotionMaterialsParamsType{
	"CUSTOM",
	"DPA",
}

// NewPromotionCreateV30PromotionMaterialsParamsTypeFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsParamsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsParamsTypeFromValue(v string) (*PromotionCreateV30PromotionMaterialsParamsType, error) {
	ev := PromotionCreateV30PromotionMaterialsParamsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsParamsType: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsParamsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsParamsType) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsParamsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_params_type value
func (v PromotionCreateV30PromotionMaterialsParamsType) Ptr() *PromotionCreateV30PromotionMaterialsParamsType {
	return &v
}
