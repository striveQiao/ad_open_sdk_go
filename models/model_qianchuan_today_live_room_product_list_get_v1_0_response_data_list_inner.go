/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanTodayLiveRoomProductListGetV10ResponseDataListInner struct for QianchuanTodayLiveRoomProductListGetV10ResponseDataListInner
type QianchuanTodayLiveRoomProductListGetV10ResponseDataListInner struct {
	//
	AdLiveOrderSettleCostPerProduct7d *float64 `json:"ad_live_order_settle_cost_per_product_7d,omitempty"`
	//
	AdLiveOrderSettleRoiPerProduct7d *float64                                                      `json:"ad_live_order_settle_roi_per_product_7d,omitempty"`
	ExplainStatus                    *QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus `json:"explain_status,omitempty"`
	//
	ImgUrl *string `json:"img_url,omitempty"`
	//
	LiveCostPerProduct *float64 `json:"live_cost_per_product,omitempty"`
	//
	LiveOrderRefundAmount *float64 `json:"live_order_refund_amount,omitempty"`
	//
	LiveOrderSettleAmount7d *float64 `json:"live_order_settle_amount_7d,omitempty"`
	//
	LiveOrderSettleCount7d *int64 `json:"live_order_settle_count_7d,omitempty"`
	//
	LiveOrderSettleCountRate7dliveOrderSettleCountRate7d *float64 `json:"live_order_settle_count_rate_7dlive_order_settle_count_rate_7d,omitempty"`
	//
	LivePayOrderCountAlias *int64 `json:"live_pay_order_count_alias,omitempty"`
	//
	LivePayOrderGmvAlias *float64 `json:"live_pay_order_gmv_alias,omitempty"`
	//
	LivePayOrderRoiPerProduct *float64 `json:"live_pay_order_roi_per_product,omitempty"`
	//
	LiveProductInventory *int64 `json:"live_product_Inventory,omitempty"`
	//
	LiveProductBindTime *string `json:"live_product_bind_time,omitempty"`
	//
	LiveProductExplainCnt *int64 `json:"live_product_explain_cnt,omitempty"`
	//
	LiveProductPrice *float64 `json:"live_product_price,omitempty"`
	//
	LubanLivePayOrderCount *int64 `json:"luban_live_pay_order_count,omitempty"`
	//
	LubanLivePayOrderGmv *float64 `json:"luban_live_pay_order_gmv,omitempty"`
	//
	ProductId *int64 `json:"productId,omitempty"`
	//
	ProductClickPayUcntRatio *float64 `json:"product_click_pay_ucnt_ratio,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	TotalLivePayOrderGpmEcom *float64 `json:"total_live_pay_order_gpm_ecom,omitempty"`
}
