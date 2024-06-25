/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoYuntuCategory 行业三级分类集合
type BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoYuntuCategory struct {
	// 一级行业分类集合
	FirstCategoryData []*BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoYuntuCategoryFirstCategoryDataInner `json:"first_category_data,omitempty"`
	// 二级行业分类集合
	SecondCategoryData *map[string][]*BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoYuntuCategorySecondCategoryDataValueInner `json:"second_category_data,omitempty"`
	// 三级级行业分类集合
	ThirdCategoryData *map[string][]*BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoYuntuCategoryThirdCategoryDataValueInner `json:"third_category_data,omitempty"`
}
