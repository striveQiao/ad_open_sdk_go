/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2RequestGeolocationInner struct for AudiencePackageUpdateV2RequestGeolocationInner
type AudiencePackageUpdateV2RequestGeolocationInner struct {
	//
	City *string `json:"city,omitempty"`
	//
	District *string `json:"district,omitempty"`
	//
	Lat *float64 `json:"lat,omitempty"`
	//
	Long *float64 `json:"long,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Province *string `json:"province,omitempty"`
	//
	Radius *int64 `json:"radius,omitempty"`
	//
	Street *string `json:"street,omitempty"`
	//
	StreetNumber *string `json:"street_number,omitempty"`
}
