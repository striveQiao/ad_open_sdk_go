/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentListGetV10ResponseDataCommentListInner struct for EnterpriseCommentListGetV10ResponseDataCommentListInner
type EnterpriseCommentListGetV10ResponseDataCommentListInner struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	CampaignId *int64 `json:"campaign_id,omitempty"`
	//
	CommentAwemeName *string `json:"comment_aweme_name,omitempty"`
	//
	CommentId *int64 `json:"comment_id,omitempty"`
	//
	CommentImage []string `json:"comment_image,omitempty"`
	//
	CommentOpenId *string `json:"comment_open_id,omitempty"`
	//
	CommentText *string `json:"comment_text,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	CreativeId *int64 `json:"creative_id,omitempty"`
	//
	DiggCount *int64 `json:"digg_count,omitempty"`
	//
	IsBannedUser *int64 `json:"is_banned_user,omitempty"`
	//
	IsStick *int64 `json:"is_stick,omitempty"`
	//
	ItemAwemeName *string `json:"item_aweme_name,omitempty"`
	//
	ItemId *int64 `json:"item_id,omitempty"`
	//
	ItemOpenId *string                                          `json:"item_open_id,omitempty"`
	Level      *EnterpriseCommentListGetV10DataCommentListLevel `json:"level,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	RepliedCommentId *int64 `json:"replied_comment_id,omitempty"`
	//
	ReplyCount *int64                                            `json:"reply_count,omitempty"`
	Source     *EnterpriseCommentListGetV10DataCommentListSource `json:"source,omitempty"`
}
