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

// CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType
type CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType string

// List of cg_transfer_query_transfer_detail_v3.0_data_transfer_target_record_list_transfer_capital_record_list_capital_type
const (
	CREDIT_BIDDING_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType = "CREDIT_BIDDING"
	CREDIT_BRAND_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType   CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType = "CREDIT_BRAND"
	CREDIT_GENERAL_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType = "CREDIT_GENERAL"
	PREPAY_BIDDING_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType = "PREPAY_BIDDING"
	PREPAY_BRAND_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType   CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType = "PREPAY_BRAND"
	PREPAY_GENERAL_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType = "PREPAY_GENERAL"
)

// All allowed values of CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType enum
var AllowedCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalTypeEnumValues = []CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType{
	"CREDIT_BIDDING",
	"CREDIT_BRAND",
	"CREDIT_GENERAL",
	"PREPAY_BIDDING",
	"PREPAY_BRAND",
	"PREPAY_GENERAL",
}

// NewCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalTypeFromValue returns a pointer to a valid CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalTypeFromValue(v string) (*CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType, error) {
	ev := CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType: valid values are %v", v, AllowedCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType) IsValid() bool {
	for _, existing := range AllowedCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_query_transfer_detail_v3.0_data_transfer_target_record_list_transfer_capital_record_list_capital_type value
func (v CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType) Ptr() *CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType {
	return &v
}
