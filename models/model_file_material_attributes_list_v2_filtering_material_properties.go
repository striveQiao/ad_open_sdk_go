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

// FileMaterialAttributesListV2FilteringMaterialProperties
type FileMaterialAttributesListV2FilteringMaterialProperties string

// List of file_material_attributes_list_v2_filtering_material_properties
const (
	AD_HIGH_QUALITY_MATERIAL_FileMaterialAttributesListV2FilteringMaterialProperties  FileMaterialAttributesListV2FilteringMaterialProperties = "AD_HIGH_QUALITY_MATERIAL"
	AD_LOW_QUALITY_MATERIAL_FileMaterialAttributesListV2FilteringMaterialProperties   FileMaterialAttributesListV2FilteringMaterialProperties = "AD_LOW_QUALITY_MATERIAL"
	AIGC_FileMaterialAttributesListV2FilteringMaterialProperties                      FileMaterialAttributesListV2FilteringMaterialProperties = "AIGC"
	CARRY_MATERIAL_FileMaterialAttributesListV2FilteringMaterialProperties            FileMaterialAttributesListV2FilteringMaterialProperties = "CARRY_MATERIAL"
	ECP_HIGH_QUALITY_MATERIAL_FileMaterialAttributesListV2FilteringMaterialProperties FileMaterialAttributesListV2FilteringMaterialProperties = "ECP_HIGH_QUALITY_MATERIAL"
	ECP_LOW_QUALITY_MATERIAL_FileMaterialAttributesListV2FilteringMaterialProperties  FileMaterialAttributesListV2FilteringMaterialProperties = "ECP_LOW_QUALITY_MATERIAL"
	FIRST_PUBLISH_MATERIAL_FileMaterialAttributesListV2FilteringMaterialProperties    FileMaterialAttributesListV2FilteringMaterialProperties = "FIRST_PUBLISH_MATERIAL"
	INEFFICIENT_MATERIAL_FileMaterialAttributesListV2FilteringMaterialProperties      FileMaterialAttributesListV2FilteringMaterialProperties = "INEFFICIENT_MATERIAL"
	SIMILAR_MATERIAL_FileMaterialAttributesListV2FilteringMaterialProperties          FileMaterialAttributesListV2FilteringMaterialProperties = "SIMILAR_MATERIAL"
	SIMILAR_QUEUE_MATERIAL_FileMaterialAttributesListV2FilteringMaterialProperties    FileMaterialAttributesListV2FilteringMaterialProperties = "SIMILAR_QUEUE_MATERIAL"
)

// All allowed values of FileMaterialAttributesListV2FilteringMaterialProperties enum
var AllowedFileMaterialAttributesListV2FilteringMaterialPropertiesEnumValues = []FileMaterialAttributesListV2FilteringMaterialProperties{
	"AD_HIGH_QUALITY_MATERIAL",
	"AD_LOW_QUALITY_MATERIAL",
	"AIGC",
	"CARRY_MATERIAL",
	"ECP_HIGH_QUALITY_MATERIAL",
	"ECP_LOW_QUALITY_MATERIAL",
	"FIRST_PUBLISH_MATERIAL",
	"INEFFICIENT_MATERIAL",
	"SIMILAR_MATERIAL",
	"SIMILAR_QUEUE_MATERIAL",
}

// NewFileMaterialAttributesListV2FilteringMaterialPropertiesFromValue returns a pointer to a valid FileMaterialAttributesListV2FilteringMaterialProperties
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileMaterialAttributesListV2FilteringMaterialPropertiesFromValue(v string) (*FileMaterialAttributesListV2FilteringMaterialProperties, error) {
	ev := FileMaterialAttributesListV2FilteringMaterialProperties(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileMaterialAttributesListV2FilteringMaterialProperties: valid values are %v", v, AllowedFileMaterialAttributesListV2FilteringMaterialPropertiesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileMaterialAttributesListV2FilteringMaterialProperties) IsValid() bool {
	for _, existing := range AllowedFileMaterialAttributesListV2FilteringMaterialPropertiesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_material_attributes_list_v2_filtering_material_properties value
func (v FileMaterialAttributesListV2FilteringMaterialProperties) Ptr() *FileMaterialAttributesListV2FilteringMaterialProperties {
	return &v
}
