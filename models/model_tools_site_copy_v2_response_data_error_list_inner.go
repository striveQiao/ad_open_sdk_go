/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteCopyV2ResponseDataErrorListInner struct for ToolsSiteCopyV2ResponseDataErrorListInner
type ToolsSiteCopyV2ResponseDataErrorListInner struct {
	// 当对origin_site_id的操作失败时，返回详细原因，成功无
	ErrorReason *string `json:"error_reason,omitempty"`
	// 返回复制失败后的原站点id
	OriginSiteId *string `json:"origin_site_id,omitempty"`
}
