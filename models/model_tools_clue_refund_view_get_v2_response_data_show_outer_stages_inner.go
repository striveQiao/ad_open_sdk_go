/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInner struct for ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInner
type ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInner struct {
	CallCountDetail *ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInnerCallCountDetail `json:"call_count_detail,omitempty"`
	// 返款标签 (参与返款、支持赔付)
	RefundCallList []*ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInnerRefundCallListInner `json:"refund_call_list,omitempty"`
	// 展示上面的小阶段，比如人工拨打时，第一个ShowOuterStage包含了两个小阶段（0-12小时，0-24小时）
	ShowInnerStages []*ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInnerShowInnerStagesInner `json:"show_inner_stages,omitempty"`
	// 线索id
	Title *string `json:"title,omitempty"`
}
