/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAnalyseListV10ResponseData
type QianchuanProductAnalyseListV10ResponseData struct {
	PageInfo *QianchuanProductAnalyseListV10ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	ProductList []*QianchuanProductAnalyseListV10ResponseDataProductListInner `json:"product_list,omitempty"`
}
