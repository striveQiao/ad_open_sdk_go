/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponGetV2ResponseDataListInner struct for ClueCouponGetV2ResponseDataListInner
type ClueCouponGetV2ResponseDataListInner struct {
	//
	ActivityId   *int64                               `json:"activity_id,omitempty"`
	ActivityType *ClueCouponGetV2DataListActivityType `json:"activity_type,omitempty"`
	//
	BindFormId *int64                                      `json:"bind_form_id,omitempty"`
	Coupon     *ClueCouponGetV2ResponseDataListInnerCoupon `json:"coupon,omitempty"`
}
