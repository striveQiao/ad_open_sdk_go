/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
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

// NativeAnchorGetV30ApiService NativeAnchorGetV30Api service
type NativeAnchorGetV30ApiService service

type ApiOpenApiV30NativeAnchorGetGetRequest struct {
	ctx          context.Context
	ApiService   *NativeAnchorGetV30ApiService
	advertiserId *int64
	page         *int32
	pageSize     *int32
	filtering    *NativeAnchorGetV30Filtering
}

func (r *ApiOpenApiV30NativeAnchorGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30NativeAnchorGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30NativeAnchorGetGetRequest) Page(page int32) *ApiOpenApiV30NativeAnchorGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30NativeAnchorGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV30NativeAnchorGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30NativeAnchorGetGetRequest) Filtering(filtering NativeAnchorGetV30Filtering) *ApiOpenApiV30NativeAnchorGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30NativeAnchorGetGetRequest) Execute() (*NativeAnchorGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30NativeAnchorGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30NativeAnchorGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30NativeAnchorGetGetRequest) WithLog(enable bool) *ApiOpenApiV30NativeAnchorGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30NativeAnchorGetGet Method for OpenApiV30NativeAnchorGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30NativeAnchorGetGetRequest
*/
func (a *NativeAnchorGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30NativeAnchorGetGetRequest {
	return &ApiOpenApiV30NativeAnchorGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NativeAnchorGetV30Response
func (a *NativeAnchorGetV30ApiService) getExecute(r *ApiOpenApiV30NativeAnchorGetGetRequest) (*NativeAnchorGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *NativeAnchorGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/native_anchor/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
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
