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

// ToolsRtaSetScopeV2AccountType
type ToolsRtaSetScopeV2AccountType string

// List of tools_rta_set_scope_v2_account_type
const (
	AD_ToolsRtaSetScopeV2AccountType   ToolsRtaSetScopeV2AccountType = "AD"
	STAR_ToolsRtaSetScopeV2AccountType ToolsRtaSetScopeV2AccountType = "STAR"
)

// All allowed values of ToolsRtaSetScopeV2AccountType enum
var AllowedToolsRtaSetScopeV2AccountTypeEnumValues = []ToolsRtaSetScopeV2AccountType{
	"AD",
	"STAR",
}

// NewToolsRtaSetScopeV2AccountTypeFromValue returns a pointer to a valid ToolsRtaSetScopeV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRtaSetScopeV2AccountTypeFromValue(v string) (*ToolsRtaSetScopeV2AccountType, error) {
	ev := ToolsRtaSetScopeV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRtaSetScopeV2AccountType: valid values are %v", v, AllowedToolsRtaSetScopeV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRtaSetScopeV2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsRtaSetScopeV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rta_set_scope_v2_account_type value
func (v ToolsRtaSetScopeV2AccountType) Ptr() *ToolsRtaSetScopeV2AccountType {
	return &v
}
