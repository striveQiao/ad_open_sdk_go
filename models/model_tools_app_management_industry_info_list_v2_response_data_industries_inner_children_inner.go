/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementIndustryInfoListV2ResponseDataIndustriesInnerChildrenInner struct for ToolsAppManagementIndustryInfoListV2ResponseDataIndustriesInnerChildrenInner
type ToolsAppManagementIndustryInfoListV2ResponseDataIndustriesInnerChildrenInner struct {
	// 子分类
	Children []*ToolsAppManagementIndustryInfoListV2ResponseDataIndustriesInnerChildrenInnerChildrenInner `json:"children,omitempty"`
	// 分类id
	Id *string `json:"id,omitempty"`
	// 级别
	Level *int64 `json:"level,omitempty"`
	// 分类名称
	Name *string `json:"name,omitempty"`
	// 题材标签列表
	ThemeTags []*ToolsAppManagementIndustryInfoListV2ResponseDataIndustriesInnerChildrenInnerThemeTagsInner `json:"theme_tags,omitempty"`
}
