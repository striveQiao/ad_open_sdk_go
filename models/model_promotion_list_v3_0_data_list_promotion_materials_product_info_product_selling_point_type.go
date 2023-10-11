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

// PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType
type PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType string

// List of promotion_list_v3.0_data_list_promotion_materials_product_info_product_selling_point_type
const (
	CUSTOM_PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType = "CUSTOM"
	DPA_PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType    PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType = "DPA"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType enum
var AllowedPromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointTypeEnumValues = []PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType{
	"CUSTOM",
	"DPA",
}

// NewPromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointTypeFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointTypeFromValue(v string) (*PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType, error) {
	ev := PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_product_info_product_selling_point_type value
func (v PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType) Ptr() *PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType {
	return &v
}
