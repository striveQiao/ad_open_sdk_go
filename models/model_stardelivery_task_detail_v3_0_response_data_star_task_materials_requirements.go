/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskDetailV30ResponseDataStarTaskMaterialsRequirements 制作要求
type StardeliveryTaskDetailV30ResponseDataStarTaskMaterialsRequirements struct {
	// 补充说明，1-3000个字符
	AdditionalInformation *string `json:"additional_information,omitempty"`
	// 镜头要求，1-80个字符，例如产品实物出镜、近景特写、达人出镜、镜头时长、指定画面等 例如：需产品实物出镜，并给到带logo的产品近景特写，特写时长不低于3s，达人需出镜亲自安利
	OnCameraRequirement *string `json:"on_camera_requirement,omitempty"`
	//
	OtherRequirements []string `json:"other_requirements,omitempty"`
	// 示例视频id，最多可添加3个
	SampleMaterialIds []string `json:"sample_material_ids,omitempty"`
	// 示例视频url，最多可添加3个，需要输入抖音视频链接并确保该视频为在线状态 （如 https://www.douyin.com/video/xxx）
	SampleVideoUrls []string `json:"sample_video_urls,omitempty"`
	// 要@ 的抖音号，抖音用户个人主页抖音昵称下方可找到
	TitleMentionsAwemeId *string `json:"title_mentions_aweme_id,omitempty"`
	// 标题要求，1-80个字符，例如标题包含商品名称、商品性能/特点关键词等；若需提及账号或话题，请在其他要求参数中填写
	TitleRequirement *string `json:"title_requirement,omitempty"`
	//
	TitleSpecifiesTopicIds []int64 `json:"title_specifies_topic_ids,omitempty"`
	// 口播要求，1-80个字符，例如口播介绍指定文字、品牌、slogan、商品卖点，是否添加口播字幕等
	VoRequirement *string `json:"vo_requirement,omitempty"`
}
