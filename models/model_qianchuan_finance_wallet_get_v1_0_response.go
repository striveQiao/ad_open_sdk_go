/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFinanceWalletGetV10Response struct for QianchuanFinanceWalletGetV10Response
type QianchuanFinanceWalletGetV10Response struct {
	//
	Code *int64                                    `json:"code,omitempty"`
	Data *QianchuanFinanceWalletGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
