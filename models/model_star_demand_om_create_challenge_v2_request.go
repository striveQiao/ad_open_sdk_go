/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmCreateChallengeV2Request struct for StarDemandOmCreateChallengeV2Request
type StarDemandOmCreateChallengeV2Request struct {
	ChallengeInfo StarDemandOmCreateChallengeV2RequestChallengeInfo `json:"challenge_info"`
	DemandInfo    StarDemandOmCreateChallengeV2RequestDemandInfo    `json:"demand_info"`
	//
	DeveloperId *int64 `json:"developer_id,omitempty"`
	// 客户星图ID
	StarId int64 `json:"star_id"`
}
