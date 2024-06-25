/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoBrandNameListInner struct for BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoBrandNameListInner
type BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoBrandNameListInner struct {
	// 所属一级行业ID
	FirstCategoryIds []int64 `json:"first_category_ids,omitempty"`
	// 品牌名称
	Label *string `json:"label,omitempty"`
	// 品牌ID
	Value *int64 `json:"value,omitempty"`
}
