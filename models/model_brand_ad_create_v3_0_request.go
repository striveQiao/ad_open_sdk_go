/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdCreateV30Request struct for BrandAdCreateV30Request
type BrandAdCreateV30Request struct {
	AdForm BrandAdCreateV30AdForm `json:"ad_form"`
	// 广告主ID
	AdvertiserId int64                                `json:"advertiser_id"`
	AppOrigin    BrandAdCreateV30AppOrigin            `json:"app_origin"`
	AudienceInfo *BrandAdCreateV30RequestAudienceInfo `json:"audience_info,omitempty"`
	// 广告组Id
	CampaignId   int64                               `json:"campaign_id"`
	Classify     BrandAdCreateV30Classify            `json:"classify"`
	DateQuantity BrandAdCreateV30RequestDateQuantity `json:"date_quantity"`
	GdSendType   BrandAdCreateV30GdSendType          `json:"gd_send_type"`
	LandingType  BrandAdCreateV30LandingType         `json:"landing_type"`
	// 招商订单编号
	MerchantIntentionNo *string `json:"merchant_Intention_no,omitempty"`
	// 计划名称
	Name *string `json:"name,omitempty"`
	// 政策编号
	PolicyNo string `json:"policy_no"`
	// 备注
	Remark *string `json:"remark,omitempty"`
	// 代理商员工ID
	StaffId *int64 `json:"staff_id,omitempty"`
}
