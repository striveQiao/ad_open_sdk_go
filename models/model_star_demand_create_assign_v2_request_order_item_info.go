/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateAssignV2RequestOrderItemInfo 任务供给侧相关信息
type StarDemandCreateAssignV2RequestOrderItemInfo struct {
	// order_type
	AuthorList []*StarDemandCreateAssignV2RequestOrderItemInfoAuthorListInner `json:"author_list"`
	// 任务付款类型 (1)：全款； (2)：预付
	OrderType int64 `json:"order_type"`
}
