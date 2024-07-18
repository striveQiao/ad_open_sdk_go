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

// LocalReportMaterialGetV30DataMaterialListMaterialType
type LocalReportMaterialGetV30DataMaterialListMaterialType string

// List of local_report_material_get_v3.0_data_material_list_material_type
const (
	CASURAL_LocalReportMaterialGetV30DataMaterialListMaterialType LocalReportMaterialGetV30DataMaterialListMaterialType = "CASURAL"
	VIDEO_LocalReportMaterialGetV30DataMaterialListMaterialType   LocalReportMaterialGetV30DataMaterialListMaterialType = "VIDEO"
)

// All allowed values of LocalReportMaterialGetV30DataMaterialListMaterialType enum
var AllowedLocalReportMaterialGetV30DataMaterialListMaterialTypeEnumValues = []LocalReportMaterialGetV30DataMaterialListMaterialType{
	"CASURAL",
	"VIDEO",
}

// NewLocalReportMaterialGetV30DataMaterialListMaterialTypeFromValue returns a pointer to a valid LocalReportMaterialGetV30DataMaterialListMaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocalReportMaterialGetV30DataMaterialListMaterialTypeFromValue(v string) (*LocalReportMaterialGetV30DataMaterialListMaterialType, error) {
	ev := LocalReportMaterialGetV30DataMaterialListMaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocalReportMaterialGetV30DataMaterialListMaterialType: valid values are %v", v, AllowedLocalReportMaterialGetV30DataMaterialListMaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocalReportMaterialGetV30DataMaterialListMaterialType) IsValid() bool {
	for _, existing := range AllowedLocalReportMaterialGetV30DataMaterialListMaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to local_report_material_get_v3.0_data_material_list_material_type value
func (v LocalReportMaterialGetV30DataMaterialListMaterialType) Ptr() *LocalReportMaterialGetV30DataMaterialListMaterialType {
	return &v
}
