/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRefundViewGetV2ResponseDataShowTimeDetails 开始/结束/变红的时间戳
type ToolsClueRefundViewGetV2ResponseDataShowTimeDetails struct {
	// 开始时间
	BeginTime *int64 `json:"begin_time,omitempty"`
	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`
	// 开始展示红色倒计时的时间，飞鱼SDK前端所需字段，传入即可
	RedLineTime *int64 `json:"red_line_time,omitempty"`
}
