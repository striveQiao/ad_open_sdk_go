/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceActionV2
type AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceActionV2 struct {
	//
	ActionCategories *string `json:"action_categories,omitempty"`
	//
	ActionDays []int64 `json:"action_days,omitempty"`
	//
	ActionScenes []*AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes `json:"action_scenes,omitempty"`
	//
	VideoActions *string `json:"video_actions,omitempty"`
}
