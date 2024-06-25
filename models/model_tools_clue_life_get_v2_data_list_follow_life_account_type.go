/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsClueLifeGetV2DataListFollowLifeAccountType
type ToolsClueLifeGetV2DataListFollowLifeAccountType string

// List of tools_clue_life_get_v2_data_list_follow_life_account_type
const (
	HEAD_ToolsClueLifeGetV2DataListFollowLifeAccountType   ToolsClueLifeGetV2DataListFollowLifeAccountType = "HEAD"
	REGION_ToolsClueLifeGetV2DataListFollowLifeAccountType ToolsClueLifeGetV2DataListFollowLifeAccountType = "REGION"
	SINGLE_ToolsClueLifeGetV2DataListFollowLifeAccountType ToolsClueLifeGetV2DataListFollowLifeAccountType = "SINGLE"
)

// All allowed values of ToolsClueLifeGetV2DataListFollowLifeAccountType enum
var AllowedToolsClueLifeGetV2DataListFollowLifeAccountTypeEnumValues = []ToolsClueLifeGetV2DataListFollowLifeAccountType{
	"HEAD",
	"REGION",
	"SINGLE",
}

// NewToolsClueLifeGetV2DataListFollowLifeAccountTypeFromValue returns a pointer to a valid ToolsClueLifeGetV2DataListFollowLifeAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueLifeGetV2DataListFollowLifeAccountTypeFromValue(v string) (*ToolsClueLifeGetV2DataListFollowLifeAccountType, error) {
	ev := ToolsClueLifeGetV2DataListFollowLifeAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueLifeGetV2DataListFollowLifeAccountType: valid values are %v", v, AllowedToolsClueLifeGetV2DataListFollowLifeAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueLifeGetV2DataListFollowLifeAccountType) IsValid() bool {
	for _, existing := range AllowedToolsClueLifeGetV2DataListFollowLifeAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_life_get_v2_data_list_follow_life_account_type value
func (v ToolsClueLifeGetV2DataListFollowLifeAccountType) Ptr() *ToolsClueLifeGetV2DataListFollowLifeAccountType {
	return &v
}
