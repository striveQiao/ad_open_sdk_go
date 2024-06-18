/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
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

// ToolsStarTaskSettlementConfigV2ApiService ToolsStarTaskSettlementConfigV2Api service
type ToolsStarTaskSettlementConfigV2ApiService service

type ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest struct {
	ctx                    context.Context
	ApiService             *ToolsStarTaskSettlementConfigV2ApiService
	advertiserId           *int64
	firstIndustryId        *int64
	secondIndustryId       *int64
	starMaterialFirstType  *int64
	starMaterialSecondType *int64
	starTaskExternalAction *ToolsStarTaskSettlementConfigV2StarTaskExternalAction
}

// 广告账户id
func (r *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 一级行业id
func (r *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest) FirstIndustryId(firstIndustryId int64) *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest {
	r.firstIndustryId = &firstIndustryId
	return r
}

// 二级行业id
func (r *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest) SecondIndustryId(secondIndustryId int64) *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest {
	r.secondIndustryId = &secondIndustryId
	return r
}

// 素材一级类目id
func (r *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest) StarMaterialFirstType(starMaterialFirstType int64) *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest {
	r.starMaterialFirstType = &starMaterialFirstType
	return r
}

// 素材二级类目id
func (r *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest) StarMaterialSecondType(starMaterialSecondType int64) *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest {
	r.starMaterialSecondType = &starMaterialSecondType
	return r
}

// 优化目标，仅对达人流量生效。支持3种可选值）
func (r *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest) StarTaskExternalAction(starTaskExternalAction ToolsStarTaskSettlementConfigV2StarTaskExternalAction) *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest {
	r.starTaskExternalAction = &starTaskExternalAction
	return r
}

func (r *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest) Execute() (*ToolsStarTaskSettlementConfigV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsStarTaskSettlementConfigGet Method for OpenApi2ToolsStarTaskSettlementConfigGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest
*/
func (a *ToolsStarTaskSettlementConfigV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest {
	return &ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsStarTaskSettlementConfigV2Response
func (a *ToolsStarTaskSettlementConfigV2ApiService) getExecute(r *ApiOpenApi2ToolsStarTaskSettlementConfigGetRequest) (*ToolsStarTaskSettlementConfigV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsStarTaskSettlementConfigV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/star_task/settlement_config/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.firstIndustryId == nil {
		return localVarReturnValue, nil, ReportError("firstIndustryId is required and must be specified")
	}
	if r.secondIndustryId == nil {
		return localVarReturnValue, nil, ReportError("secondIndustryId is required and must be specified")
	}
	if r.starMaterialFirstType == nil {
		return localVarReturnValue, nil, ReportError("starMaterialFirstType is required and must be specified")
	}
	if r.starMaterialSecondType == nil {
		return localVarReturnValue, nil, ReportError("starMaterialSecondType is required and must be specified")
	}
	if r.starTaskExternalAction == nil {
		return localVarReturnValue, nil, ReportError("starTaskExternalAction is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "first_industry_id", r.firstIndustryId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "second_industry_id", r.secondIndustryId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "star_material_first_type", r.starMaterialFirstType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "star_material_second_type", r.starMaterialSecondType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "star_task_external_action", r.starTaskExternalAction)
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
