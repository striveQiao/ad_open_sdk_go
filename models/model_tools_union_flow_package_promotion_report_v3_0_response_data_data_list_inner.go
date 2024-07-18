/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsUnionFlowPackagePromotionReportV30ResponseDataDataListInner struct for ToolsUnionFlowPackagePromotionReportV30ResponseDataDataListInner
type ToolsUnionFlowPackagePromotionReportV30ResponseDataDataListInner struct {
	// 激活数
	Active *float64 `json:"active,omitempty"`
	// 激活成本
	ActiveCost *float64 `json:"active_cost,omitempty"`
	// 付费数
	ActivePay *float64 `json:"active_pay,omitempty"`
	// 付费成本
	ActivePayCost *float64 `json:"active_pay_cost,omitempty"`
	// 付费率
	ActivePayRate *float64 `json:"active_pay_rate,omitempty"`
	// 注册数
	ActiveRegister *float64 `json:"active_register,omitempty"`
	// 注册成本
	ActiveRegisterCost *float64 `json:"active_register_cost,omitempty"`
	// 注册率
	ActiveRegisterRate *float64 `json:"active_register_rate,omitempty"`
	// 点击指标
	ClickCnt *float64 `json:"click_cnt,omitempty"`
	// 转化成本
	ConversionCost *float64 `json:"conversion_cost,omitempty"`
	// 转化量
	ConvertCnt *float64 `json:"convert_cnt,omitempty"`
	// 表单提交数
	Form *float64 `json:"form,omitempty"`
	// 表单
	FormCost *float64 `json:"form_cost,omitempty"`
	// 关键行为数
	GameAddiction *float64 `json:"game_addiction,omitempty"`
	// 关键行为成本
	GameAddictionCost *float64 `json:"game_addiction_cost,omitempty"`
	// 关键行为率
	GameAddictionRate *float64 `json:"game_addiction_rate,omitempty"`
	// 完件数
	LoanCompletion *float64 `json:"loan_completion,omitempty"`
	// 完件成本
	LoanCompletionCost *float64 `json:"loan_completion_cost,omitempty"`
	// 授信数
	LoanCredit *float64 `json:"loan_credit,omitempty"`
	// 授信成本
	LoanCreditCost *float64 `json:"loan_credit_cost,omitempty"`
	// 次留数
	NextDayOpen *float64 `json:"next_day_open,omitempty"`
	// 次留成本
	NextDayOpenCost *float64 `json:"next_day_open_cost,omitempty"`
	// 次留率
	NextDayOpenRate *float64 `json:"next_day_open_rate,omitempty"`
	// 预授信数
	PreLoanCredit *float64 `json:"pre_loan_credit,omitempty"`
	// 预授信成本
	PreLoanCreditCost *float64 `json:"pre_loan_credit_cost,omitempty"`
	// 广告位
	Rit *int64 `json:"rit,omitempty"`
	// 展示指标
	ShowCnt *float64 `json:"show_cnt,omitempty"`
	// 消耗，单位：元
	StatCost *float64 `json:"stat_cost,omitempty"`
}
