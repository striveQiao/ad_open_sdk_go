/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateCreateV2ResponseDataBricksInnerButtonPhoneEvent phoneEvent事件行为描述
type ToolsSiteTemplateCreateV2ResponseDataBricksInnerButtonPhoneEvent struct {
	// 智能电话ID，当phone_event不为空时，有值。用户可以通过[【获取智能电话列表】](https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710644556812 )接口或[【青鸟线索通平台】]（https://clue.oceanengine.com/）获取智能电话ID
	InstanceId int64 `json:"instance_id"`
}
