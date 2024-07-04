/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CgTransferWalletTransferListV30DataRecordListTransferTargetStatus
type CgTransferWalletTransferListV30DataRecordListTransferTargetStatus string

// List of cg_transfer_wallet_transfer_list_v3.0_data_record_list_transfer_target_status
const (
	NO_TRANSFER_CgTransferWalletTransferListV30DataRecordListTransferTargetStatus      CgTransferWalletTransferListV30DataRecordListTransferTargetStatus = "NO_TRANSFER"
	TRANSFER_FAILED_CgTransferWalletTransferListV30DataRecordListTransferTargetStatus  CgTransferWalletTransferListV30DataRecordListTransferTargetStatus = "TRANSFER_FAILED"
	TRANSFER_ING_CgTransferWalletTransferListV30DataRecordListTransferTargetStatus     CgTransferWalletTransferListV30DataRecordListTransferTargetStatus = "TRANSFER_ING"
	TRANSFER_PART_CgTransferWalletTransferListV30DataRecordListTransferTargetStatus    CgTransferWalletTransferListV30DataRecordListTransferTargetStatus = "TRANSFER_PART"
	TRANSFER_SUCCESS_CgTransferWalletTransferListV30DataRecordListTransferTargetStatus CgTransferWalletTransferListV30DataRecordListTransferTargetStatus = "TRANSFER_SUCCESS"
)

// All allowed values of CgTransferWalletTransferListV30DataRecordListTransferTargetStatus enum
var AllowedCgTransferWalletTransferListV30DataRecordListTransferTargetStatusEnumValues = []CgTransferWalletTransferListV30DataRecordListTransferTargetStatus{
	"NO_TRANSFER",
	"TRANSFER_FAILED",
	"TRANSFER_ING",
	"TRANSFER_PART",
	"TRANSFER_SUCCESS",
}

// NewCgTransferWalletTransferListV30DataRecordListTransferTargetStatusFromValue returns a pointer to a valid CgTransferWalletTransferListV30DataRecordListTransferTargetStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferWalletTransferListV30DataRecordListTransferTargetStatusFromValue(v string) (*CgTransferWalletTransferListV30DataRecordListTransferTargetStatus, error) {
	ev := CgTransferWalletTransferListV30DataRecordListTransferTargetStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferWalletTransferListV30DataRecordListTransferTargetStatus: valid values are %v", v, AllowedCgTransferWalletTransferListV30DataRecordListTransferTargetStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferWalletTransferListV30DataRecordListTransferTargetStatus) IsValid() bool {
	for _, existing := range AllowedCgTransferWalletTransferListV30DataRecordListTransferTargetStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_wallet_transfer_list_v3.0_data_record_list_transfer_target_status value
func (v CgTransferWalletTransferListV30DataRecordListTransferTargetStatus) Ptr() *CgTransferWalletTransferListV30DataRecordListTransferTargetStatus {
	return &v
}
