/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorUpdateV30Request struct for NativeAnchorUpdateV30Request
type NativeAnchorUpdateV30Request struct {
	//
	AdvertiserId int64                                  `json:"advertiser_id"`
	AnchorInfo   NativeAnchorUpdateV30RequestAnchorInfo `json:"anchor_info"`
}
