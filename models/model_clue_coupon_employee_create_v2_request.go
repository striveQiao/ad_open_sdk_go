/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponEmployeeCreateV2Request struct for ClueCouponEmployeeCreateV2Request
type ClueCouponEmployeeCreateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	List []*ClueCouponEmployeeCreateV2RequestListInner `json:"list"`
}
