/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeCreateV2RequestAdDataMiniProgramInfo
type CreativeProceduralCreativeCreateV2RequestAdDataMiniProgramInfo struct {
	//
	AppId string `json:"app_id"`
	//
	Params *string `json:"params,omitempty"`
	//
	StartPath *string                                                     `json:"start_path,omitempty"`
	Type      CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType `json:"type"`
	//
	Url *string `json:"url,omitempty"`
}
