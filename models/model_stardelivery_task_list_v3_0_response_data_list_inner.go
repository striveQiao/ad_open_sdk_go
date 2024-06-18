/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskListV30ResponseDataListInner struct for StardeliveryTaskListV30ResponseDataListInner
type StardeliveryTaskListV30ResponseDataListInner struct {
	//
	StarAdCostDivideRatio *float64 `json:"star_ad_cost_divide_ratio,omitempty"`
	//
	StarMaterialBid *float64 `json:"star_material_bid,omitempty"`
	//
	StarMaterialFirstName *string `json:"star_material_first_name,omitempty"`
	//
	StarMaterialFirstType *int64 `json:"star_material_first_type,omitempty"`
	//
	StarMaterialSecondName *string `json:"star_material_second_name,omitempty"`
	//
	StarMaterialSecondType *int64 `json:"star_material_second_type,omitempty"`
	//
	StarSaleCostEndTime *string                                            `json:"star_sale_cost_end_time,omitempty"`
	StarTaskAnchorType  *StardeliveryTaskListV30DataListStarTaskAnchorType `json:"star_task_anchor_type,omitempty"`
	//
	StarTaskBid *float64 `json:"star_task_bid,omitempty"`
	//
	StarTaskBudget *float64 `json:"star_task_budget,omitempty"`
	//
	StarTaskCanDeliveryItemCount *int64 `json:"star_task_can_delivery_item_count,omitempty"`
	//
	StarTaskCreateTime     *string                                                `json:"star_task_create_time,omitempty"`
	StarTaskExternalAction *StardeliveryTaskListV30DataListStarTaskExternalAction `json:"star_task_external_action,omitempty"`
	//
	StarTaskId *int64 `json:"star_task_id,omitempty"`
	//
	StarTaskModifyTime *string `json:"star_task_modify_time,omitempty"`
	//
	StarTaskName *string `json:"star_task_name,omitempty"`
	//
	StarTaskPostEndTime *string `json:"star_task_post_end_time,omitempty"`
	//
	StarTaskPostItemCount *int64 `json:"star_task_post_item_count,omitempty"`
	//
	StarTaskRelateProjectCount *int64                                         `json:"star_task_relate_project_count,omitempty"`
	StarTaskSource             *StardeliveryTaskListV30DataListStarTaskSource `json:"star_task_source,omitempty"`
	//
	StarTaskStartTime *string                                           `json:"star_task_start_time,omitempty"`
	StarTaskStatus    *StardeliveryTaskListV30DataListStarTaskStatus    `json:"star_task_status,omitempty"`
	StarTaskSubStatus *StardeliveryTaskListV30DataListStarTaskSubStatus `json:"star_task_sub_status,omitempty"`
	//
	StarTimeCostEndTime *string `json:"star_time_cost_end_time,omitempty"`
}
