/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentGetV30ResponseData
type ToolsCommentGetV30ResponseData struct {
	// 评论列表
	CommentList []*ToolsCommentGetV30ResponseDataCommentListInner `json:"comment_list"`
	// 负评数
	NegativeVolume *int64 `json:"negative_volume,omitempty"`
	// 负评率
	NegativeVolumeRate *float64                               `json:"negative_volume_rate,omitempty"`
	PageInfo           ToolsCommentGetV30ResponseDataPageInfo `json:"page_info"`
}
