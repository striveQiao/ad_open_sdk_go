/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeProductAvailableGetV10ResponseData
type QianchuanAwemeProductAvailableGetV10ResponseData struct {
	PageInfo *QianchuanAwemeProductAvailableGetV10ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	ProductList []*QianchuanAwemeProductAvailableGetV10ResponseDataProductListInner `json:"product_list"`
}
