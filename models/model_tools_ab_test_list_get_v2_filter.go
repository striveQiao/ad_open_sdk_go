/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestListGetV2Filter
type ToolsAbTestListGetV2Filter struct {
	// 根据实验结论过滤，允许值：`\"NOT_START\"`：未开始，`\"OBVIOUS\"`：有明显结论，`\"INSUFFICIENT\"`：数据量不足；`\"FAILED\"`：不满足实验条件，`\"TMP_OBVIOUS\"`：有初步结论。
	Conclusions []*ToolsAbTestListGetV2FilterConclusions `json:"conclusions,omitempty"`
	// 返回创建时间在此之后的实验，格式：\"2020-12-25 07:12:08\"
	CreateTimeAfter *string `json:"create_time_after,omitempty"`
	// 返回创建时间在此之前的实验，格式：\"2020-12-25 07:12:08\"
	CreateTimeBefore *string `json:"create_time_before,omitempty"`
	// 根据实验状态过滤，允许值：`\"CREATED\"`：排期中，`\"PROCESSING\"`：进行中，`\"FINISH\"`：结束，`\"FAILED\"`：不满足实验条件。
	Status []*ToolsAbTestListGetV2FilterStatus `json:"status,omitempty"`
	// 根据实验ID列表过滤，列表长度1-100
	TestIds []int64 `json:"test_ids,omitempty"`
	// 根据实验名称过滤，支持模糊匹配，长度限制100字符，中文算2个字符。
	TestName *string `json:"test_name,omitempty"`
	// 返回测试时间在此之后的实验，格式：\"2020-12-25 07:12:08\"
	TestTimeAfter *string `json:"test_time_after,omitempty"`
	// 返回测试时间在此之前的实验，格式：\"2020-12-25 07:12:08\"
	TestTimeBefore *string `json:"test_time_before,omitempty"`
}
