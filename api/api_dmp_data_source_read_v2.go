/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
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

// DmpDataSourceReadV2ApiService DmpDataSourceReadV2Api service
type DmpDataSourceReadV2ApiService service

type ApiOpenApi2DmpDataSourceReadGetRequest struct {
	ctx              context.Context
	ApiService       *DmpDataSourceReadV2ApiService
	advertiserId     *int64
	dataSourceIdList *[]string
}

func (r *ApiOpenApi2DmpDataSourceReadGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2DmpDataSourceReadGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2DmpDataSourceReadGetRequest) DataSourceIdList(dataSourceIdList []string) *ApiOpenApi2DmpDataSourceReadGetRequest {
	r.dataSourceIdList = &dataSourceIdList
	return r
}

func (r *ApiOpenApi2DmpDataSourceReadGetRequest) Execute() (*DmpDataSourceReadV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2DmpDataSourceReadGetRequest) AccessToken(accessToken string) *ApiOpenApi2DmpDataSourceReadGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DmpDataSourceReadGetRequest) WithLog(enable bool) *ApiOpenApi2DmpDataSourceReadGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DmpDataSourceReadGet Method for OpenApi2DmpDataSourceReadGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DmpDataSourceReadGetRequest
*/
func (a *DmpDataSourceReadV2ApiService) Get(ctx context.Context) *ApiOpenApi2DmpDataSourceReadGetRequest {
	return &ApiOpenApi2DmpDataSourceReadGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DmpDataSourceReadV2Response
func (a *DmpDataSourceReadV2ApiService) getExecute(r *ApiOpenApi2DmpDataSourceReadGetRequest) (*DmpDataSourceReadV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *DmpDataSourceReadV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/dmp/data_source/read/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.dataSourceIdList != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "data_source_id_list", r.dataSourceIdList)
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
