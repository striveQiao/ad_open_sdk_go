/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionRejectReasonGetV30ResponseDataListInnerMaterialRejectInner struct for PromotionRejectReasonGetV30ResponseDataListInnerMaterialRejectInner
type PromotionRejectReasonGetV30ResponseDataListInnerMaterialRejectInner struct {
	AuditPlatform *PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform `json:"audit_platform,omitempty"`
	//
	Item *string `json:"item,omitempty"`
	// 拒绝建议
	RejectReason []string `json:"reject_reason,omitempty"`
	// 审核建议
	Suggestion []string                                               `json:"suggestion,omitempty"`
	Type       *PromotionRejectReasonGetV30DataListMaterialRejectType `json:"type,omitempty"`
}
