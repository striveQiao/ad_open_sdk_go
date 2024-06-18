/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderDetailV2ResponseData
type StarOrderDetailV2ResponseData struct {
	// 订单审核状态 未审核 = 0 脚本审核失败 = 1 脚本审核通过 = 2 视频审核失败 = 3 视频审核通过 = 4 组件审核拒绝 = 5 组件审核通过 = 6 订单跳过审核 = 7
	AuditStatus *int64 `json:"audit_status,omitempty"`
	// 达人取消原因
	AuthorCancelReason *string `json:"author_cancel_reason,omitempty"`
	// 达人取消状态 0 = 未发起 1 = 已发起
	AuthorCancelStatus *int64                                   `json:"author_cancel_status,omitempty"`
	AuthorInfo         *StarOrderDetailV2ResponseDataAuthorInfo `json:"author_info,omitempty"`
	// 组件信息
	ComponentInfo []*StarOrderDetailV2ResponseDataComponentInfoInner `json:"component_info,omitempty"`
	// 任务创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
	// 任务完结时间
	FinishTime *int64 `json:"finish_time,omitempty"`
	// 任务ID
	OrderId     *int64                                    `json:"order_id,omitempty"`
	PaymentInfo *StarOrderDetailV2ResponseDataPaymentInfo `json:"payment_info,omitempty"`
	// 脚本列表
	ScriptList []*StarOrderDetailV2ResponseDataScriptListInner `json:"script_list,omitempty"`
	// 任务状态 待接收 = -1 达人已接单 = 1 已关闭 = 2 已完成 = 3 已取消 = 4 待付尾款 = 10 脚本已上传 = 41 脚本已拒绝 = 42 脚本已确认 = 43 脚本已跳过 = 44 视频已上传 = 51 视频已拒绝 = 52 视频已确认 = 53 视频已发布 = 54
	Status *int64 `json:"status,omitempty"`
	// 视频列表
	VideoList []*StarOrderDetailV2ResponseDataVideoListInner `json:"video_list,omitempty"`
}
