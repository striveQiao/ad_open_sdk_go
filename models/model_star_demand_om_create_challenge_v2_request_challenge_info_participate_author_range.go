/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmCreateChallengeV2RequestChallengeInfoParticipateAuthorRange
type StarDemandOmCreateChallengeV2RequestChallengeInfoParticipateAuthorRange struct {
	// 达人抖音号,此字段非必填，仅在专属任务challenge_info.author_scope=2时需要传
	UniqueIds []int64 `json:"unique_ids,omitempty"`
}
