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

// QianchuanAwemeOrderGetV10ApiService QianchuanAwemeOrderGetV10Api service
type QianchuanAwemeOrderGetV10ApiService service

type ApiOpenApiV10QianchuanAwemeOrderGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAwemeOrderGetV10ApiService
	advertiserId *int64
	filtering    *QianchuanAwemeOrderGetV10Filtering
	cursor       *int64
	count        *QianchuanAwemeOrderGetV10Count
}

// 千川广告主id
func (r *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest) Filtering(filtering QianchuanAwemeOrderGetV10Filtering) *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest {
	r.filtering = &filtering
	return r
}

// 页码，默认值：0
func (r *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest) Cursor(cursor int64) *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest {
	r.cursor = &cursor
	return r
}

// 页面大小，允许值：10, 20, 50，默认值：10
func (r *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest) Count(count QianchuanAwemeOrderGetV10Count) *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest {
	r.count = &count
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest) Execute() (*QianchuanAwemeOrderGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAwemeOrderGetGet Method for OpenApiV10QianchuanAwemeOrderGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAwemeOrderGetGetRequest
*/
func (a *QianchuanAwemeOrderGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest {
	return &ApiOpenApiV10QianchuanAwemeOrderGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAwemeOrderGetV10Response
func (a *QianchuanAwemeOrderGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAwemeOrderGetGetRequest) (*QianchuanAwemeOrderGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAwemeOrderGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/aweme/order/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor)
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count)
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
