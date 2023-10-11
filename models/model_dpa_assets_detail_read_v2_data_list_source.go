/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaAssetsDetailReadV2DataListSource
type DpaAssetsDetailReadV2DataListSource string

// List of dpa_assets_detail_read_v2_data_list_source
const (
	AUTO_DpaAssetsDetailReadV2DataListSource   DpaAssetsDetailReadV2DataListSource = "AUTO"
	MANUAL_DpaAssetsDetailReadV2DataListSource DpaAssetsDetailReadV2DataListSource = "MANUAL"
)

// All allowed values of DpaAssetsDetailReadV2DataListSource enum
var AllowedDpaAssetsDetailReadV2DataListSourceEnumValues = []DpaAssetsDetailReadV2DataListSource{
	"AUTO",
	"MANUAL",
}

// NewDpaAssetsDetailReadV2DataListSourceFromValue returns a pointer to a valid DpaAssetsDetailReadV2DataListSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetsDetailReadV2DataListSourceFromValue(v string) (*DpaAssetsDetailReadV2DataListSource, error) {
	ev := DpaAssetsDetailReadV2DataListSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetsDetailReadV2DataListSource: valid values are %v", v, AllowedDpaAssetsDetailReadV2DataListSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetsDetailReadV2DataListSource) IsValid() bool {
	for _, existing := range AllowedDpaAssetsDetailReadV2DataListSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_assets_detail_read_v2_data_list_source value
func (v DpaAssetsDetailReadV2DataListSource) Ptr() *DpaAssetsDetailReadV2DataListSource {
	return &v
}
