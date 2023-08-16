/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.12
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

// ToolsKeywordsBidRatioGetV30ApiService ToolsKeywordsBidRatioGetV30Api service
type ToolsKeywordsBidRatioGetV30ApiService service

type ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsKeywordsBidRatioGetV30ApiService
	advertiserId *int64
	filtering    *ToolsKeywordsBidRatioGetV30Filtering
}

func (r *ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest) Filtering(filtering ToolsKeywordsBidRatioGetV30Filtering) *ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest) Execute() (*ToolsKeywordsBidRatioGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsKeywordsBidRatioGetGet Method for OpenApiV30ToolsKeywordsBidRatioGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest
*/
func (a *ToolsKeywordsBidRatioGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest {
	return &ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsKeywordsBidRatioGetV30Response
func (a *ToolsKeywordsBidRatioGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsKeywordsBidRatioGetGetRequest) (*ToolsKeywordsBidRatioGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsKeywordsBidRatioGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/keywords_bid_ratio/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
