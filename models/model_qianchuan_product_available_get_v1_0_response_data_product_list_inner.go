/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAvailableGetV10ResponseDataProductListInner struct for QianchuanProductAvailableGetV10ResponseDataProductListInner
type QianchuanProductAvailableGetV10ResponseDataProductListInner struct {
	// 商品权益
	Benefits []string `json:"benefits,omitempty"`
	//
	CategoryName *string `json:"category_name,omitempty"`
	//
	ChannelId   *int64                                                     `json:"channel_id,omitempty"`
	ChannelType *QianchuanProductAvailableGetV10DataProductListChannelType `json:"channel_type,omitempty"`
	//
	DiscountHigherPrice *float64 `json:"discount_higher_price,omitempty"`
	//
	DiscountLowerPrice *float64 `json:"discount_lower_price,omitempty"`
	//
	DiscountPrice *int64 `json:"discount_price,omitempty"`
	// 首次上架时间
	FirstOnShelfTime *int64 `json:"first_on_shelf_time,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	Img *string `json:"img,omitempty"`
	//
	ImgList []*QianchuanProductAvailableGetV10ResponseDataProductListInnerImgListInner `json:"img_list,omitempty"`
	//
	Inventory *int64 `json:"inventory,omitempty"`
	//
	MarketPrice *float64 `json:"market_price,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	ProductRate *int64 `json:"product_rate,omitempty"`
	// 推荐理由
	RecommendReasons []string `json:"recommend_reasons,omitempty"`
	//
	SaleTime *string `json:"sale_time,omitempty"`
	// 销量
	SellNum *int64 `json:"sell_num,omitempty"`
	//
	SupportProductNewOpen *bool                                               `json:"support_product_new_open,omitempty"`
	Tags                  *QianchuanProductAvailableGetV10DataProductListTags `json:"tags,omitempty"`
}
