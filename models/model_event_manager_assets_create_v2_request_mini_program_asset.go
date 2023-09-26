/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerAssetsCreateV2RequestMiniProgramAsset 小程序信息
type EventManagerAssetsCreateV2RequestMiniProgramAsset struct {
	// 小程序instance_id
	InstanceId int64 `json:"instance_id"`
	// 小程序appId
	MiniProgramId string `json:"mini_program_id"`
	// 小程序的名称
	MiniProgramName string                                                    `json:"mini_program_name"`
	MiniProgramType EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType `json:"mini_program_type"`
}
