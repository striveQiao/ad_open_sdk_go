/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetInfoV2ResponseData
type StarOrderGetInfoV2ResponseData struct {
	// 任务列表
	OrderList []*StarOrderGetInfoV2ResponseDataOrderListInner `json:"order_list,omitempty"`
}
