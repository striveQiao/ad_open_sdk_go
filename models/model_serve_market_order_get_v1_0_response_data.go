/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ServeMarketOrderGetV10ResponseData
type ServeMarketOrderGetV10ResponseData struct {
	//
	OrderList []*ServeMarketOrderGetV10ResponseDataOrderListInner `json:"order_list,omitempty"`
	PageInfo  *ServeMarketOrderGetV10ResponseDataPageInfo         `json:"page_info,omitempty"`
}
