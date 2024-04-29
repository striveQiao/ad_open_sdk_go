/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponGetV2ResponseDataListInnerCoupon
type ClueCouponGetV2ResponseDataListInnerCoupon struct {
	//
	CouponId *int64 `json:"coupon_id,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	DeliverEnd **string `json:"deliver_end,omitempty"`
	//
	DeliverStart  **string                                               `json:"deliver_start,omitempty"`
	GlobalLimit   *ClueCouponGetV2ResponseDataListInnerCouponGlobalLimit `json:"global_limit,omitempty"`
	NeedPhone     *ClueCouponGetV2DataListCouponNeedPhone                `json:"need_phone,omitempty"`
	NeedSmsVerify *ClueCouponGetV2DataListCouponNeedSmsVerify            `json:"need_sms_verify,omitempty"`
	//
	ResourceList []*ClueCouponGetV2ResponseDataListInnerCouponResourceListInner `json:"resource_list,omitempty"`
	Status       *ClueCouponGetV2DataListCouponStatus                           `json:"status,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	UpdateTime *string                                              `json:"update_time,omitempty"`
	UserLimit  *ClueCouponGetV2ResponseDataListInnerCouponUserLimit `json:"user_limit,omitempty"`
}
