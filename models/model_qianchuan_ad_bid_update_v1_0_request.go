/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdBidUpdateV10Request struct for QianchuanAdBidUpdateV10Request
type QianchuanAdBidUpdateV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Data []*QianchuanAdBidUpdateV10RequestDataInner `json:"data"`
}
