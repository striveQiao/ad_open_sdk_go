/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementIndustryInfoListV2AccountType
type ToolsAppManagementIndustryInfoListV2AccountType string

// List of tools_app_management_industry_info_list_v2_account_type
const (
	AD_ToolsAppManagementIndustryInfoListV2AccountType ToolsAppManagementIndustryInfoListV2AccountType = "AD"
	BP_ToolsAppManagementIndustryInfoListV2AccountType ToolsAppManagementIndustryInfoListV2AccountType = "BP"
)

// All allowed values of ToolsAppManagementIndustryInfoListV2AccountType enum
var AllowedToolsAppManagementIndustryInfoListV2AccountTypeEnumValues = []ToolsAppManagementIndustryInfoListV2AccountType{
	"AD",
	"BP",
}

// NewToolsAppManagementIndustryInfoListV2AccountTypeFromValue returns a pointer to a valid ToolsAppManagementIndustryInfoListV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementIndustryInfoListV2AccountTypeFromValue(v string) (*ToolsAppManagementIndustryInfoListV2AccountType, error) {
	ev := ToolsAppManagementIndustryInfoListV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementIndustryInfoListV2AccountType: valid values are %v", v, AllowedToolsAppManagementIndustryInfoListV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementIndustryInfoListV2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementIndustryInfoListV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_industry_info_list_v2_account_type value
func (v ToolsAppManagementIndustryInfoListV2AccountType) Ptr() *ToolsAppManagementIndustryInfoListV2AccountType {
	return &v
}
