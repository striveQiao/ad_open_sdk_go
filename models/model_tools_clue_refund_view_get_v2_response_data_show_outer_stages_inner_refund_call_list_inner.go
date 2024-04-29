/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInnerRefundCallListInner struct for ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInnerRefundCallListInner
type ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInnerRefundCallListInner struct {
	// 通话状态
	BillStateText *string `json:"bill_state_text,omitempty"`
	// 通话id, 和话单一一对应
	CallId *string `json:"call_id,omitempty"`
	// UI颜色，飞鱼SDK前端所需字段，传入即可
	Color *string `json:"color,omitempty"`
	// 不满足原因，, 仅在不满足要求时展示<br>例如：\"原因： 未打够35s\"
	Reason *string `json:"reason,omitempty"`
	// 0,1,2
	RefundCheckResult *int64 `json:"refund_check_result,omitempty"`
	// 满足要求 | 不满足要求 | 已接通
	RefundCheckResultText *string `json:"refund_check_result_text,omitempty"`
}
