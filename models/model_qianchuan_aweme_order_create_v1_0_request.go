/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderCreateV10Request struct for QianchuanAwemeOrderCreateV10Request
type QianchuanAwemeOrderCreateV10Request struct {
	//
	AdvertiserId int64                                        `json:"advertiser_id"`
	Audience     *QianchuanAwemeOrderCreateV10RequestAudience `json:"audience,omitempty"`
	// 推广对象的抖音id 1. 视频加热和视频加热直播间：传递视频作者的抖音uid 2. 直接加热直播间：传递主播的抖音uid
	AwemeId int64 `json:"aweme_id"`
	//
	AwemeItemId     *int64                                             `json:"aweme_item_id,omitempty"`
	DeliverySetting QianchuanAwemeOrderCreateV10RequestDeliverySetting `json:"delivery_setting"`
	MarketingGoal   QianchuanAwemeOrderCreateV10MarketingGoal          `json:"marketing_goal"`
}
