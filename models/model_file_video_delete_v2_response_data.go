/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoDeleteV2ResponseData
type FileVideoDeleteV2ResponseData struct {
	// 操作失败的视频ID列表，不在此列表内的素材表示删除成功
	FailMaterialIds []int64 `json:"fail_material_ids,omitempty"`
	// 操作失败的素材ID列表，不在此列表内的素材表示删除成功
	FailVideoIds []string `json:"fail_video_ids,omitempty"`
}
