/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdminInfoV2ResponseData
type ToolsAdminInfoV2ResponseData struct {
	// 行政层级信息
	Districts []*ToolsAdminInfoV2ResponseDataDistrictsInner `json:"districts,omitempty"`
	// 行政信息版本号
	Version *string `json:"version,omitempty"`
}
