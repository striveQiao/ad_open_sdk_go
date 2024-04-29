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

// ClueSmartphoneGetV2ApiService ClueSmartphoneGetV2Api service
type ClueSmartphoneGetV2ApiService service

type ApiOpenApi2ClueSmartphoneGetGetRequest struct {
	ctx          context.Context
	ApiService   *ClueSmartphoneGetV2ApiService
	advertiserId *int64
	endTime      **string
	instanceIds  *[]int64
	isDel        *ClueSmartphoneGetV2IsDel
	page         *int64
	pageSize     *int64
	startTime    **string
}

func (r *ApiOpenApi2ClueSmartphoneGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ClueSmartphoneGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ClueSmartphoneGetGetRequest) EndTime(endTime *string) *ApiOpenApi2ClueSmartphoneGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ClueSmartphoneGetGetRequest) InstanceIds(instanceIds []int64) *ApiOpenApi2ClueSmartphoneGetGetRequest {
	r.instanceIds = &instanceIds
	return r
}

func (r *ApiOpenApi2ClueSmartphoneGetGetRequest) IsDel(isDel ClueSmartphoneGetV2IsDel) *ApiOpenApi2ClueSmartphoneGetGetRequest {
	r.isDel = &isDel
	return r
}

func (r *ApiOpenApi2ClueSmartphoneGetGetRequest) Page(page int64) *ApiOpenApi2ClueSmartphoneGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ClueSmartphoneGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ClueSmartphoneGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ClueSmartphoneGetGetRequest) StartTime(startTime *string) *ApiOpenApi2ClueSmartphoneGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApi2ClueSmartphoneGetGetRequest) Execute() (*ClueSmartphoneGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ClueSmartphoneGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ClueSmartphoneGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ClueSmartphoneGetGetRequest) WithLog(enable bool) *ApiOpenApi2ClueSmartphoneGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ClueSmartphoneGetGet Method for OpenApi2ClueSmartphoneGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ClueSmartphoneGetGetRequest
*/
func (a *ClueSmartphoneGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ClueSmartphoneGetGetRequest {
	return &ApiOpenApi2ClueSmartphoneGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ClueSmartphoneGetV2Response
func (a *ClueSmartphoneGetV2ApiService) getExecute(r *ApiOpenApi2ClueSmartphoneGetGetRequest) (*ClueSmartphoneGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ClueSmartphoneGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/clue/smartphone/get/"

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
	if r.instanceIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "instance_ids", r.instanceIds)
	}
	if r.isDel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_del", r.isDel)
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
