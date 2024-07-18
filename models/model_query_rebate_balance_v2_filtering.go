/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryRebateBalanceV2Filtering
type QueryRebateBalanceV2Filtering struct {
	ApprovalStatus *QueryRebateBalanceV2FilteringApprovalStatus `json:"approval_status,omitempty"`
	// 回收状态 （1：未回收，2：已回收，3：已归档）
	ArchiveStatusList []*QueryRebateBalanceV2FilteringArchiveStatusList `json:"archive_status_list,omitempty"`
	// 合同编号
	ContractSerial *string `json:"contract_serial,omitempty"`
	// 创建时间范围结束时间
	CreateEndDate *string `json:"create_end_date,omitempty"`
	// 创建时间范围开始时间
	CreateStartDate *string `json:"create_start_date,omitempty"`
	// 开票状态 （1：未开票，2：部分开票，3：已开票）
	InvoiceStatusList []*QueryRebateBalanceV2FilteringInvoiceStatusList `json:"invoice_status_list,omitempty"`
	MonthQuarter      *QueryRebateBalanceV2FilteringMonthQuarter        `json:"month_quarter,omitempty"`
	//
	Platforms []*QueryRebateBalanceV2FilteringPlatforms `json:"platforms,omitempty"`
	// 返点流水编号 （与平台披露编号一致，建议使用）
	RebateBalanceSerial *string `json:"rebate_balance_serial,omitempty"`
	// 盖章状态 （1：未盖章，2：审批中，3：已盖章）
	StampStatusList []*QueryRebateBalanceV2FilteringStampStatusList `json:"stamp_status_list,omitempty"`
	StampType       *QueryRebateBalanceV2FilteringStampType         `json:"stamp_type,omitempty"`
	Standard        *QueryRebateBalanceV2FilteringStandard          `json:"standard,omitempty"`
	UseType         *QueryRebateBalanceV2FilteringUseType           `json:"use_type,omitempty"`
	// 结算周期年
	Year *int32 `json:"year,omitempty"`
}
