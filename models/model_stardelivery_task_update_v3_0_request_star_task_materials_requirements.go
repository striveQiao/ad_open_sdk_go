/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskUpdateV30RequestStarTaskMaterialsRequirements
type StardeliveryTaskUpdateV30RequestStarTaskMaterialsRequirements struct {
	//
	AdditionalInformation *string `json:"additional_information,omitempty"`
	//
	OnCameraRequirement string `json:"on_camera_requirement"`
	//
	OtherRequirements []string `json:"other_requirements,omitempty"`
	// 参考素材 uri
	SampleMaterialIds []string `json:"sample_material_ids"`
	// 示例视频
	SampleVideoUrls []string `json:"sample_video_urls,omitempty"`
	//
	TitleMentionsAwemeId *string `json:"title_mentions_aweme_id,omitempty"`
	//
	TitleRequirement string `json:"title_requirement"`
	//
	TitleSpecifiesTopicIds []int64 `json:"title_specifies_topic_ids,omitempty"`
	//
	VoRequirement string `json:"vo_requirement"`
}
