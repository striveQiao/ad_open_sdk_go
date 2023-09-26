/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

type ApiOpenApi2ToolsBidSuggestGetRequestExample struct {
	Ac                        []*ToolsBidSuggestV2Ac                         `json:"ac,omitempty"`
	ActionCategories          []int64                                        `json:"action_categories,omitempty"`
	ActionDays                ToolsBidSuggestV2ActionDays                    `json:"action_days,omitempty"`
	ActionScene               []*ToolsBidSuggestV2ActionScene                `json:"action_scene,omitempty"`
	ActionWords               []int64                                        `json:"action_words,omitempty"`
	ActivateType              []*ToolsBidSuggestV2ActivateType               `json:"activate_type,omitempty"`
	AdId                      int64                                          `json:"ad_id,omitempty"`
	AdTag                     []int64                                        `json:"ad_tag,omitempty"`
	AdvertiserId              int64                                          `json:"advertiser_id,omitempty"`
	Age                       []*ToolsBidSuggestV2Age                        `json:"age,omitempty"`
	AndroidOsv                ToolsBidSuggestV2AndroidOsv                    `json:"android_osv,omitempty"`
	AppBehaviorTarget         ToolsBidSuggestV2AppBehaviorTarget             `json:"app_behavior_target,omitempty"`
	AppCategory               []int64                                        `json:"app_category,omitempty"`
	AppIds                    []int64                                        `json:"app_ids,omitempty"`
	ArticleCategory           []*ToolsBidSuggestV2ArticleCategory            `json:"article_category,omitempty"`
	AudiencePackageId         int64                                          `json:"audience_package_id,omitempty"`
	AutoExtendTargets         []*ToolsBidSuggestV2AutoExtendTargets          `json:"auto_extend_targets,omitempty"`
	AwemeFanAccounts          []int64                                        `json:"aweme_fan_accounts,omitempty"`
	AwemeFanBehaviors         []*ToolsBidSuggestV2AwemeFanBehaviors          `json:"aweme_fan_behaviors,omitempty"`
	AwemeFanCategories        []int64                                        `json:"aweme_fan_categories,omitempty"`
	AwemeFanTimeScope         ToolsBidSuggestV2AwemeFanTimeScope             `json:"aweme_fan_time_scope,omitempty"`
	AwemeFansNumbers          []int64                                        `json:"aweme_fans_numbers,omitempty"`
	BidMode                   ToolsBidSuggestV2BidMode                       `json:"bid_mode,omitempty"`
	Budget                    int64                                          `json:"budget,omitempty"`
	BudgetMode                ToolsBidSuggestV2BudgetMode                    `json:"budget_mode,omitempty"`
	BusinessIds               []int64                                        `json:"business_ids,omitempty"`
	CampaignId                int64                                          `json:"campaign_id,omitempty"`
	Career                    []*ToolsBidSuggestV2Career                     `json:"career,omitempty"`
	Carrier                   []*ToolsBidSuggestV2Carrier                    `json:"carrier,omitempty"`
	City                      []int64                                        `json:"city,omitempty"`
	ConvertId                 int64                                          `json:"convert_id,omitempty"`
	ConvertedTimeDuration     ToolsBidSuggestV2ConvertedTimeDuration         `json:"converted_time_duration,omitempty"`
	DeviceBrand               []*ToolsBidSuggestV2DeviceBrand                `json:"device_brand,omitempty"`
	DeviceType                []*ToolsBidSuggestV2DeviceType                 `json:"device_type,omitempty"`
	District                  ToolsBidSuggestV2District                      `json:"district,omitempty"`
	DpaLocalAudience          ToolsBidSuggestV2DpaLocalAudience              `json:"dpa_local_audience,omitempty"`
	DpaRtaRecommendType       ToolsBidSuggestV2DpaRtaRecommendType           `json:"dpa_rta_recommend_type,omitempty"`
	DpaRtaSwitch              ToolsBidSuggestV2DpaRtaSwitch                  `json:"dpa_rta_switch,omitempty"`
	ExcludeCustomActions      []*ToolsBidSuggestV2ExcludeCustomActionsInner  `json:"exclude_custom_actions,omitempty"`
	ExcludeFlowPackage        []int64                                        `json:"exclude_flow_package,omitempty"`
	FilterAwemeAbnormalActive ToolsBidSuggestV2FilterAwemeAbnormalActive     `json:"filter_aweme_abnormal_active,omitempty"`
	FilterAwemeFansCount      int64                                          `json:"filter_aweme_fans_count,omitempty"`
	FilterOwnAwemeFans        ToolsBidSuggestV2FilterOwnAwemeFans            `json:"filter_own_aweme_fans,omitempty"`
	FlowControlMode           ToolsBidSuggestV2FlowControlMode               `json:"flow_control_mode,omitempty"`
	FlowPackage               []int64                                        `json:"flow_package,omitempty"`
	Gender                    ToolsBidSuggestV2Gender                        `json:"gender,omitempty"`
	Geolocation               []*AdGetV2ResponseDataAudienceGeolocationInner `json:"geolocation,omitempty"`
	HideIfConverted           ToolsBidSuggestV2HideIfConverted               `json:"hide_if_converted,omitempty"`
	HideIfExists              ToolsBidSuggestV2HideIfExists                  `json:"hide_if_exists,omitempty"`
	IncludeCustomActions      []*ToolsBidSuggestV2ExcludeCustomActionsInner  `json:"include_custom_actions,omitempty"`
	InterestActionMode        ToolsBidSuggestV2InterestActionMode            `json:"interest_action_mode,omitempty"`
	InterestCategories        []int64                                        `json:"interest_categories,omitempty"`
	InterestTags              []int64                                        `json:"interest_tags,omitempty"`
	InterestWords             []int64                                        `json:"interest_words,omitempty"`
	IosOsv                    ToolsBidSuggestV2IosOsv                        `json:"ios_osv,omitempty"`
	LaunchPrice               []int64                                        `json:"launch_price,omitempty"`
	LocationType              ToolsBidSuggestV2LocationType                  `json:"location_type,omitempty"`
	Platform                  []*ToolsBidSuggestV2Platform                   `json:"platform,omitempty"`
	Pricing                   ToolsBidSuggestV2Pricing                       `json:"pricing,omitempty"`
	RegionVersion             string                                         `json:"region_version,omitempty"`
	RetargetingTags           []int64                                        `json:"retargeting_tags,omitempty"`
	RetargetingTagsExclude    []int64                                        `json:"retargeting_tags_exclude,omitempty"`
	RetargetingTagsInclude    []int64                                        `json:"retargeting_tags_include,omitempty"`
	RetargetingType           ToolsBidSuggestV2RetargetingType               `json:"retargeting_type,omitempty"`
	ScheduleType              ToolsBidSuggestV2ScheduleType                  `json:"schedule_type,omitempty"`
	SuperiorPopularityType    ToolsBidSuggestV2SuperiorPopularityType        `json:"superior_popularity_type,omitempty"`
	UserType                  []int64                                        `json:"user_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/bid/suggest/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsBidSuggestGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsBidSuggestV2Api().
		Get(ctx).
		AccessToken(accessToken).
		Ac(request.Ac).ActionCategories(request.ActionCategories).ActionDays(request.ActionDays).ActionScene(request.ActionScene).ActionWords(request.ActionWords).ActivateType(request.ActivateType).AdId(request.AdId).AdTag(request.AdTag).AdvertiserId(request.AdvertiserId).Age(request.Age).AndroidOsv(request.AndroidOsv).AppBehaviorTarget(request.AppBehaviorTarget).AppCategory(request.AppCategory).AppIds(request.AppIds).ArticleCategory(request.ArticleCategory).AudiencePackageId(request.AudiencePackageId).AutoExtendTargets(request.AutoExtendTargets).AwemeFanAccounts(request.AwemeFanAccounts).AwemeFanBehaviors(request.AwemeFanBehaviors).AwemeFanCategories(request.AwemeFanCategories).AwemeFanTimeScope(request.AwemeFanTimeScope).AwemeFansNumbers(request.AwemeFansNumbers).BidMode(request.BidMode).Budget(request.Budget).BudgetMode(request.BudgetMode).BusinessIds(request.BusinessIds).CampaignId(request.CampaignId).Career(request.Career).Carrier(request.Carrier).City(request.City).ConvertId(request.ConvertId).ConvertedTimeDuration(request.ConvertedTimeDuration).DeviceBrand(request.DeviceBrand).DeviceType(request.DeviceType).District(request.District).DpaLocalAudience(request.DpaLocalAudience).DpaRtaRecommendType(request.DpaRtaRecommendType).DpaRtaSwitch(request.DpaRtaSwitch).ExcludeCustomActions(request.ExcludeCustomActions).ExcludeFlowPackage(request.ExcludeFlowPackage).FilterAwemeAbnormalActive(request.FilterAwemeAbnormalActive).FilterAwemeFansCount(request.FilterAwemeFansCount).FilterOwnAwemeFans(request.FilterOwnAwemeFans).FlowControlMode(request.FlowControlMode).FlowPackage(request.FlowPackage).Gender(request.Gender).Geolocation(request.Geolocation).HideIfConverted(request.HideIfConverted).HideIfExists(request.HideIfExists).IncludeCustomActions(request.IncludeCustomActions).InterestActionMode(request.InterestActionMode).InterestCategories(request.InterestCategories).InterestTags(request.InterestTags).InterestWords(request.InterestWords).IosOsv(request.IosOsv).LaunchPrice(request.LaunchPrice).LocationType(request.LocationType).Platform(request.Platform).Pricing(request.Pricing).RegionVersion(request.RegionVersion).RetargetingTags(request.RetargetingTags).RetargetingTagsExclude(request.RetargetingTagsExclude).RetargetingTagsInclude(request.RetargetingTagsInclude).RetargetingType(request.RetargetingType).ScheduleType(request.ScheduleType).SuperiorPopularityType(request.SuperiorPopularityType).UserType(request.UserType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
