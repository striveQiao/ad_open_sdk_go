/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerVideoOnlineVideo 在线视频内容
type ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerVideoOnlineVideo struct {
	// 视频原始地址，例如https://v.youku.com/v_show/id_xxx.html，当前仅支持优酷视频，当`online_video`不为空时，有值
	OriginUrl *string `json:"origin_url,omitempty"`
	// 视频封面图链接，当`online_video`不为空时，有值
	PosterUrl *string `json:"poster_url,omitempty"`
}
