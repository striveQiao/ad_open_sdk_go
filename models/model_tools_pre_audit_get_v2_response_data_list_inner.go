/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPreAuditGetV2ResponseDataListInner struct for ToolsPreAuditGetV2ResponseDataListInner
type ToolsPreAuditGetV2ResponseDataListInner struct {
	// 物料内容
	Content      *string                                 `json:"content,omitempty"`
	MaterialType *ToolsPreAuditGetV2DataListMaterialType `json:"material_type,omitempty"`
	// 预审物料 - 文本物料：用户上传 - 图片物料：用户上传图片链接或图片ID - 视频物料：用户上传的视频云ID
	PreAuditId *int64 `json:"pre_audit_id,omitempty"`
	// 拒绝理由
	RejectReason *string                           `json:"reject_reason,omitempty"`
	Status       *ToolsPreAuditGetV2DataListStatus `json:"status,omitempty"`
}
