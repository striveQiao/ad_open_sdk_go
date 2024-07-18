/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmGetChallengeV2ResponseDataTaskInfoChallengeInfoAuthorListInner struct for StarDemandOmGetChallengeV2ResponseDataTaskInfoChallengeInfoAuthorListInner
type StarDemandOmGetChallengeV2ResponseDataTaskInfoChallengeInfoAuthorListInner struct {
	// - 取消方：1-开发者，2-达人，3-平台
	CancelRole *int64 `json:"cancel_role,omitempty"`
	// - 合作状态，1-已发布视频 2-待发布视频 3-合作已取消
	ParticipateStatus *int64 `json:"participate_status,omitempty"`
	// - 达人抖音号
	UniqueId *string `json:"unique_id,omitempty"`
}
