/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileMaterialAttributesListV2AccountType
type FileMaterialAttributesListV2AccountType string

// List of file_material_attributes_list_v2_account_type
const (
	BP_FileMaterialAttributesListV2AccountType        FileMaterialAttributesListV2AccountType = "BP"
	AGENT_FileMaterialAttributesListV2AccountType     FileMaterialAttributesListV2AccountType = "AGENT"
	AD_FileMaterialAttributesListV2AccountType        FileMaterialAttributesListV2AccountType = "AD"
	QIANCHUAN_FileMaterialAttributesListV2AccountType FileMaterialAttributesListV2AccountType = "QIANCHUAN"
)

// All allowed values of FileMaterialAttributesListV2AccountType enum
var AllowedFileMaterialAttributesListV2AccountTypeEnumValues = []FileMaterialAttributesListV2AccountType{
	"BP",
	"AGENT",
	"AD",
	"QIANCHUAN",
}

// NewFileMaterialAttributesListV2AccountTypeFromValue returns a pointer to a valid FileMaterialAttributesListV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileMaterialAttributesListV2AccountTypeFromValue(v string) (*FileMaterialAttributesListV2AccountType, error) {
	ev := FileMaterialAttributesListV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileMaterialAttributesListV2AccountType: valid values are %v", v, AllowedFileMaterialAttributesListV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileMaterialAttributesListV2AccountType) IsValid() bool {
	for _, existing := range AllowedFileMaterialAttributesListV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_material_attributes_list_v2_account_type value
func (v FileMaterialAttributesListV2AccountType) Ptr() *FileMaterialAttributesListV2AccountType {
	return &v
}
