/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppUpdateV30RequestAppPageInner struct for ToolsMicroAppUpdateV30RequestAppPageInner
type ToolsMicroAppUpdateV30RequestAppPageInner struct {
	//
	Id *int64 `json:"id,omitempty"`
	//
	Link        string                                   `json:"link"`
	OperateType ToolsMicroAppUpdateV30AppPageOperateType `json:"operate_type"`
	//
	Remark string `json:"remark"`
	//
	StartPage *string `json:"start_page,omitempty"`
	//
	StartParam *string `json:"start_param,omitempty"`
}
