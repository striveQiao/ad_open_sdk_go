/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
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

// ToolsAdConvertOptimizedTargetGetV2ApiService ToolsAdConvertOptimizedTargetGetV2Api service
type ToolsAdConvertOptimizedTargetGetV2ApiService service

type ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest struct {
	ctx                context.Context
	ApiService         *ToolsAdConvertOptimizedTargetGetV2ApiService
	advertiserId       *int64
	appSchema          *string
	appType            *ToolsAdConvertOptimizedTargetGetV2AppType
	campaignType       *ToolsAdConvertOptimizedTargetGetV2CampaignType
	convertId          *int64
	convertName        *string
	convertType        *ToolsAdConvertOptimizedTargetGetV2ConvertType
	dedicateType       *ToolsAdConvertOptimizedTargetGetV2DedicateType
	deepExternalAction *ToolsAdConvertOptimizedTargetGetV2DeepExternalAction
	externalAction     *ToolsAdConvertOptimizedTargetGetV2ExternalAction
	externalUrl        *string
	itunesUrl          *string
	landingType        *ToolsAdConvertOptimizedTargetGetV2LandingType
	launchTargetType   *ToolsAdConvertOptimizedTargetGetV2LaunchTargetType
	marketingPurpose   *ToolsAdConvertOptimizedTargetGetV2MarketingPurpose
	packageName        *string
	page               *int64
	pageSize           *int64
	promotionContent   *ToolsAdConvertOptimizedTargetGetV2PromotionContent
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) AppSchema(appSchema string) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.appSchema = &appSchema
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) AppType(appType ToolsAdConvertOptimizedTargetGetV2AppType) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.appType = &appType
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) CampaignType(campaignType ToolsAdConvertOptimizedTargetGetV2CampaignType) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.campaignType = &campaignType
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) ConvertId(convertId int64) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.convertId = &convertId
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) ConvertName(convertName string) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.convertName = &convertName
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) ConvertType(convertType ToolsAdConvertOptimizedTargetGetV2ConvertType) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.convertType = &convertType
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) DedicateType(dedicateType ToolsAdConvertOptimizedTargetGetV2DedicateType) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.dedicateType = &dedicateType
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) DeepExternalAction(deepExternalAction ToolsAdConvertOptimizedTargetGetV2DeepExternalAction) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.deepExternalAction = &deepExternalAction
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) ExternalAction(externalAction ToolsAdConvertOptimizedTargetGetV2ExternalAction) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.externalAction = &externalAction
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) ExternalUrl(externalUrl string) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.externalUrl = &externalUrl
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) ItunesUrl(itunesUrl string) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.itunesUrl = &itunesUrl
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) LandingType(landingType ToolsAdConvertOptimizedTargetGetV2LandingType) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.landingType = &landingType
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) LaunchTargetType(launchTargetType ToolsAdConvertOptimizedTargetGetV2LaunchTargetType) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.launchTargetType = &launchTargetType
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) MarketingPurpose(marketingPurpose ToolsAdConvertOptimizedTargetGetV2MarketingPurpose) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.marketingPurpose = &marketingPurpose
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) PackageName(packageName string) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.packageName = &packageName
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) Page(page int64) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) PromotionContent(promotionContent ToolsAdConvertOptimizedTargetGetV2PromotionContent) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.promotionContent = &promotionContent
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) Execute() (*ToolsAdConvertOptimizedTargetGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAdConvertOptimizedTargetGetGet Method for OpenApi2ToolsAdConvertOptimizedTargetGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest
*/
func (a *ToolsAdConvertOptimizedTargetGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest {
	return &ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAdConvertOptimizedTargetGetV2Response
func (a *ToolsAdConvertOptimizedTargetGetV2ApiService) getExecute(r *ApiOpenApi2ToolsAdConvertOptimizedTargetGetGetRequest) (*ToolsAdConvertOptimizedTargetGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAdConvertOptimizedTargetGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/ad_convert/optimized_target/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.appSchema != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_schema", r.appSchema)
	}
	if r.appType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_type", r.appType)
	}
	if r.campaignType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_type", r.campaignType)
	}
	if r.convertId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "convert_id", r.convertId)
	}
	if r.convertName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "convert_name", r.convertName)
	}
	if r.convertType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "convert_type", r.convertType)
	}
	if r.dedicateType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dedicate_type", r.dedicateType)
	}
	if r.deepExternalAction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deep_external_action", r.deepExternalAction)
	}
	if r.externalAction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "external_action", r.externalAction)
	}
	if r.externalUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "external_url", r.externalUrl)
	}
	if r.itunesUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itunes_url", r.itunesUrl)
	}
	if r.landingType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "landing_type", r.landingType)
	}
	if r.launchTargetType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "launch_target_type", r.launchTargetType)
	}
	if r.marketingPurpose != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_purpose", r.marketingPurpose)
	}
	if r.packageName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "package_name", r.packageName)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.promotionContent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "promotion_content", r.promotionContent)
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
