/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
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

// ProjectListV30ApiService ProjectListV30Api service
type ProjectListV30ApiService service

type ApiOpenApiV30ProjectListGetRequest struct {
	ctx          context.Context
	ApiService   *ProjectListV30ApiService
	advertiserId *int64
	fields       *[]string
	filtering    *ProjectListV30Filtering
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApiV30ProjectListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ProjectListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ProjectListGetRequest) Fields(fields []string) *ApiOpenApiV30ProjectListGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApiV30ProjectListGetRequest) Filtering(filtering ProjectListV30Filtering) *ApiOpenApiV30ProjectListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30ProjectListGetRequest) Page(page int64) *ApiOpenApiV30ProjectListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30ProjectListGetRequest) PageSize(pageSize int64) *ApiOpenApiV30ProjectListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ProjectListGetRequest) Execute() (*ProjectListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ProjectListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ProjectListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ProjectListGetRequest) WithLog(enable bool) *ApiOpenApiV30ProjectListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ProjectListGet Method for OpenApiV30ProjectListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ProjectListGetRequest
*/
func (a *ProjectListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ProjectListGetRequest {
	return &ApiOpenApiV30ProjectListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProjectListV30Response
func (a *ProjectListV30ApiService) getExecute(r *ApiOpenApiV30ProjectListGetRequest) (*ProjectListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ProjectListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/project/list/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
