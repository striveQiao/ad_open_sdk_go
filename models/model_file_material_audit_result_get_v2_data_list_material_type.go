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

// FileMaterialAuditResultGetV2DataListMaterialType
type FileMaterialAuditResultGetV2DataListMaterialType string

// List of file_material_audit_result_get_v2_data_list_material_type
const (
	VIDEO_FileMaterialAuditResultGetV2DataListMaterialType   FileMaterialAuditResultGetV2DataListMaterialType = "VIDEO"
	PICTURE_FileMaterialAuditResultGetV2DataListMaterialType FileMaterialAuditResultGetV2DataListMaterialType = "PICTURE"
	TITLE_FileMaterialAuditResultGetV2DataListMaterialType   FileMaterialAuditResultGetV2DataListMaterialType = "TITLE"
)

// All allowed values of FileMaterialAuditResultGetV2DataListMaterialType enum
var AllowedFileMaterialAuditResultGetV2DataListMaterialTypeEnumValues = []FileMaterialAuditResultGetV2DataListMaterialType{
	"VIDEO",
	"PICTURE",
	"TITLE",
}

// NewFileMaterialAuditResultGetV2DataListMaterialTypeFromValue returns a pointer to a valid FileMaterialAuditResultGetV2DataListMaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileMaterialAuditResultGetV2DataListMaterialTypeFromValue(v string) (*FileMaterialAuditResultGetV2DataListMaterialType, error) {
	ev := FileMaterialAuditResultGetV2DataListMaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileMaterialAuditResultGetV2DataListMaterialType: valid values are %v", v, AllowedFileMaterialAuditResultGetV2DataListMaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileMaterialAuditResultGetV2DataListMaterialType) IsValid() bool {
	for _, existing := range AllowedFileMaterialAuditResultGetV2DataListMaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_material_audit_result_get_v2_data_list_material_type value
func (v FileMaterialAuditResultGetV2DataListMaterialType) Ptr() *FileMaterialAuditResultGetV2DataListMaterialType {
	return &v
}
