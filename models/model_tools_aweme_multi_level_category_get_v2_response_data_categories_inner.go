/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeMultiLevelCategoryGetV2ResponseDataCategoriesInner struct for ToolsAwemeMultiLevelCategoryGetV2ResponseDataCategoriesInner
type ToolsAwemeMultiLevelCategoryGetV2ResponseDataCategoriesInner struct {
	//
	Children []*ToolsAwemeMultiLevelCategoryGetV2ResponseDataCategoriesInnerChildrenInner `json:"children,omitempty"`
	//
	CoverNumStr *string `json:"cover_num_str,omitempty"`
	//
	FansNumStr *string `json:"fans_num_str,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	Value *string `json:"value,omitempty"`
}
