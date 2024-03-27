/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmGetChallengeV2ResponseDataTaskInfoChallengeInfo
type StarDemandOmGetChallengeV2ResponseDataTaskInfoChallengeInfo struct {
	// 最长分账周期
	AccountDivideDay *int64 `json:"account_divide_day,omitempty"`
	// 组件标题
	AnchorTitle *string `json:"anchor_title,omitempty"`
	//
	AuthorList []*StarDemandOmGetChallengeV2ResponseDataTaskInfoChallengeInfoAuthorListInner `json:"author_list,omitempty"`
	// 达人选择
	AuthorScope *int64 `json:"author_scope,omitempty"`
	// 达人侧任务名称
	AuthorTaskName string `json:"author_task_name"`
	//
	Budget int64 `json:"budget"`
	// 分佣比例
	CommissionRate *int64 `json:"commission_rate,omitempty"`
	// 结算方式
	CommissionType *int64 `json:"commission_type,omitempty"`
	// 任务介绍
	DemandDesc *string `json:"demand_desc,omitempty"`
	// 任务截止时间
	EndTime int64 `json:"end_time"`
	//
	MaxUploadItemCnt int64 `json:"max_upload_item_cnt"`
	// 小程序ID
	MicroAppId *string `json:"micro_app_id,omitempty"`
	// 任务状态
	OmTaskStatus *int64 `json:"om_task_status,omitempty"`
	// 任务标签
	OmTaskTag              []string                                                                          `json:"om_task_tag,omitempty"`
	ParticipateAuthorRange StarDemandOmGetChallengeV2ResponseDataTaskInfoChallengeInfoParticipateAuthorRange `json:"participate_author_range"`
	//
	SampleDelivery int64 `json:"sample_delivery"`
	// 示例视频
	SampleVideo    []int64                                                                   `json:"sample_video,omitempty"`
	SettlementInfo StarDemandOmGetChallengeV2ResponseDataTaskInfoChallengeInfoSettlementInfo `json:"settlement_info"`
	// 小程序落地页地址
	StartPage *string `json:"start_page,omitempty"`
	// 投稿开始时间
	StartTime int64 `json:"start_time"`
	// 任务头图
	TaskHeadImage *string `json:"task_head_image,omitempty"`
	// 任务图标
	TaskIcon string `json:"task_icon"`
}
