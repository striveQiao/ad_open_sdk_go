/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CdpBrandGetV30ResponseDataBrandCategoryInfoInner struct for CdpBrandGetV30ResponseDataBrandCategoryInfoInner
type CdpBrandGetV30ResponseDataBrandCategoryInfoInner struct {
	//
	Children []*CdpBrandGetV30ResponseDataBrandCategoryInfoInnerChildrenInner `json:"children,omitempty"`
	// 一级类别id: yuntu_category_id
	Id *string `json:"id,omitempty"`
	// 一级类别标签
	Lable *string `json:"lable,omitempty"`
}
