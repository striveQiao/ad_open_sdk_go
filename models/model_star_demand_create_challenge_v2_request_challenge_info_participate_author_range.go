/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateChallengeV2RequestChallengeInfoParticipateAuthorRange 投稿达人范围
type StarDemandCreateChallengeV2RequestChallengeInfoParticipateAuthorRange struct {
	// 达人ID列表 达人的抖音uid（仅自定义结算可用）
	AuthorUids       []int64                                                                                `json:"author_uids,omitempty"`
	AuthorWatcherTag *StarDemandCreateChallengeV2RequestChallengeInfoParticipateAuthorRangeAuthorWatcherTag `json:"author_watcher_tag,omitempty"`
	// 达人内容类型标签列表 「投稿任务创建数据字典」返回的content_tags中标签的tag_value值构成的列表
	ContentTags []int64 `json:"content_tags,omitempty"`
	// 达人性别 “男”，“女”
	Gender []string `json:"gender,omitempty"`
	// 达人粉丝量范围最大值 不得小于min_follower(如有)
	MaxFollower *int64 `json:"max_follower,omitempty"`
	// 达人粉丝量范围最小值 非负整数
	MinFollower *int64 `json:"min_follower,omitempty"`
	// 达人省份 「投稿任务创建数据字典」返回的provinces中的项 形如“北京市”、“江苏省”
	Provinces []string `json:"provinces,omitempty"`
}
