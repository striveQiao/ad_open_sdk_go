/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserQualificationSubmitV30Request struct for AdvertiserQualificationSubmitV30Request
type AdvertiserQualificationSubmitV30Request struct {
	// 广告主ID
	AdvertiserId int64                                          `json:"advertiser_id"`
	Subject      AdvertiserQualificationSubmitV30RequestSubject `json:"subject"`
}
