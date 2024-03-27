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

// ToolsLogSearchV2ApiService ToolsLogSearchV2Api service
type ToolsLogSearchV2ApiService service

type ApiOpenApi2ToolsLogSearchGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsLogSearchV2ApiService
	advertiserId *int64
	endTime      **string
	objectId     *[]int64
	page         *int64
	pageSize     *int64
	startTime    **string
}

func (r *ApiOpenApi2ToolsLogSearchGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsLogSearchGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsLogSearchGetRequest) EndTime(endTime *string) *ApiOpenApi2ToolsLogSearchGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ToolsLogSearchGetRequest) ObjectId(objectId []int64) *ApiOpenApi2ToolsLogSearchGetRequest {
	r.objectId = &objectId
	return r
}

func (r *ApiOpenApi2ToolsLogSearchGetRequest) Page(page int64) *ApiOpenApi2ToolsLogSearchGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsLogSearchGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsLogSearchGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsLogSearchGetRequest) StartTime(startTime *string) *ApiOpenApi2ToolsLogSearchGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApi2ToolsLogSearchGetRequest) Execute() (*ToolsLogSearchV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsLogSearchGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsLogSearchGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsLogSearchGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsLogSearchGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsLogSearchGet Method for OpenApi2ToolsLogSearchGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsLogSearchGetRequest
*/
func (a *ToolsLogSearchV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsLogSearchGetRequest {
	return &ApiOpenApi2ToolsLogSearchGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsLogSearchV2Response
func (a *ToolsLogSearchV2ApiService) getExecute(r *ApiOpenApi2ToolsLogSearchGetRequest) (*ToolsLogSearchV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsLogSearchV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/log_search/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	if r.objectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "object_id", r.objectId)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
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
