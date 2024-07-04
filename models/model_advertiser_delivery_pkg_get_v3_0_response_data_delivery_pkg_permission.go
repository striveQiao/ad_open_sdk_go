/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgPermission 权限信息
type AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgPermission struct {
	// 是否支持编辑
	CanDelete bool `json:"can_delete"`
	// 是否支持编辑
	CanEdit bool `json:"can_edit"`
	// 不支持删除的原因
	CantDeleteReason *string `json:"cant_delete_reason,omitempty"`
	// 不支持编辑的原因
	CantEditReason *string `json:"cant_edit_reason,omitempty"`
}
