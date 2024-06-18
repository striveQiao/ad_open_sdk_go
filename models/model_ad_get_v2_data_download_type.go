/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataDownloadType
type AdGetV2DataDownloadType string

// List of ad_get_v2_data_download_type
const (
	EXTERNAL_URL_AdGetV2DataDownloadType AdGetV2DataDownloadType = "EXTERNAL_URL"
	DOWNLOAD_URL_AdGetV2DataDownloadType AdGetV2DataDownloadType = "DOWNLOAD_URL"
)

// All allowed values of AdGetV2DataDownloadType enum
var AllowedAdGetV2DataDownloadTypeEnumValues = []AdGetV2DataDownloadType{
	"EXTERNAL_URL",
	"DOWNLOAD_URL",
}

// NewAdGetV2DataDownloadTypeFromValue returns a pointer to a valid AdGetV2DataDownloadType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataDownloadTypeFromValue(v string) (*AdGetV2DataDownloadType, error) {
	ev := AdGetV2DataDownloadType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataDownloadType: valid values are %v", v, AllowedAdGetV2DataDownloadTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataDownloadType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataDownloadTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_download_type value
func (v AdGetV2DataDownloadType) Ptr() *AdGetV2DataDownloadType {
	return &v
}
