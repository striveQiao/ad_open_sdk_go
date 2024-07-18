/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryRebateAccountingInfoV2Filtering
type QueryRebateAccountingInfoV2Filtering struct {
	ApprovalStatus *QueryRebateAccountingInfoV2FilteringApprovalStatus `json:"approval_status,omitempty"`
	// 合同编号
	ContractSerial *string                                             `json:"contract_serial,omitempty"`
	IsCreateRebate *QueryRebateAccountingInfoV2FilteringIsCreateRebate `json:"is_create_rebate,omitempty"`
	MonthQuarter   *QueryRebateAccountingInfoV2FilteringMonthQuarter   `json:"month_quarter,omitempty"`
	// 业务线/平台 (0:AD 3:星图 10031:广告 10063:巨量本地推)
	Platforms []*QueryRebateAccountingInfoV2FilteringPlatforms `json:"platforms,omitempty"`
	// 返点核算信息编号
	RebateAccountingInfoSerial *string `json:"rebate_accounting_info_serial,omitempty"`
	// 返点流水号
	RebateBalanceSerial *string                                      `json:"rebate_balance_serial,omitempty"`
	UseType             *QueryRebateAccountingInfoV2FilteringUseType `json:"use_type,omitempty"`
	// 结算周期年 例:2024
	Year *int64 `json:"year,omitempty"`
}
