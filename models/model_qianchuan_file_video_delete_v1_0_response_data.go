/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFileVideoDeleteV10ResponseData
type QianchuanFileVideoDeleteV10ResponseData struct {
	// 操作失败的video_id列表，不在此列表内的素材表示删除成功
	FailVideoIds []string `json:"fail_video_ids"`
}
