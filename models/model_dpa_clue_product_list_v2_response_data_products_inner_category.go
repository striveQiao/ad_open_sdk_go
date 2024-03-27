/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaClueProductListV2ResponseDataProductsInnerCategory 商品类目信息
type DpaClueProductListV2ResponseDataProductsInnerCategory struct {
	// 一级类目ID
	FirstCategoryId *int64 `json:"first_category_id,omitempty"`
	// 一级类目名称
	FirstCategoryName *string `json:"first_category_name,omitempty"`
	// 四级类目ID
	FourthCategoryId *int64 `json:"fourth_category_id,omitempty"`
	// 四级类目名称
	FourthCategoryName *string `json:"fourth_category_name,omitempty"`
	// 二级类目ID
	SecondCategoryId *int64 `json:"second_category_id,omitempty"`
	// 二级类目名称
	SecondCategoryName *string `json:"second_category_name,omitempty"`
	// 三级类目ID
	ThirdCategoryId *int64 `json:"third_category_id,omitempty"`
	// 三级类目名称
	ThirdCategoryName *string `json:"third_category_name,omitempty"`
}
