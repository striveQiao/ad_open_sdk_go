/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
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

// ToolsEventConvertOptimizedGoalGetV30ApiService ToolsEventConvertOptimizedGoalGetV30Api service
type ToolsEventConvertOptimizedGoalGetV30ApiService service

type ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest struct {
	ctx                context.Context
	ApiService         *ToolsEventConvertOptimizedGoalGetV30ApiService
	landingType        *ToolsEventConvertOptimizedGoalGetV30LandingType
	marketingPurpose   *ToolsEventConvertOptimizedGoalGetV30MarketingPurpose
	assetType          *ToolsEventConvertOptimizedGoalGetV30AssetType
	advertiserId       *int64
	siteId             *int64
	miniProgramId      *string
	assetId            *int64
	campaignType       *ToolsEventConvertOptimizedGoalGetV30CampaignType
	microAppInstanceId *int64
}

func (r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) LandingType(landingType ToolsEventConvertOptimizedGoalGetV30LandingType) *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest {
	r.landingType = &landingType
	return r
}

func (r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) MarketingPurpose(marketingPurpose ToolsEventConvertOptimizedGoalGetV30MarketingPurpose) *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest {
	r.marketingPurpose = &marketingPurpose
	return r
}

func (r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) AssetType(assetType ToolsEventConvertOptimizedGoalGetV30AssetType) *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest {
	r.assetType = &assetType
	return r
}

func (r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) SiteId(siteId int64) *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest {
	r.siteId = &siteId
	return r
}

func (r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) MiniProgramId(miniProgramId string) *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest {
	r.miniProgramId = &miniProgramId
	return r
}

func (r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) AssetId(assetId int64) *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest {
	r.assetId = &assetId
	return r
}

func (r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) CampaignType(campaignType ToolsEventConvertOptimizedGoalGetV30CampaignType) *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest {
	r.campaignType = &campaignType
	return r
}

func (r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) MicroAppInstanceId(microAppInstanceId int64) *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest {
	r.microAppInstanceId = &microAppInstanceId
	return r
}

func (r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) Execute() (*ToolsEventConvertOptimizedGoalGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsEventConvertOptimizedGoalGetGet Method for OpenApiV30ToolsEventConvertOptimizedGoalGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest
*/
func (a *ToolsEventConvertOptimizedGoalGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest {
	return &ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsEventConvertOptimizedGoalGetV30Response
func (a *ToolsEventConvertOptimizedGoalGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequest) (*ToolsEventConvertOptimizedGoalGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsEventConvertOptimizedGoalGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/event_convert/optimized_goal/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.landingType == nil {
		return localVarReturnValue, nil, ReportError("landingType is required and must be specified")
	}
	if r.marketingPurpose == nil {
		return localVarReturnValue, nil, ReportError("marketingPurpose is required and must be specified")
	}
	if r.assetType == nil {
		return localVarReturnValue, nil, ReportError("assetType is required and must be specified")
	}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId)
	}
	if r.miniProgramId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mini_program_id", r.miniProgramId)
	}
	if r.assetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "asset_id", r.assetId)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "landing_type", r.landingType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_purpose", r.marketingPurpose)
	parameterAddToHeaderOrQuery(localVarQueryParams, "asset_type", r.assetType)
	if r.campaignType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_type", r.campaignType)
	}
	if r.microAppInstanceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "micro_app_instance_id", r.microAppInstanceId)
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
