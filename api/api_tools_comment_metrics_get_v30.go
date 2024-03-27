/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
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

// ToolsCommentMetricsGetV30ApiService ToolsCommentMetricsGetV30Api service
type ToolsCommentMetricsGetV30ApiService service

type ApiOpenApiV30ToolsCommentMetricsGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsCommentMetricsGetV30ApiService
	advertiserId *int64
	startTime    *string
	endTime      *string
	filtering    *ToolsCommentMetricsGetV30Filtering
}

// 广告主id
func (r *ApiOpenApiV30ToolsCommentMetricsGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsCommentMetricsGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ToolsCommentMetricsGetGetRequest) StartTime(startTime string) *ApiOpenApiV30ToolsCommentMetricsGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV30ToolsCommentMetricsGetGetRequest) EndTime(endTime string) *ApiOpenApiV30ToolsCommentMetricsGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV30ToolsCommentMetricsGetGetRequest) Filtering(filtering ToolsCommentMetricsGetV30Filtering) *ApiOpenApiV30ToolsCommentMetricsGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30ToolsCommentMetricsGetGetRequest) Execute() (*ToolsCommentMetricsGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsCommentMetricsGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsCommentMetricsGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsCommentMetricsGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsCommentMetricsGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsCommentMetricsGetGet Method for OpenApiV30ToolsCommentMetricsGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsCommentMetricsGetGetRequest
*/
func (a *ToolsCommentMetricsGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsCommentMetricsGetGetRequest {
	return &ApiOpenApiV30ToolsCommentMetricsGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsCommentMetricsGetV30Response
func (a *ToolsCommentMetricsGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsCommentMetricsGetGetRequest) (*ToolsCommentMetricsGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsCommentMetricsGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/comment_metrics/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
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
