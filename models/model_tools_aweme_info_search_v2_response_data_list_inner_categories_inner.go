/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeInfoSearchV2ResponseDataListInnerCategoriesInner struct for ToolsAwemeInfoSearchV2ResponseDataListInnerCategoriesInner
type ToolsAwemeInfoSearchV2ResponseDataListInnerCategoriesInner struct {
	//
	Children []*ToolsAwemeInfoSearchV2ResponseDataListInnerCategoriesInnerChildrenInner `json:"children,omitempty"`
	//
	CoverNumStr *string `json:"cover_num_str,omitempty"`
	//
	FansNumStr *string `json:"fans_num_str,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	Value *string `json:"value,omitempty"`
}
