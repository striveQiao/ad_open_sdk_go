/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30PromotionMaterialsOpenUrlType
type PromotionCreateV30PromotionMaterialsOpenUrlType string

// List of promotion_create_v3.0_promotion_materials_open_url_type
const (
	CUSTOM_PromotionCreateV30PromotionMaterialsOpenUrlType PromotionCreateV30PromotionMaterialsOpenUrlType = "CUSTOM"
	DPA_PromotionCreateV30PromotionMaterialsOpenUrlType    PromotionCreateV30PromotionMaterialsOpenUrlType = "DPA"
	NONE_PromotionCreateV30PromotionMaterialsOpenUrlType   PromotionCreateV30PromotionMaterialsOpenUrlType = "NONE"
)

// All allowed values of PromotionCreateV30PromotionMaterialsOpenUrlType enum
var AllowedPromotionCreateV30PromotionMaterialsOpenUrlTypeEnumValues = []PromotionCreateV30PromotionMaterialsOpenUrlType{
	"CUSTOM",
	"DPA",
	"NONE",
}

// NewPromotionCreateV30PromotionMaterialsOpenUrlTypeFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsOpenUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsOpenUrlTypeFromValue(v string) (*PromotionCreateV30PromotionMaterialsOpenUrlType, error) {
	ev := PromotionCreateV30PromotionMaterialsOpenUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsOpenUrlType: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsOpenUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsOpenUrlType) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsOpenUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_open_url_type value
func (v PromotionCreateV30PromotionMaterialsOpenUrlType) Ptr() *PromotionCreateV30PromotionMaterialsOpenUrlType {
	return &v
}
