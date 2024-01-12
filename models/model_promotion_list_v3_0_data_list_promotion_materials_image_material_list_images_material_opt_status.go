/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus
type PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus string

// List of promotion_list_v3.0_data_list_promotion_materials_image_material_list_images_material_opt_status
const (
	DISABLE_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus = "DISABLE"
	ENABLE_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus  PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus = "ENABLE"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus enum
var AllowedPromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatusEnumValues = []PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus{
	"DISABLE",
	"ENABLE",
}

// NewPromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatusFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatusFromValue(v string) (*PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus, error) {
	ev := PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_image_material_list_images_material_opt_status value
func (v PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus) Ptr() *PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus {
	return &v
}
