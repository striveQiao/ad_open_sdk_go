/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateCreateV2ResponseDataBricksInnerWechatApplet 微信小程序组件
type ToolsSiteTemplateCreateV2ResponseDataBricksInnerWechatApplet struct {
	// 微信小程序组件ID
	InstanceId *int64 `json:"instance_id,omitempty"`
	// 简介，长度不超过40；
	Introduction *string `json:"introduction,omitempty"`
	// 标签，个数不超过2，字数不超过5
	Items []string `json:"items,omitempty"`
	// logo链接
	Logo *string `json:"logo,omitempty"`
}
