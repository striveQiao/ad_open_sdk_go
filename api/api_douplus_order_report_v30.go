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

// DouplusOrderReportV30ApiService DouplusOrderReportV30Api service
type DouplusOrderReportV30ApiService service

type ApiOpenApiV30DouplusOrderReportGetRequest struct {
	ctx         context.Context
	ApiService  *DouplusOrderReportV30ApiService
	awemeSecUid *string
	statTime    *DouplusOrderReportV30StatTime
	groupBy     *[]*DouplusOrderReportV30GroupBy
	filter      *DouplusOrderReportV30Filter
	pageSize    *int64
	page        *int64
}

func (r *ApiOpenApiV30DouplusOrderReportGetRequest) AwemeSecUid(awemeSecUid string) *ApiOpenApiV30DouplusOrderReportGetRequest {
	r.awemeSecUid = &awemeSecUid
	return r
}

func (r *ApiOpenApiV30DouplusOrderReportGetRequest) StatTime(statTime DouplusOrderReportV30StatTime) *ApiOpenApiV30DouplusOrderReportGetRequest {
	r.statTime = &statTime
	return r
}

func (r *ApiOpenApiV30DouplusOrderReportGetRequest) GroupBy(groupBy []*DouplusOrderReportV30GroupBy) *ApiOpenApiV30DouplusOrderReportGetRequest {
	r.groupBy = &groupBy
	return r
}

func (r *ApiOpenApiV30DouplusOrderReportGetRequest) Filter(filter DouplusOrderReportV30Filter) *ApiOpenApiV30DouplusOrderReportGetRequest {
	r.filter = &filter
	return r
}

func (r *ApiOpenApiV30DouplusOrderReportGetRequest) PageSize(pageSize int64) *ApiOpenApiV30DouplusOrderReportGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30DouplusOrderReportGetRequest) Page(page int64) *ApiOpenApiV30DouplusOrderReportGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30DouplusOrderReportGetRequest) Execute() (*DouplusOrderReportV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30DouplusOrderReportGetRequest) AccessToken(accessToken string) *ApiOpenApiV30DouplusOrderReportGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30DouplusOrderReportGetRequest) WithLog(enable bool) *ApiOpenApiV30DouplusOrderReportGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30DouplusOrderReportGet Method for OpenApiV30DouplusOrderReportGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30DouplusOrderReportGetRequest
*/
func (a *DouplusOrderReportV30ApiService) Get(ctx context.Context) *ApiOpenApiV30DouplusOrderReportGetRequest {
	return &ApiOpenApiV30DouplusOrderReportGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DouplusOrderReportV30Response
func (a *DouplusOrderReportV30ApiService) getExecute(r *ApiOpenApiV30DouplusOrderReportGetRequest) (*DouplusOrderReportV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *DouplusOrderReportV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/douplus/order/report/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.awemeSecUid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_sec_uid", r.awemeSecUid)
	}
	if r.statTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stat_time", r.statTime)
	}
	if r.groupBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group_by", r.groupBy)
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
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
