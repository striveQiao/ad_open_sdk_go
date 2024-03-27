/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaClueProductDetailV2DataProductsAuditStatus
type DpaClueProductDetailV2DataProductsAuditStatus string

// List of dpa_clue_product_detail_v2_data_products_audit_status
const (
	AUDIT_STATUS_APPROVE_DpaClueProductDetailV2DataProductsAuditStatus DpaClueProductDetailV2DataProductsAuditStatus = "AUDIT_STATUS_APPROVE"
	AUDIT_STATUS_INIT_DpaClueProductDetailV2DataProductsAuditStatus    DpaClueProductDetailV2DataProductsAuditStatus = "AUDIT_STATUS_INIT"
	AUDIT_STATUS_REJECT_DpaClueProductDetailV2DataProductsAuditStatus  DpaClueProductDetailV2DataProductsAuditStatus = "AUDIT_STATUS_REJECT"
)

// All allowed values of DpaClueProductDetailV2DataProductsAuditStatus enum
var AllowedDpaClueProductDetailV2DataProductsAuditStatusEnumValues = []DpaClueProductDetailV2DataProductsAuditStatus{
	"AUDIT_STATUS_APPROVE",
	"AUDIT_STATUS_INIT",
	"AUDIT_STATUS_REJECT",
}

// NewDpaClueProductDetailV2DataProductsAuditStatusFromValue returns a pointer to a valid DpaClueProductDetailV2DataProductsAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaClueProductDetailV2DataProductsAuditStatusFromValue(v string) (*DpaClueProductDetailV2DataProductsAuditStatus, error) {
	ev := DpaClueProductDetailV2DataProductsAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaClueProductDetailV2DataProductsAuditStatus: valid values are %v", v, AllowedDpaClueProductDetailV2DataProductsAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaClueProductDetailV2DataProductsAuditStatus) IsValid() bool {
	for _, existing := range AllowedDpaClueProductDetailV2DataProductsAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_clue_product_detail_v2_data_products_audit_status value
func (v DpaClueProductDetailV2DataProductsAuditStatus) Ptr() *DpaClueProductDetailV2DataProductsAuditStatus {
	return &v
}
