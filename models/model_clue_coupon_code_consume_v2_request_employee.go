/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponCodeConsumeV2RequestEmployee
type ClueCouponCodeConsumeV2RequestEmployee struct {
	//
	StoreId *int64 `json:"store_id,omitempty"`
	//
	UserId   int64                                   `json:"user_id"`
	UserType ClueCouponCodeConsumeV2EmployeeUserType `json:"user_type"`
}
