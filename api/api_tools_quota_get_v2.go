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

// ToolsQuotaGetV2ApiService ToolsQuotaGetV2Api service
type ToolsQuotaGetV2ApiService service

type ApiOpenApi2ToolsQuotaGetGetRequest struct {
	ctx           context.Context
	ApiService    *ToolsQuotaGetV2ApiService
	advertiserId  *int64
	campaignType  *ToolsQuotaGetV2CampaignType
	deliveryRange *ToolsQuotaGetV2DeliveryRange
}

// 广告账户id
func (r *ApiOpenApi2ToolsQuotaGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsQuotaGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告组类型，&#x60;FEED&#x60;：信息流  &#x60;SEARCH&#x60;:搜索广告
func (r *ApiOpenApi2ToolsQuotaGetGetRequest) CampaignType(campaignType ToolsQuotaGetV2CampaignType) *ApiOpenApi2ToolsQuotaGetGetRequest {
	r.campaignType = &campaignType
	return r
}

func (r *ApiOpenApi2ToolsQuotaGetGetRequest) DeliveryRange(deliveryRange ToolsQuotaGetV2DeliveryRange) *ApiOpenApi2ToolsQuotaGetGetRequest {
	r.deliveryRange = &deliveryRange
	return r
}

func (r *ApiOpenApi2ToolsQuotaGetGetRequest) Execute() (*ToolsQuotaGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsQuotaGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsQuotaGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsQuotaGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsQuotaGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsQuotaGetGet Method for OpenApi2ToolsQuotaGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsQuotaGetGetRequest
*/
func (a *ToolsQuotaGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsQuotaGetGetRequest {
	return &ApiOpenApi2ToolsQuotaGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsQuotaGetV2Response
func (a *ToolsQuotaGetV2ApiService) getExecute(r *ApiOpenApi2ToolsQuotaGetGetRequest) (*ToolsQuotaGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsQuotaGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/quota/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.campaignType == nil {
		return localVarReturnValue, nil, ReportError("campaignType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_type", r.campaignType)
	if r.deliveryRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delivery_range", r.deliveryRange)
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
