/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10RequestChannelProductInfosInner struct for QianchuanAdCreateV10RequestChannelProductInfosInner
type QianchuanAdCreateV10RequestChannelProductInfosInner struct {
	//
	ChannelId   *int64                                              `json:"channel_id,omitempty"`
	ChannelType *QianchuanAdCreateV10ChannelProductInfosChannelType `json:"channel_type,omitempty"`
	//
	ProductId *int64 `json:"product_id,omitempty"`
}
