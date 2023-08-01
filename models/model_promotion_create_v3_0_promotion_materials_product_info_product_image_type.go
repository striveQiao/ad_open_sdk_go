/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30PromotionMaterialsProductInfoProductImageType
type PromotionCreateV30PromotionMaterialsProductInfoProductImageType string

// List of promotion_create_v3.0_promotion_materials_product_info_product_image_type
const (
	CUSTOM_PromotionCreateV30PromotionMaterialsProductInfoProductImageType PromotionCreateV30PromotionMaterialsProductInfoProductImageType = "CUSTOM"
	DPA_PromotionCreateV30PromotionMaterialsProductInfoProductImageType    PromotionCreateV30PromotionMaterialsProductInfoProductImageType = "DPA"
)

// All allowed values of PromotionCreateV30PromotionMaterialsProductInfoProductImageType enum
var AllowedPromotionCreateV30PromotionMaterialsProductInfoProductImageTypeEnumValues = []PromotionCreateV30PromotionMaterialsProductInfoProductImageType{
	"CUSTOM",
	"DPA",
}

// NewPromotionCreateV30PromotionMaterialsProductInfoProductImageTypeFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsProductInfoProductImageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsProductInfoProductImageTypeFromValue(v string) (*PromotionCreateV30PromotionMaterialsProductInfoProductImageType, error) {
	ev := PromotionCreateV30PromotionMaterialsProductInfoProductImageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsProductInfoProductImageType: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsProductInfoProductImageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsProductInfoProductImageType) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsProductInfoProductImageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_product_info_product_image_type value
func (v PromotionCreateV30PromotionMaterialsProductInfoProductImageType) Ptr() *PromotionCreateV30PromotionMaterialsProductInfoProductImageType {
	return &v
}
