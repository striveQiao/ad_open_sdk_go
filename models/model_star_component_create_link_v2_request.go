/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarComponentCreateLinkV2Request struct for StarComponentCreateLinkV2Request
type StarComponentCreateLinkV2Request struct {
	LinkComponentInfo *StarComponentCreateLinkV2RequestLinkComponentInfo `json:"link_component_info,omitempty"`
	// 星图客户ID
	StarId int64 `json:"star_id"`
}
