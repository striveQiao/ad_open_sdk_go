/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderCloseV30Request struct for DouplusOrderCloseV30Request
type DouplusOrderCloseV30Request struct {
	//
	AwemeSecUid string `json:"aweme_sec_uid"`
	// 需要关停的订单ID列表，最多10
	OrderIds []int64 `json:"order_ids"`
}
