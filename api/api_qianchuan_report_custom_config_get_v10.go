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

// QianchuanReportCustomConfigGetV10ApiService QianchuanReportCustomConfigGetV10Api service
type QianchuanReportCustomConfigGetV10ApiService service

type ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanReportCustomConfigGetV10ApiService
	advertiserId *int64
	dataTopics   *[]*QianchuanReportCustomConfigGetV10DataTopics
}

func (r *ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest) DataTopics(dataTopics []*QianchuanReportCustomConfigGetV10DataTopics) *ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest {
	r.dataTopics = &dataTopics
	return r
}

func (r *ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest) Execute() (*QianchuanReportCustomConfigGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanReportCustomConfigGetGet Method for OpenApiV10QianchuanReportCustomConfigGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest
*/
func (a *QianchuanReportCustomConfigGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest {
	return &ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanReportCustomConfigGetV10Response
func (a *QianchuanReportCustomConfigGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanReportCustomConfigGetGetRequest) (*QianchuanReportCustomConfigGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanReportCustomConfigGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/report/custom/config/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.dataTopics == nil {
		return localVarReturnValue, nil, ReportError("dataTopics is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "data_topics", r.dataTopics)
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
