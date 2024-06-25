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

// ReportCustomAsyncTaskDownloadV30ApiService ReportCustomAsyncTaskDownloadV30Api service
type ReportCustomAsyncTaskDownloadV30ApiService service

type ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest struct {
	ctx          context.Context
	ApiService   *ReportCustomAsyncTaskDownloadV30ApiService
	advertiserId *int64
	taskId       *int64
}

func (r *ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest) TaskId(taskId int64) *ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest {
	r.taskId = &taskId
	return r
}

func (r *ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest) Execute() (*ReportCustomAsyncTaskDownloadV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest) WithLog(enable bool) *ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ReportCustomAsyncTaskDownloadGet Method for OpenApiV30ReportCustomAsyncTaskDownloadGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest
*/
func (a *ReportCustomAsyncTaskDownloadV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest {
	return &ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportCustomAsyncTaskDownloadV30Response
func (a *ReportCustomAsyncTaskDownloadV30ApiService) getExecute(r *ApiOpenApiV30ReportCustomAsyncTaskDownloadGetRequest) (*ReportCustomAsyncTaskDownloadV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportCustomAsyncTaskDownloadV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/report/custom/async_task/download/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.taskId == nil {
		return localVarReturnValue, nil, ReportError("taskId is required and must be specified")
	}
	if *r.taskId < 1 {
		return localVarReturnValue, nil, ReportError("taskId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "task_id", r.taskId)
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
