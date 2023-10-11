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

// FileMaterialListV2PropertiesFilter
type FileMaterialListV2PropertiesFilter string

// List of file_material_list_v2_properties_filter
const (
	AD_HIGH_QUALITY_MATERIAL_FileMaterialListV2PropertiesFilter  FileMaterialListV2PropertiesFilter = "AD_HIGH_QUALITY_MATERIAL"
	ECP_HIGH_QUALITY_MATERIAL_FileMaterialListV2PropertiesFilter FileMaterialListV2PropertiesFilter = "ECP_HIGH_QUALITY_MATERIAL"
	FIRST_PUBLISH_MATERIAL_FileMaterialListV2PropertiesFilter    FileMaterialListV2PropertiesFilter = "FIRST_PUBLISH_MATERIAL"
	INEFFICIENT_MATERIAL_FileMaterialListV2PropertiesFilter      FileMaterialListV2PropertiesFilter = "INEFFICIENT_MATERIAL"
	SIMILAR_MATERIAL_FileMaterialListV2PropertiesFilter          FileMaterialListV2PropertiesFilter = "SIMILAR_MATERIAL"
	SIMILAR_QUEUE_MATERIAL_FileMaterialListV2PropertiesFilter    FileMaterialListV2PropertiesFilter = "SIMILAR_QUEUE_MATERIAL"
)

// All allowed values of FileMaterialListV2PropertiesFilter enum
var AllowedFileMaterialListV2PropertiesFilterEnumValues = []FileMaterialListV2PropertiesFilter{
	"AD_HIGH_QUALITY_MATERIAL",
	"ECP_HIGH_QUALITY_MATERIAL",
	"FIRST_PUBLISH_MATERIAL",
	"INEFFICIENT_MATERIAL",
	"SIMILAR_MATERIAL",
	"SIMILAR_QUEUE_MATERIAL",
}

// NewFileMaterialListV2PropertiesFilterFromValue returns a pointer to a valid FileMaterialListV2PropertiesFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileMaterialListV2PropertiesFilterFromValue(v string) (*FileMaterialListV2PropertiesFilter, error) {
	ev := FileMaterialListV2PropertiesFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileMaterialListV2PropertiesFilter: valid values are %v", v, AllowedFileMaterialListV2PropertiesFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileMaterialListV2PropertiesFilter) IsValid() bool {
	for _, existing := range AllowedFileMaterialListV2PropertiesFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_material_list_v2_properties_filter value
func (v FileMaterialListV2PropertiesFilter) Ptr() *FileMaterialListV2PropertiesFilter {
	return &v
}
