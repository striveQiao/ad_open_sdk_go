/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaClueProductDetailV2DataProductsStatus
type DpaClueProductDetailV2DataProductsStatus string

// List of dpa_clue_product_detail_v2_data_products_status
const (
	STATUS_OFFLINE_DpaClueProductDetailV2DataProductsStatus DpaClueProductDetailV2DataProductsStatus = "STATUS_OFFLINE"
	STATUS_ONLINE_DpaClueProductDetailV2DataProductsStatus  DpaClueProductDetailV2DataProductsStatus = "STATUS_ONLINE"
)

// All allowed values of DpaClueProductDetailV2DataProductsStatus enum
var AllowedDpaClueProductDetailV2DataProductsStatusEnumValues = []DpaClueProductDetailV2DataProductsStatus{
	"STATUS_OFFLINE",
	"STATUS_ONLINE",
}

// NewDpaClueProductDetailV2DataProductsStatusFromValue returns a pointer to a valid DpaClueProductDetailV2DataProductsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaClueProductDetailV2DataProductsStatusFromValue(v string) (*DpaClueProductDetailV2DataProductsStatus, error) {
	ev := DpaClueProductDetailV2DataProductsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaClueProductDetailV2DataProductsStatus: valid values are %v", v, AllowedDpaClueProductDetailV2DataProductsStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaClueProductDetailV2DataProductsStatus) IsValid() bool {
	for _, existing := range AllowedDpaClueProductDetailV2DataProductsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_clue_product_detail_v2_data_products_status value
func (v DpaClueProductDetailV2DataProductsStatus) Ptr() *DpaClueProductDetailV2DataProductsStatus {
	return &v
}
