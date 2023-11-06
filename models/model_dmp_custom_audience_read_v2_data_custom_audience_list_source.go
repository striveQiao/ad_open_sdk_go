/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DmpCustomAudienceReadV2DataCustomAudienceListSource
type DmpCustomAudienceReadV2DataCustomAudienceListSource string

// List of dmp_custom_audience_read_v2_data_custom_audience_list_source
const (
	CUSTOM_AUDIENCE_TYPE_UPLOAD_DmpCustomAudienceReadV2DataCustomAudienceListSource      DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_UPLOAD"
	CUSTOM_AUDIENCE_TYPE_DOU_PLUS_DmpCustomAudienceReadV2DataCustomAudienceListSource    DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_DOU_PLUS"
	CUSTOM_AUDIENCE_TYPE_DATA_SOURCE_DmpCustomAudienceReadV2DataCustomAudienceListSource DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_DATA_SOURCE"
	CUSTOM_AUDIENCE_TYPE_OPERATE_DmpCustomAudienceReadV2DataCustomAudienceListSource     DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_OPERATE"
	CUSTOM_AUDIENCE_TYPE_FRIEND_DmpCustomAudienceReadV2DataCustomAudienceListSource      DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_FRIEND"
	CUSTOM_AUDIENCE_TYPE_THIRD_PARTY_DmpCustomAudienceReadV2DataCustomAudienceListSource DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_THIRD_PARTY"
	CUSTOM_AUDIENCE_TYPE_FINANCE_DmpCustomAudienceReadV2DataCustomAudienceListSource     DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_FINANCE"
	CUSTOM_AUDIENCE_TYPE_RULE_DmpCustomAudienceReadV2DataCustomAudienceListSource        DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_RULE"
	CUSTOM_AUDIENCE_TYPE_THEME_DmpCustomAudienceReadV2DataCustomAudienceListSource       DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_THEME"
	CUSTOM_AUDIENCE_TYPE_BRAND_DmpCustomAudienceReadV2DataCustomAudienceListSource       DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_BRAND"
	CUSTOM_AUDIENCE_TYPE_EXTEND_DmpCustomAudienceReadV2DataCustomAudienceListSource      DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_EXTEND"
	CUSTOM_AUDIENCE_TYPE_PACK_RULE_DmpCustomAudienceReadV2DataCustomAudienceListSource   DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_PACK_RULE"
	CUSTOM_AUDIENCE_TYPE_ONE_KEY_DmpCustomAudienceReadV2DataCustomAudienceListSource     DmpCustomAudienceReadV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_ONE_KEY"
)

// All allowed values of DmpCustomAudienceReadV2DataCustomAudienceListSource enum
var AllowedDmpCustomAudienceReadV2DataCustomAudienceListSourceEnumValues = []DmpCustomAudienceReadV2DataCustomAudienceListSource{
	"CUSTOM_AUDIENCE_TYPE_UPLOAD",
	"CUSTOM_AUDIENCE_TYPE_DOU_PLUS",
	"CUSTOM_AUDIENCE_TYPE_DATA_SOURCE",
	"CUSTOM_AUDIENCE_TYPE_OPERATE",
	"CUSTOM_AUDIENCE_TYPE_FRIEND",
	"CUSTOM_AUDIENCE_TYPE_THIRD_PARTY",
	"CUSTOM_AUDIENCE_TYPE_FINANCE",
	"CUSTOM_AUDIENCE_TYPE_RULE",
	"CUSTOM_AUDIENCE_TYPE_THEME",
	"CUSTOM_AUDIENCE_TYPE_BRAND",
	"CUSTOM_AUDIENCE_TYPE_EXTEND",
	"CUSTOM_AUDIENCE_TYPE_PACK_RULE",
	"CUSTOM_AUDIENCE_TYPE_ONE_KEY",
}

// NewDmpCustomAudienceReadV2DataCustomAudienceListSourceFromValue returns a pointer to a valid DmpCustomAudienceReadV2DataCustomAudienceListSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpCustomAudienceReadV2DataCustomAudienceListSourceFromValue(v string) (*DmpCustomAudienceReadV2DataCustomAudienceListSource, error) {
	ev := DmpCustomAudienceReadV2DataCustomAudienceListSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpCustomAudienceReadV2DataCustomAudienceListSource: valid values are %v", v, AllowedDmpCustomAudienceReadV2DataCustomAudienceListSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpCustomAudienceReadV2DataCustomAudienceListSource) IsValid() bool {
	for _, existing := range AllowedDmpCustomAudienceReadV2DataCustomAudienceListSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_custom_audience_read_v2_data_custom_audience_list_source value
func (v DmpCustomAudienceReadV2DataCustomAudienceListSource) Ptr() *DmpCustomAudienceReadV2DataCustomAudienceListSource {
	return &v
}
