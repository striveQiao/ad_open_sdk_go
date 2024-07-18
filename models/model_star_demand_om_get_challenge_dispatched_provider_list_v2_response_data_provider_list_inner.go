/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmGetChallengeDispatchedProviderListV2ResponseDataProviderListInner struct for StarDemandOmGetChallengeDispatchedProviderListV2ResponseDataProviderListInner
type StarDemandOmGetChallengeDispatchedProviderListV2ResponseDataProviderListInner struct {
	// 是否可以和服务商取消合作
	CanCancelCooperate *bool `json:"can_cancel_cooperate,omitempty"`
	// 服务商是否接单
	HasProviderReceived *bool `json:"has_provider_received,omitempty"`
	//
	ProviderId *int64 `json:"provider_id,omitempty"`
}
