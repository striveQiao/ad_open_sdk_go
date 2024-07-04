/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// Oauth2AdvertiserGetResponseDataListInner struct for Oauth2AdvertiserGetResponseDataListInner
type Oauth2AdvertiserGetResponseDataListInner struct {
	//
	AccountRole *string `json:"account_role,omitempty"`
	//
	AccountStringId *string `json:"account_string_id,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AdvertiserName *string `json:"advertiser_name,omitempty"`
	//
	AdvertiserRole *int64 `json:"advertiser_role,omitempty"`
	//
	CompanyList []*Oauth2AdvertiserGetResponseDataListInnerCompanyListInner `json:"company_list,omitempty"`
	//
	IsValid *bool `json:"is_valid,omitempty"`
}
