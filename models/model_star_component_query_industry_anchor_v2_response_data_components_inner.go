/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarComponentQueryIndustryAnchorV2ResponseDataComponentsInner struct for StarComponentQueryIndustryAnchorV2ResponseDataComponentsInner
type StarComponentQueryIndustryAnchorV2ResponseDataComponentsInner struct {
	// 组件名称
	AnchorName *string `json:"anchor_name,omitempty"`
	// 组件状态
	AnchorStatus *int64 `json:"anchor_status,omitempty"`
	// 主标题
	AnchorTitle *string `json:"anchor_title,omitempty"`
	// 行业锚点类型
	AnchorType *int64 `json:"anchor_type,omitempty"`
	// 审核驳回原因（如果组件被审核驳回）
	AuditRejectReason *string `json:"audit_reject_reason,omitempty"`
	// 行业组件id
	IndustryAnchorId *int64 `json:"industry_anchor_id,omitempty"`
	// 六分屏预览链接
	PreviewUrl *string `json:"preview_url,omitempty"`
}
