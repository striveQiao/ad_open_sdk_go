/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DmpCustomAudienceSelectV2DataCustomAudienceListSource
type DmpCustomAudienceSelectV2DataCustomAudienceListSource string

// List of dmp_custom_audience_select_v2_data_custom_audience_list_source
const (
	CUSTOM_AUDIENCE_TYPE_FINANCE_DmpCustomAudienceSelectV2DataCustomAudienceListSource     DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_FINANCE"
	CUSTOM_AUDIENCE_TYPE_DATA_SOURCE_DmpCustomAudienceSelectV2DataCustomAudienceListSource DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_DATA_SOURCE"
	CUSTOM_AUDIENCE_TYPE_RULE_DmpCustomAudienceSelectV2DataCustomAudienceListSource        DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_RULE"
	CUSTOM_AUDIENCE_TYPE_OPERATE_DmpCustomAudienceSelectV2DataCustomAudienceListSource     DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_OPERATE"
	CUSTOM_AUDIENCE_TYPE_PACK_RULE_DmpCustomAudienceSelectV2DataCustomAudienceListSource   DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_PACK_RULE"
	CUSTOM_AUDIENCE_TYPE_THEME_DmpCustomAudienceSelectV2DataCustomAudienceListSource       DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_THEME"
	CUSTOM_AUDIENCE_TYPE_THIRD_PARTY_DmpCustomAudienceSelectV2DataCustomAudienceListSource DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_THIRD_PARTY"
	CUSTOM_AUDIENCE_TYPE_FRIEND_DmpCustomAudienceSelectV2DataCustomAudienceListSource      DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_FRIEND"
	CUSTOM_AUDIENCE_TYPE_BRAND_DmpCustomAudienceSelectV2DataCustomAudienceListSource       DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_BRAND"
	CUSTOM_AUDIENCE_TYPE_UPLOAD_DmpCustomAudienceSelectV2DataCustomAudienceListSource      DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_UPLOAD"
	CUSTOM_AUDIENCE_TYPE_DOU_PLUS_DmpCustomAudienceSelectV2DataCustomAudienceListSource    DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_DOU_PLUS"
	CUSTOM_AUDIENCE_TYPE_ONE_KEY_DmpCustomAudienceSelectV2DataCustomAudienceListSource     DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_ONE_KEY"
	CUSTOM_AUDIENCE_TYPE_EXTEND_DmpCustomAudienceSelectV2DataCustomAudienceListSource      DmpCustomAudienceSelectV2DataCustomAudienceListSource = "CUSTOM_AUDIENCE_TYPE_EXTEND"
)

// All allowed values of DmpCustomAudienceSelectV2DataCustomAudienceListSource enum
var AllowedDmpCustomAudienceSelectV2DataCustomAudienceListSourceEnumValues = []DmpCustomAudienceSelectV2DataCustomAudienceListSource{
	"CUSTOM_AUDIENCE_TYPE_FINANCE",
	"CUSTOM_AUDIENCE_TYPE_DATA_SOURCE",
	"CUSTOM_AUDIENCE_TYPE_RULE",
	"CUSTOM_AUDIENCE_TYPE_OPERATE",
	"CUSTOM_AUDIENCE_TYPE_PACK_RULE",
	"CUSTOM_AUDIENCE_TYPE_THEME",
	"CUSTOM_AUDIENCE_TYPE_THIRD_PARTY",
	"CUSTOM_AUDIENCE_TYPE_FRIEND",
	"CUSTOM_AUDIENCE_TYPE_BRAND",
	"CUSTOM_AUDIENCE_TYPE_UPLOAD",
	"CUSTOM_AUDIENCE_TYPE_DOU_PLUS",
	"CUSTOM_AUDIENCE_TYPE_ONE_KEY",
	"CUSTOM_AUDIENCE_TYPE_EXTEND",
}

// NewDmpCustomAudienceSelectV2DataCustomAudienceListSourceFromValue returns a pointer to a valid DmpCustomAudienceSelectV2DataCustomAudienceListSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpCustomAudienceSelectV2DataCustomAudienceListSourceFromValue(v string) (*DmpCustomAudienceSelectV2DataCustomAudienceListSource, error) {
	ev := DmpCustomAudienceSelectV2DataCustomAudienceListSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpCustomAudienceSelectV2DataCustomAudienceListSource: valid values are %v", v, AllowedDmpCustomAudienceSelectV2DataCustomAudienceListSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpCustomAudienceSelectV2DataCustomAudienceListSource) IsValid() bool {
	for _, existing := range AllowedDmpCustomAudienceSelectV2DataCustomAudienceListSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_custom_audience_select_v2_data_custom_audience_list_source value
func (v DmpCustomAudienceSelectV2DataCustomAudienceListSource) Ptr() *DmpCustomAudienceSelectV2DataCustomAudienceListSource {
	return &v
}
