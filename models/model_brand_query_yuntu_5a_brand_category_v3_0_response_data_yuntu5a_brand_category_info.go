/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfo
type BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfo struct {
	// 品牌分类名称
	BrandNameList []*BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoBrandNameListInner `json:"brand_name_list,omitempty"`
	YuntuCategory *BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoYuntuCategory        `json:"yuntu_category,omitempty"`
}
