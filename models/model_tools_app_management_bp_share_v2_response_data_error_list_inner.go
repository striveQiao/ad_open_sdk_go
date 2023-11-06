/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementBpShareV2ResponseDataErrorListInner struct for ToolsAppManagementBpShareV2ResponseDataErrorListInner
type ToolsAppManagementBpShareV2ResponseDataErrorListInner struct {
	AccountInfo         *ToolsAppManagementBpShareV2ResponseDataErrorListInnerAccountInfo         `json:"account_info,omitempty"`
	AllAccount          *ToolsAppManagementBpShareV2ResponseDataErrorListInnerAllAccount          `json:"all_account,omitempty"`
	AllAccountByCompany *ToolsAppManagementBpShareV2ResponseDataErrorListInnerAllAccountByCompany `json:"all_account_by_company,omitempty"`
	//
	ErrorMessage *string                                            `json:"error_message,omitempty"`
	ShareMode    *ToolsAppManagementBpShareV2DataErrorListShareMode `json:"share_mode,omitempty"`
}
