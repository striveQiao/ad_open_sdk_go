/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentStickOnTopV30Request struct for ToolsCommentStickOnTopV30Request
type ToolsCommentStickOnTopV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 评论id
	CommentId int64                              `json:"comment_id"`
	StickType ToolsCommentStickOnTopV30StickType `json:"stick_type"`
}
