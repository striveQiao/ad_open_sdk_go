/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeItemsDataV2ResponseDataDataListInner struct for StarChallengeItemsDataV2ResponseDataDataListInner
type StarChallengeItemsDataV2ResponseDataDataListInner struct {
	// Android激活数
	AndroidActivateCnt *int64 `json:"android_activate_cnt,omitempty"`
	// 作者昵称
	AuthorNickname *string `json:"author_nickname,omitempty"`
	// 评论量
	CommentCnt *int64 `json:"comment_cnt,omitempty"`
	// 组件点击数
	ComponentClickCnt *int64 `json:"component_click_cnt,omitempty"`
	// iOS激活数
	IosActivateCnt *int64 `json:"ios_activate_cnt,omitempty"`
	// 作品ID
	ItemId *int64 `json:"item_id,omitempty"`
	// 点赞量
	LikeCnt *int64 `json:"like_cnt,omitempty"`
	// 作品链接
	Link *string `json:"link,omitempty"`
	// 播放量
	PlayVv *int64 `json:"play_vv,omitempty"`
	// 促活转化数
	PromoteCnt *int64 `json:"promote_cnt,omitempty"`
	// 发布时间
	ReleaseTime *int64 `json:"release_time,omitempty"`
	// 1：审核中 2：相关 3：不相关
	RelevanceAuditResult *int64 `json:"relevance_audit_result,omitempty"`
	// 获奖金额（分）
	RewardAmount *int64 `json:"reward_amount,omitempty"`
	// 获奖等级
	RewardLevel *int64 `json:"reward_level,omitempty"`
	// 分享量
	ShareCnt *int64 `json:"share_cnt,omitempty"`
	// 作品标题
	Title *string `json:"title,omitempty"`
	// 作者uid
	Uid *int64 `json:"uid,omitempty"`
	// 有效点赞量
	ValidLikeCnt *int64 `json:"valid_like_cnt,omitempty"`
	// 有效播放量
	ValidPlayVv *int64 `json:"valid_play_vv,omitempty"`
}
