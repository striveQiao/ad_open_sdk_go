/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeUpdateV2RequestAdDataMiniProgramInfo
type CreativeProceduralCreativeUpdateV2RequestAdDataMiniProgramInfo struct {
	//
	AppId string `json:"app_id"`
	//
	Params *string `json:"params,omitempty"`
	//
	StartPath *string                                                     `json:"start_path,omitempty"`
	Type      CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType `json:"type"`
	//
	Url *string `json:"url,omitempty"`
}
