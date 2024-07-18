/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponDetailV2ResponseDataCoupon
type ClueCouponDetailV2ResponseDataCoupon struct {
	//
	CouponId *int64 `json:"coupon_id,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	DeliverEnd **string `json:"deliver_end,omitempty"`
	//
	DeliverStart  **string                                         `json:"deliver_start,omitempty"`
	GlobalLimit   *ClueCouponDetailV2ResponseDataCouponGlobalLimit `json:"global_limit,omitempty"`
	NeedPhone     *ClueCouponDetailV2DataCouponNeedPhone           `json:"need_phone,omitempty"`
	NeedSmsVerify *ClueCouponDetailV2DataCouponNeedSmsVerify       `json:"need_sms_verify,omitempty"`
	//
	ResourceList []*ClueCouponDetailV2ResponseDataCouponResourceListInner `json:"resource_list,omitempty"`
	Status       *ClueCouponDetailV2DataCouponStatus                      `json:"status,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	UpdateTime *string                                        `json:"update_time,omitempty"`
	UserLimit  *ClueCouponDetailV2ResponseDataCouponUserLimit `json:"user_limit,omitempty"`
}
