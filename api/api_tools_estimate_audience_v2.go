/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ToolsEstimateAudienceV2ApiService ToolsEstimateAudienceV2Api service
type ToolsEstimateAudienceV2ApiService service

type ApiOpenApi2ToolsEstimateAudienceGetRequest struct {
	ctx                       context.Context
	ApiService                *ToolsEstimateAudienceV2ApiService
	ac                        *[]*ToolsEstimateAudienceV2Ac
	actionCategories          *[]int64
	actionDays                *ToolsEstimateAudienceV2ActionDays
	actionWords               *[]int64
	activateType              *[]*ToolsEstimateAudienceV2ActivateType
	adTag                     *[]int64
	advertiserId              *int64
	age                       *[]*ToolsEstimateAudienceV2Age
	androidOsv                *ToolsEstimateAudienceV2AndroidOsv
	appBehaviorTarget         *ToolsEstimateAudienceV2AppBehaviorTarget
	appCategory               *[]int64
	appIds                    *[]int64
	articleCategory           *[]*ToolsEstimateAudienceV2ArticleCategory
	audiencePackageId         *int64
	autoExtendEnabled         *ToolsEstimateAudienceV2AutoExtendEnabled
	autoExtendTargets         *[]*ToolsEstimateAudienceV2AutoExtendTargets
	awemeFansNumbers          *[]int64
	carrier                   *[]*ToolsEstimateAudienceV2Carrier
	city                      *[]int64
	deviceBrand               *[]*ToolsEstimateAudienceV2DeviceBrand
	deviceType                *[]*ToolsEstimateAudienceV2DeviceType
	district                  *ToolsEstimateAudienceV2District
	dpaLocalAudience          *ToolsEstimateAudienceV2DpaLocalAudience
	excludeCustomActions      *[]*ToolsEstimateAudienceV2ExcludeCustomActionsInner
	excludeFlowPackage        *[]int64
	filterAwemeAbnormalActive *ToolsEstimateAudienceV2FilterAwemeAbnormalActive
	filterAwemeFansCount      *int64
	filterOwnAwemeFans        *ToolsEstimateAudienceV2FilterOwnAwemeFans
	flowPackage               *[]int64
	gender                    *ToolsEstimateAudienceV2Gender
	geolocation               *[]*ToolsEstimateAudienceV2GeolocationInner
	includeCustomActions      *[]*ToolsEstimateAudienceV2IncludeCustomActionsInner
	interestActionMode        *ToolsEstimateAudienceV2InterestActionMode
	interestCategories        *[]int64
	interestTags              *[]int64
	interestWords             *[]int64
	iosOsv                    *ToolsEstimateAudienceV2IosOsv
	launchPrice               *[]int64
	locationType              *ToolsEstimateAudienceV2LocationType
	platform                  *[]*ToolsEstimateAudienceV2Platform
	retargetingTags           *[]int64
	retargetingTagsExclude    *[]int64
	retargetingTagsInclude    *[]int64
	retargetingType           *ToolsEstimateAudienceV2RetargetingType
	superiorPopularityType    *ToolsEstimateAudienceV2SuperiorPopularityType
	userType                  *[]int64
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) Ac(ac []*ToolsEstimateAudienceV2Ac) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.ac = &ac
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) ActionCategories(actionCategories []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.actionCategories = &actionCategories
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) ActionDays(actionDays ToolsEstimateAudienceV2ActionDays) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.actionDays = &actionDays
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) ActionWords(actionWords []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.actionWords = &actionWords
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) ActivateType(activateType []*ToolsEstimateAudienceV2ActivateType) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.activateType = &activateType
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) AdTag(adTag []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.adTag = &adTag
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) Age(age []*ToolsEstimateAudienceV2Age) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.age = &age
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) AndroidOsv(androidOsv ToolsEstimateAudienceV2AndroidOsv) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.androidOsv = &androidOsv
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) AppBehaviorTarget(appBehaviorTarget ToolsEstimateAudienceV2AppBehaviorTarget) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.appBehaviorTarget = &appBehaviorTarget
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) AppCategory(appCategory []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.appCategory = &appCategory
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) AppIds(appIds []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.appIds = &appIds
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) ArticleCategory(articleCategory []*ToolsEstimateAudienceV2ArticleCategory) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.articleCategory = &articleCategory
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) AudiencePackageId(audiencePackageId int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.audiencePackageId = &audiencePackageId
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) AutoExtendEnabled(autoExtendEnabled ToolsEstimateAudienceV2AutoExtendEnabled) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.autoExtendEnabled = &autoExtendEnabled
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) AutoExtendTargets(autoExtendTargets []*ToolsEstimateAudienceV2AutoExtendTargets) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.autoExtendTargets = &autoExtendTargets
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) AwemeFansNumbers(awemeFansNumbers []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.awemeFansNumbers = &awemeFansNumbers
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) Carrier(carrier []*ToolsEstimateAudienceV2Carrier) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.carrier = &carrier
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) City(city []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.city = &city
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) DeviceBrand(deviceBrand []*ToolsEstimateAudienceV2DeviceBrand) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.deviceBrand = &deviceBrand
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) DeviceType(deviceType []*ToolsEstimateAudienceV2DeviceType) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.deviceType = &deviceType
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) District(district ToolsEstimateAudienceV2District) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.district = &district
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) DpaLocalAudience(dpaLocalAudience ToolsEstimateAudienceV2DpaLocalAudience) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.dpaLocalAudience = &dpaLocalAudience
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) ExcludeCustomActions(excludeCustomActions []*ToolsEstimateAudienceV2ExcludeCustomActionsInner) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.excludeCustomActions = &excludeCustomActions
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) ExcludeFlowPackage(excludeFlowPackage []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.excludeFlowPackage = &excludeFlowPackage
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) FilterAwemeAbnormalActive(filterAwemeAbnormalActive ToolsEstimateAudienceV2FilterAwemeAbnormalActive) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.filterAwemeAbnormalActive = &filterAwemeAbnormalActive
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) FilterAwemeFansCount(filterAwemeFansCount int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.filterAwemeFansCount = &filterAwemeFansCount
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) FilterOwnAwemeFans(filterOwnAwemeFans ToolsEstimateAudienceV2FilterOwnAwemeFans) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.filterOwnAwemeFans = &filterOwnAwemeFans
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) FlowPackage(flowPackage []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.flowPackage = &flowPackage
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) Gender(gender ToolsEstimateAudienceV2Gender) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.gender = &gender
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) Geolocation(geolocation []*ToolsEstimateAudienceV2GeolocationInner) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.geolocation = &geolocation
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) IncludeCustomActions(includeCustomActions []*ToolsEstimateAudienceV2IncludeCustomActionsInner) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.includeCustomActions = &includeCustomActions
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) InterestActionMode(interestActionMode ToolsEstimateAudienceV2InterestActionMode) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.interestActionMode = &interestActionMode
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) InterestCategories(interestCategories []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.interestCategories = &interestCategories
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) InterestTags(interestTags []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.interestTags = &interestTags
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) InterestWords(interestWords []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.interestWords = &interestWords
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) IosOsv(iosOsv ToolsEstimateAudienceV2IosOsv) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.iosOsv = &iosOsv
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) LaunchPrice(launchPrice []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.launchPrice = &launchPrice
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) LocationType(locationType ToolsEstimateAudienceV2LocationType) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.locationType = &locationType
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) Platform(platform []*ToolsEstimateAudienceV2Platform) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.platform = &platform
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) RetargetingTags(retargetingTags []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.retargetingTags = &retargetingTags
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) RetargetingTagsExclude(retargetingTagsExclude []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.retargetingTagsExclude = &retargetingTagsExclude
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) RetargetingTagsInclude(retargetingTagsInclude []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.retargetingTagsInclude = &retargetingTagsInclude
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) RetargetingType(retargetingType ToolsEstimateAudienceV2RetargetingType) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.retargetingType = &retargetingType
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) SuperiorPopularityType(superiorPopularityType ToolsEstimateAudienceV2SuperiorPopularityType) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.superiorPopularityType = &superiorPopularityType
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) UserType(userType []int64) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.userType = &userType
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) Execute() (*ToolsEstimateAudienceV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsEstimateAudienceGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsEstimateAudienceGet Method for OpenApi2ToolsEstimateAudienceGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsEstimateAudienceGetRequest
*/
func (a *ToolsEstimateAudienceV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsEstimateAudienceGetRequest {
	return &ApiOpenApi2ToolsEstimateAudienceGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsEstimateAudienceV2Response
func (a *ToolsEstimateAudienceV2ApiService) getExecute(r *ApiOpenApi2ToolsEstimateAudienceGetRequest) (*ToolsEstimateAudienceV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsEstimateAudienceV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/estimate_audience/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ac != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ac", r.ac)
	}
	if r.actionCategories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action_categories", r.actionCategories)
	}
	if r.actionDays != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action_days", r.actionDays)
	}
	if r.actionWords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action_words", r.actionWords)
	}
	if r.activateType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activate_type", r.activateType)
	}
	if r.adTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_tag", r.adTag)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.age != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "age", r.age)
	}
	if r.androidOsv != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "android_osv", r.androidOsv)
	}
	if r.appBehaviorTarget != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_behavior_target", r.appBehaviorTarget)
	}
	if r.appCategory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_category", r.appCategory)
	}
	if r.appIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_ids", r.appIds)
	}
	if r.articleCategory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "article_category", r.articleCategory)
	}
	if r.audiencePackageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "audience_package_id", r.audiencePackageId)
	}
	if r.autoExtendEnabled != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auto_extend_enabled", r.autoExtendEnabled)
	}
	if r.autoExtendTargets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auto_extend_targets", r.autoExtendTargets)
	}
	if r.awemeFansNumbers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_fans_numbers", r.awemeFansNumbers)
	}
	if r.carrier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "carrier", r.carrier)
	}
	if r.city != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "city", r.city)
	}
	if r.deviceBrand != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "device_brand", r.deviceBrand)
	}
	if r.deviceType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "device_type", r.deviceType)
	}
	if r.district != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "district", r.district)
	}
	if r.dpaLocalAudience != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dpa_local_audience", r.dpaLocalAudience)
	}
	if r.excludeCustomActions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_custom_actions", r.excludeCustomActions)
	}
	if r.excludeFlowPackage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_flow_package", r.excludeFlowPackage)
	}
	if r.filterAwemeAbnormalActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_aweme_abnormal_active", r.filterAwemeAbnormalActive)
	}
	if r.filterAwemeFansCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_aweme_fans_count", r.filterAwemeFansCount)
	}
	if r.filterOwnAwemeFans != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_own_aweme_fans", r.filterOwnAwemeFans)
	}
	if r.flowPackage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "flow_package", r.flowPackage)
	}
	if r.gender != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gender", r.gender)
	}
	if r.geolocation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "geolocation", r.geolocation)
	}
	if r.includeCustomActions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_custom_actions", r.includeCustomActions)
	}
	if r.interestActionMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interest_action_mode", r.interestActionMode)
	}
	if r.interestCategories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interest_categories", r.interestCategories)
	}
	if r.interestTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interest_tags", r.interestTags)
	}
	if r.interestWords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interest_words", r.interestWords)
	}
	if r.iosOsv != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ios_osv", r.iosOsv)
	}
	if r.launchPrice != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "launch_price", r.launchPrice)
	}
	if r.locationType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "location_type", r.locationType)
	}
	if r.platform != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "platform", r.platform)
	}
	if r.retargetingTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retargeting_tags", r.retargetingTags)
	}
	if r.retargetingTagsExclude != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retargeting_tags_exclude", r.retargetingTagsExclude)
	}
	if r.retargetingTagsInclude != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retargeting_tags_include", r.retargetingTagsInclude)
	}
	if r.retargetingType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retargeting_type", r.retargetingType)
	}
	if r.superiorPopularityType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "superior_popularity_type", r.superiorPopularityType)
	}
	if r.userType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user_type", r.userType)
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
